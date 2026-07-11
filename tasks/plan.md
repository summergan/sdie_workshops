# Plan: Repository Tags Search

## SDIE Phase

Design.

## Architecture

- Existing modules:
  - `models/repo/release.go`: owns release/tag query options and SQL
    conditions.
  - `routers/web/repo/release.go`: owns `/{owner}/{repo}/tags` request parsing,
    pagination, query execution, and template data.
  - `templates/repo/tag/list.tmpl`: owns the server-rendered tags page.
  - `templates/shared/search/*`: owns standard Gitea search input markup.
  - `services/context/pagination.go`: preserves `Keyword` as `q` in pagination
    links.
  - `models/repo/release_test.go`: owns UT verification for model/query
    semantics.
  - `tests/integration/release_test.go`: owns IT verification for route,
    fixture database, and rendered template behavior.
  - `tests/e2e/tag-search.test.e2e.js`: owns E2E verification for browser form
    submission and DOM-visible filtered results.
- New or changed modules:
  - Add `Keyword string` to `FindReleasesOptions`.
  - Add keyword filtering in `FindReleasesOptions.ToConds()`.
  - Read `q` in `TagsList`, pass it to model queries, and expose it to the
    template as `Keyword`.
  - Add a GET search form to the tags template.
- Data/API contracts:
  - No database schema change.
  - No REST API or git protocol contract change.
  - `Keyword` is trimmed before use.
  - Non-empty `Keyword` maps to `lower_tag_name LIKE lower(keyword)`.
  - Filtered list and filtered count use the same query options.
- UI or template contracts:
  - Search form method is `GET`.
  - Input name is `q`.
  - Input value is `.Keyword`.
  - Use existing shared search partials.
  - Show `no_results` only for non-empty search with no matching tags.

## Task Split

- [x] Task: Specify tags search behavior.
  - Acceptance: `tasks/spec.md` defines scope, assumptions, domain semantics,
    BDD criteria, boundaries, and commands.
  - Verify: spec maps directly to the tags page slice in
    `.codex/skills/sdie-harness-loop/references/repo-map.md`.
  - Files: `tasks/spec.md`.

- [x] Task: Add RED model coverage.
  - Acceptance: model test describes case-insensitive partial tag-name matching
    through `FindReleasesOptions.Keyword`.
  - Verify: test fails before the `Keyword` option exists.
  - Files: `models/repo/release_test.go`.

- [x] Task: Add RED web coverage.
  - Acceptance: integration test describes `/tags?q=DELETE`, visible filtered
    rows, and retained search input value.
  - Verify: test fails before route/template support exists.
  - Files: `tests/integration/release_test.go`.

- [x] Task: Add browser E2E coverage.
  - Acceptance: Playwright test opens the tags page, submits `q=DELETE` through
    the search form, and observes only `delete-tag` in browser-rendered output.
  - Verify: `make test-e2e-sqlite#tag-search` passes after route/template
    support exists.
  - Files: `tests/e2e/tag-search.test.e2e.js`.

- [x] Task: Implement model keyword contract.
  - Acceptance: `FindReleasesOptions.Keyword` filters tags by
    `lower_tag_name`.
  - Verify: model keyword test passes.
  - Files: `models/repo/release.go`.

- [x] Task: Implement route and pagination contract.
  - Acceptance: `TagsList` reads `q`, sets `Keyword`, applies it to rows and
    count, and relies on pagination to preserve it.
  - Verify: integration test and affected model package pass.
  - Files: `routers/web/repo/release.go`.

- [x] Task: Implement tags page search UI.
  - Acceptance: tags page renders a GET search box, keeps the submitted keyword,
    and shows no-results feedback for unmatched searches.
  - Verify: integration test passes.
  - Files: `templates/repo/tag/list.tmpl`.

- [x] Task: Evaluate the full slice.
  - Acceptance: targeted model test, tags integration test, affected model
    package test, and build are run; host caveats are recorded.
  - Verify: results captured in `tasks/todo.md`.
  - Files: `tasks/todo.md`.

## Six Cross-Checks

- Business Analysis x Technical Design: the user problem is "find a tag by tag
  name"; the design maps that directly to `FindReleasesOptions.Keyword` over
  `lower_tag_name`, without broadening scope to commits, release notes, or API
  search.
- Business Analysis x Engineering Implementation: the vertical slice spans
  model query, web route, template form, and integration test, proving the
  requirement is feasible inside existing Gitea boundaries.
- Business Analysis x Quality Verification: BDD criteria are represented by the
  UT keyword test, IT `/tags?q=DELETE` route/template test, E2E browser search
  flow, and by the empty-search/no-results specification.
- Technical Design x Engineering Implementation: implementation honors module
  boundaries: model owns SQL conditions, route owns request state and count,
  template owns markup, and existing pagination owns query preservation.
- Technical Design x Quality Verification: the model test verifies the
  lower-case query contract and count behavior; the integration test verifies
  route/template behavior as an observable user flow.
- Engineering Implementation x Quality Verification: targeted UT, IT, E2E,
  affected package test, and build are the quality gates for this slice; full
  backend caveat is tracked separately as an environment issue.

## Risks

- Risk: pagination shows wrong totals if the route keeps using unfiltered
  `NumTags`.
  - Mitigation: count with the same `FindReleasesOptions` used to fetch rows.
- Risk: case-insensitive matching differs across databases.
  - Mitigation: query the existing `lower_tag_name` column with lower-cased
    input instead of depending on database collation.
- Risk: template markup drifts from Gitea UI conventions.
  - Mitigation: reuse `templates/shared/search/combo`.
- Risk: a future agent treats search as frontend-only.
  - Mitigation: `AGENTS.md`, `repo-map.md`, and this plan state server-side
    search as a hard boundary.

## Loop-Back Triggers

- Return to Specification when: users need search over commit hash, release
  notes, annotated tag messages, API responses, or multiple filter fields.
- Return to Design when: database performance requires an index, query planner
  changes, or pagination/search contracts need to move into a shared helper.
- Improve harness when: a repeated ambiguity appears around route/query naming,
  template partial choice, or how to record verification evidence.
