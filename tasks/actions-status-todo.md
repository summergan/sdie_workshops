# Todo: Gitea Actions Run Status Aggregation

## Implementation Tasks

- [x] Task: Extend the repo-local harness for Actions status work.
  - SDIE mapping: Specification / Design support.
  - Acceptance: repo map, testing reference, and AGENTS commands include the
    Actions status slice.
  - RED: harness references only described the previous Tags Search slice.
  - GREEN: Actions status files, ownership, and commands are documented.
  - Evaluation: future Actions status work can start from local references
    instead of conversational memory.
  - Files: `AGENTS.md`,
    `.codex/skills/sdie-harness-loop/references/repo-map.md`,
    `.codex/skills/sdie-harness-loop/references/testing.md`.

- [x] Task: Write the SDIE specification and design.
  - SDIE mapping: Specification / Design.
  - Acceptance: spec and plan capture requirements, semantics, tasks,
    cross-checks, risks, and loop-back triggers.
  - RED: current task had only chat requirements.
  - GREEN: task-specific artifacts exist under `tasks/actions-status-*.md`.
  - Evaluation: BDD cases now map to planned tests.
  - Files: `tasks/actions-status-spec.md`, `tasks/actions-status-plan.md`,
    `tasks/actions-status-todo.md`.

- [x] Task: Write RED model coverage.
  - SDIE mapping: Implementation / Testing.
  - Acceptance: tests reproduce all-skipped, cancelled, blocked, and
    success-plus-skipped aggregation semantics.
  - RED: targeted UT failed first. Current code returned `success` for all
    skipped jobs, `failure` for cancelled jobs, `running` for blocked/waiting
    combinations, and `success` for an empty job list.
  - GREEN: targeted UT passes after aggregation priority is made explicit.
  - Evaluation: the model contract now covers every local aggregate status and
    the unknown fallback.
  - Files: `models/actions/run_job_test.go`.

- [x] Task: Write RED status filter coverage.
  - SDIE mapping: Implementation / Testing.
  - Acceptance: tests prove cancelled, skipped, and blocked are present in the
    status filter contract.
  - RED: model test failed because `GetStatusInfoList` returned only success,
    failure, waiting, and running; integration test failed because the rendered
    Actions status menu had no cancelled, skipped, or blocked entries.
  - GREEN: model and integration tests pass after adding the missing statuses to
    the status list contract.
  - Evaluation: users can now filter for every aggregate status the local model
    can produce.
  - Files: `models/actions/run_list_test.go`,
    `tests/integration/actions_status_test.go`.

- [x] Task: Implement Actions status aggregation and filter visibility.
  - SDIE mapping: Implementation / Development.
  - Acceptance: aggregation returns the correct status for every local status
    class, and status filters expose every aggregate status except unknown.
  - RED: tests failed before implementation as described above.
  - GREEN: `aggregateJobStatus` now handles skipped, cancelled, blocked,
    waiting, running, failure, success, and unknown fallback; `GetStatusInfoList`
    exposes all non-unknown aggregate statuses.
  - Evaluation: implementation stays in `models/actions`; routers, templates,
    and Vue continue consuming status strings.
  - Files: `models/actions/run_job.go`, `models/actions/run_list.go`.

- [x] Task: Polish Actions status filter UI and add browser coverage.
  - SDIE mapping: Implementation / Testing / Evaluation.
  - Acceptance: the Actions status filter renders each aggregate status with
    the existing status icon partial and localized label.
  - RED: integration UI assertion failed when menu entries were raw lowercase
    status strings and had no icons.
  - GREEN: template now renders icon plus localized text; integration and E2E
    checks pass.
  - Evaluation: the menu now matches the visual language of Actions run status
    rows and remains keyboard/search-dropdown compatible.
  - Files: `templates/repo/actions/list.tmpl`,
    `templates/repo/actions/status.tmpl`,
    `tests/e2e/actions-status-filter.test.e2e.js`,
    `tests/integration/actions_status_test.go`.

- [x] Task: Evaluate the slice.
  - SDIE mapping: Evaluation / Validation and Verification.
  - Acceptance: UT, IT, E2E classification, build, and caveats are recorded.
  - RED: target tests failed before the model changes.
  - GREEN: targeted UT, affected package UT, targeted IT, build, and diff check
    pass.
  - Evaluation: validation maps back to the three user-visible failures and all
    acceptance criteria in `tasks/actions-status-spec.md`.
  - Files: `tasks/actions-status-todo.md`.

## Evaluation Results

- Verification:
  - UT targeted RED:
    `PATH="/Users/summer/.asdf/shims:$PATH" TAGS="sqlite sqlite_unlock_notify" go test -tags "sqlite sqlite_unlock_notify" ./models/actions -run 'TestAggregateJobStatus|TestGetStatusInfoListIncludesAggregateStatuses' -count=1`
    failed before implementation, reproducing the skipped, cancelled, blocked,
    waiting, unknown, and filter-list gaps.
  - UT targeted GREEN: the same command passed after implementation.
  - UT package:
    `PATH="/Users/summer/.asdf/shims:$PATH" TAGS="sqlite sqlite_unlock_notify" go test -tags "sqlite sqlite_unlock_notify" ./models/actions -count=1`
    passed.
  - IT targeted RED:
    `PATH="/Users/summer/.asdf/shims:$PATH" TAGS="sqlite sqlite_unlock_notify" make test-sqlite#TestActionsRunStatusFilterOptions`
    failed before implementation because the Actions status menu lacked
    cancelled, skipped, and blocked links.
  - IT targeted GREEN: the same command passed after implementation.
  - E2E targeted:
    `PATH="/Users/summer/.local/bin:/Users/summer/.asdf/shims:$PATH" TAGS="sqlite sqlite_unlock_notify" GITEA_E2E_CHROME_PATH="/Applications/Google Chrome.app/Contents/MacOS/Google Chrome" make test-e2e-sqlite#actions-status-filter`
    passed after using the repo harness Node 22 PATH and replacing
    `networkidle` login waiting with URL/DOM readiness.
  - Browser smoke:
    launched the target repo `./gitea web -c tests/sqlite.ini` on port 3003 and
    verified all seven status menu items through Playwright's Chromium API;
    screenshot saved to `outputs/actions-status-ui-smoke.png`.
  - Build:
    `PATH="/Users/summer/.asdf/shims:$PATH" TAGS="sqlite sqlite_unlock_notify" make build`
    passed.
  - Diff check: `git diff --check` passed.
- Validation:
  - All-skipped workflows aggregate as `skipped` instead of appearing active.
  - Cancelled terminal workflows aggregate as `cancelled` instead of collapsing
    into `failure`.
  - Blocked pending workflows can surface as `blocked` when no waiting/running
    job outranks them.
  - The Actions page status filter exposes cancelled, skipped, and blocked.
  - The Actions page status filter renders all aggregate statuses with icons
    and localized labels.
- Caveats:
  - Full backend suite may still hit the known macOS `dyld: missing LC_UUID load
    command` issue documented in `tasks/todo.md`.
- Harness improvements:
  - Actions status ownership and commands were added to the repo-local harness.
  - E2E verification must use `PATH="/Users/summer/.local/bin:/Users/summer/.asdf/shims:$PATH"`
    on this workstation; `/opt/homebrew/bin/node` can hang during Playwright
    test discovery.
