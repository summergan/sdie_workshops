# Spec: Gitea Actions Run Status Aggregation

## SDIE Phase

Specification.

## Assumptions

1. The requested behavior is for Gitea Actions workflow-run status aggregation
   from its jobs, plus the run-list status filter visibility.
2. This slice uses the current local model: `success`, `failure`, `cancelled`,
   `skipped`, `waiting`, `running`, and `blocked`. Newer upstream states such as
   `cancelling` are out of scope because this branch does not define them.
3. Runner protocol result conversion and commit status mapping are out of scope
   unless they directly break workflow-run display.
4. The existing status icon template and Vue component already know how to
   render `cancelled`, `skipped`, and `blocked`; the primary gap is aggregation
   and filter exposure.

## Business Context

- User/problem: repository users need the workflow run header and Actions list
  to summarize the real outcome of all jobs in a workflow.
- Current pain: the aggregate only distinguishes failure, success, waiting, and
  running. All-skipped workflows look like running workflows, cancelled jobs make
  the run look failed, and blocked runs are not exposed as a first-class state.
- Constraint-heavy brownfield context: Actions status is shared by models,
  routers, templates, Vue, runner APIs, and commit status creation. The fix must
  stay inside the existing status enum and avoid schema or frontend rewrites.
- Success outcome: every aggregate status the model can produce is visible and
  test-protected in the Actions page flow.

## Structured Requirements

### Added

- Add explicit aggregation behavior for `cancelled`, `skipped`, and `blocked`.
- Add status filter entries for `Cancelled`, `Skipped`, and `Blocked` on the
  Actions run list.
- Render Actions status filter entries with the same status icon and localized
  label style used elsewhere in Actions, instead of raw status strings.
- Add regression tests that reproduce the three reported user-visible failures.

### Modified

- Workflow-run aggregation must return `Skipped` when every job is skipped.
- Workflow-run aggregation must return `Cancelled` when a cancelled terminal job
  determines the run outcome instead of collapsing it into `Failure`.
- Workflow-run aggregation must return `Blocked` when no job is waiting/running
  and at least one job is blocked.
- Mixed `success` and `skipped` terminal jobs still aggregate as `Success`.

### Removed

- None. Existing success, failure, waiting, and running behavior remains
  supported.

## Domain Semantics

- Domain terms:
  - Job status: status of one `ActionRunJob`.
  - Run status: aggregate status stored on `ActionRun` and rendered in Actions
    list/detail pages.
  - Terminal status: success, failure, cancelled, skipped.
  - Pending status: waiting, running, blocked.
- Existing repo concepts involved:
  - `Status.IsDone()` defines terminal job states.
  - `UpdateRunJob` recalculates the run status after a job status update.
  - `GetStatusInfoList` defines status choices in the Actions list filter.
  - `templates/repo/actions/status.tmpl` and `ActionRunStatus.vue` render
    visible status icons once the model exposes the correct string.
- Boundary cases:
  - All jobs skipped: aggregate `skipped`.
  - Success plus skipped: aggregate `success`.
  - Cancelled plus success/skipped/failure: aggregate `cancelled` when no
    pending job outranks the terminal result.
  - Running job present: aggregate `running`.
  - Waiting job present and no running job: aggregate `waiting`.
  - Blocked job present and no running/waiting job: aggregate `blocked`.
  - Empty job list: aggregate `unknown` fallback.

## Acceptance Criteria / BDD

- Given all jobs in a workflow run are skipped, when Gitea aggregates the run,
  then the run status is `skipped`.
- Given a workflow run has a cancelled job and no waiting/running/blocked jobs,
  when Gitea aggregates the run, then the run status is `cancelled`.
- Given a workflow run has a blocked job and no waiting/running jobs, when Gitea
  aggregates the run, then the run status is `blocked`.
- Given a workflow run has success and skipped jobs, when Gitea aggregates the
  run, then the run status remains `success`.
- Given a user opens the Actions page status filter, when the menu renders, then
  `Canceled`, `Skipped`, and `Blocked` are visible filter options.
- Given a user opens the Actions page status filter, when status options render,
  then every aggregate status option has a status icon and localized label.

## Commands

- Narrow UT:
  `PATH="/Users/summer/.asdf/shims:$PATH" TAGS="sqlite sqlite_unlock_notify" go test -tags "sqlite sqlite_unlock_notify" ./models/actions -run 'TestAggregateJobStatus|TestGetStatusInfoListIncludesAggregateStatuses' -count=1`
- Integration test:
  `PATH="/Users/summer/.asdf/shims:$PATH" TAGS="sqlite sqlite_unlock_notify" make test-sqlite#TestActionsRunStatusFilterOptions`
- E2E UI test:
  `PATH="/Users/summer/.local/bin:/Users/summer/.asdf/shims:$PATH" TAGS="sqlite sqlite_unlock_notify" GITEA_E2E_CHROME_PATH="/Applications/Google Chrome.app/Contents/MacOS/Google Chrome" make test-e2e-sqlite#actions-status-filter`
- Build:
  `PATH="/Users/summer/.asdf/shims:$PATH" TAGS="sqlite sqlite_unlock_notify" make build`

## Boundaries

- Always: keep aggregation in `models/actions`; routers and templates consume
  the resulting status string.
- Always: expose every aggregate status in `GetStatusInfoList` except
  `StatusUnknown`.
- Always: write RED tests before changing aggregation behavior.
- Ask first: adding statuses, schema migrations, runner protocol changes,
  commit status semantics changes, or frontend rewrites.
- Never: make `cancelled`, `skipped`, or `blocked` appear by hardcoding strings
  in templates while leaving model state wrong.

## Open Questions

- None for this slice.
