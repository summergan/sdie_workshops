# Plan: Gitea Actions Run Status Aggregation

## SDIE Phase

Design.

## Architecture

- Existing modules:
  - `models/actions/status.go`: status enum and terminal/pending helpers.
  - `models/actions/run_job.go`: owns aggregation after job status updates.
  - `models/actions/run_list.go`: owns Actions run-list status filter entries.
  - `routers/web/repo/actions/actions.go`: renders the Actions run list and
    applies status filters.
  - `templates/repo/actions/status.tmpl`: server-rendered status icon.
  - `web_src/js/components/ActionRunStatus.vue`: Vue-rendered status icon.
- New or changed modules:
  - Add model tests for aggregation precedence and status filter entries.
  - Add integration coverage for the Actions page status filter menu.
  - Add E2E coverage for the browser-visible Actions status filter menu.
  - Update `aggregateJobStatus` to cover all local statuses.
  - Update `GetStatusInfoList` to include all aggregate statuses.
- Data/API contracts:
  - No database schema change.
  - No REST API contract change.
  - `ActionRun.Status` remains the single persisted aggregate.
  - `StatusUnknown` remains fallback only and is not added to the filter menu.
- UI or template contracts:
  - Existing status icon template/Vue component render the status string.
  - Actions list filter menu is driven by `GetStatusInfoList`.
  - Localized labels come from existing `actions.status.*` locale keys.
  - Server-rendered filter items reuse the status icon partial and localized
    labels so list filters match the rest of the Actions UI.

## Task Split

- [x] Task: Specify Actions status behavior.
  - Acceptance: `tasks/actions-status-spec.md` captures semantics, BDD, and
    boundaries.
  - Verify: spec maps to Actions code in the updated repo map.
  - Files: `tasks/actions-status-spec.md`,
    `.codex/skills/sdie-harness-loop/references/repo-map.md`.

- [x] Task: Add RED model coverage.
  - Acceptance: tests fail on current aggregation for all-skipped, cancelled,
    and blocked cases.
  - Verify: targeted Actions model test fails before implementation.
  - Files: `models/actions/run_job_test.go`.

- [x] Task: Add RED status filter coverage.
  - Acceptance: test fails because `GetStatusInfoList` and the rendered Actions
    filter do not include cancelled/skipped/blocked.
  - Verify: targeted model/integration tests fail before implementation.
  - Files: `models/actions/run_list_test.go`,
    `tests/integration/actions_status_test.go`.

- [x] Task: Implement aggregation contract.
  - Acceptance: aggregate status handles success/skipped, all skipped, running,
    waiting, blocked, cancelled, failure, and unknown fallback.
  - Verify: targeted Actions model tests pass.
  - Files: `models/actions/run_job.go`.

- [x] Task: Implement filter visibility contract.
  - Acceptance: `GetStatusInfoList` returns success, failure, cancelled,
    skipped, waiting, running, and blocked.
  - Verify: model and integration tests pass.
  - Files: `models/actions/run_list.go`.

- [x] Task: Polish Actions status filter UI.
  - Acceptance: status filter entries render an icon plus localized label for
    every aggregate status.
  - Verify: integration HTML assertion and browser E2E pass.
  - Files: `templates/repo/actions/list.tmpl`,
    `templates/repo/actions/status.tmpl`,
    `tests/e2e/actions-status-filter.test.e2e.js`.

- [x] Task: Evaluate and record evidence.
  - Acceptance: UT, IT, E2E, build, and caveats are recorded.
  - Verify: `tasks/actions-status-todo.md` contains command results.
  - Files: `tasks/actions-status-todo.md`.

## Six Cross-Checks

- Business Analysis x Technical Design: the user-visible failures map directly
  to model aggregation and status filter exposure, not to a new UI subsystem.
- Business Analysis x Engineering Implementation: the vertical slice changes
  the model output that the existing list/detail pages already render.
- Business Analysis x Quality Verification: BDD cases are encoded as unit
  tests, an integration test for rendered filter options, and a browser E2E
  for visible UI behavior.
- Technical Design x Engineering Implementation: implementation stays in
  `models/actions`; routers and templates continue to consume status data.
- Technical Design x Quality Verification: tests assert the model contract, the
  web-rendered filter contract, and the browser-visible dropdown contract.
- Engineering Implementation x Quality Verification: targeted UT, IT, E2E, and
  build form the quality gate for this UI-visible slice.

## Risks

- Risk: aggregation precedence differs from upstream behavior.
  - Mitigation: align with current Gitea main for the subset of statuses present
    in this branch.
- Risk: adding status filter entries exposes a status the aggregator cannot
  produce.
  - Mitigation: keep the filter list in the same order and subset as the
    aggregator, excluding `unknown`.
- Risk: a blocked pending run is treated as terminal.
  - Mitigation: keep `StatusBlocked` outside `IsDone`; only aggregation exposes
    it when no waiting/running job exists.
- Risk: integration test mutates shared fixtures.
  - Mitigation: use `tests.PrepareTestEnv(t)` and runtime unit updates instead
    of editing fixture YAML.

## Loop-Back Triggers

- Return to Specification when: users require new states, `continue-on-error`,
  or commit-status behavior changes.
- Return to Design when: runner protocol or task scheduling semantics conflict
  with the aggregate precedence.
- Improve harness when: a new feature starts in a code area missing from
  `repo-map.md` or `testing.md`.
