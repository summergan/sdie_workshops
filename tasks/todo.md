# Todo: Repository Tags Search

## Implementation Tasks

- [x] Task: Build the repo-local SDIE harness.
  - SDIE mapping: Specification/Design support.
  - Acceptance: `AGENTS.md`, harness references, templates, and task artifacts
    exist and define the local loop.
  - RED: earlier task flow lacked repo-local harness instructions.
  - GREEN: `AGENTS.md` and `.codex/skills/sdie-harness-loop/*` define the SDIE loop,
    references, and templates.
  - Evaluation: harness now requires spec, design, implementation, evaluation,
    cross-checks, and steering-loop updates for non-trivial changes.
  - Files: `AGENTS.md`, `.codex/skills/sdie-harness-loop/*`.

- [x] Task: Complete the SDIE specification for tags search.
  - SDIE mapping: Specification.
  - Acceptance: spec includes assumptions, business context, structured
    requirements, domain semantics, BDD criteria, commands, and boundaries.
  - RED: old spec captured feature intent but did not include SDIE phase output
    structure or explicit cross-domain semantics.
  - GREEN: `tasks/spec.md` now follows
    `.codex/skills/sdie-harness-loop/assets/templates/sdie-spec.md`.
  - Evaluation: spec can drive design and test assertions without relying on
    conversational memory.
  - Files: `tasks/spec.md`.

- [x] Task: Complete the SDIE design plan.
  - SDIE mapping: Design.
  - Acceptance: plan includes architecture, data/template contracts, MECE task
    split, risks, six cross-checks, and loop-back triggers.
  - RED: old plan listed components and risks but did not expose SDIE
    cross-check evidence.
  - GREEN: `tasks/plan.md` now follows
    `.codex/skills/sdie-harness-loop/assets/templates/sdie-plan.md`.
  - Evaluation: design states exactly where model, route, template, and tests
    own behavior.
  - Files: `tasks/plan.md`.

- [x] Task: Write RED model coverage.
  - SDIE mapping: Implementation / Testing.
  - Acceptance: model test describes case-insensitive partial tag-name matching
    through `FindReleasesOptions.Keyword`.
  - RED: test failed first because `FindReleasesOptions.Keyword` did not exist.
  - GREEN: targeted model test passes after adding the keyword query contract.
  - Evaluation: protects the database query semantics independently from the
    web route.
  - Files: `models/repo/release_test.go`.

- [x] Task: Write RED integration coverage.
  - SDIE mapping: Implementation / Testing.
  - Acceptance: integration test describes `/tags?q=DELETE`, visible filtered
    rows, retained search input value, no-results state, and pagination links
    preserving `q`.
  - RED: test failed first because `q=DELETE` returned unfiltered tags and no
    search input preserved the value.
  - GREEN: tags integration slice passes after route and template wiring.
  - Evaluation: protects the user-visible behavior at HTTP/template level.
  - Files: `tests/integration/release_test.go`.

- [x] Task: Write browser E2E coverage.
  - SDIE mapping: Implementation / Testing.
  - Acceptance: Playwright opens the tags page, submits `DELETE` through the
    search form, verifies `input[name="q"]` retains `DELETE`, and verifies the
    rendered tag list contains only `delete-tag`.
  - RED: not applicable as a historical RED for this late-added layer; the
    harness improvement is to add E2E before declaring future browser-facing
    slices complete.
  - GREEN: local Chrome E2E passed with Node 22 and
    `GITEA_E2E_CHROME_PATH`.
  - Evaluation: protects the real browser flow beyond UT and IT.
  - Files: `tests/e2e/tag-search.test.e2e.js`.

- [x] Task: Implement tags search behavior.
  - SDIE mapping: Implementation / Development.
  - Acceptance: `q` filters tags by case-insensitive partial tag name,
    pagination count uses the filtered result, and empty `q` keeps old behavior.
  - RED: model and integration tests failed before implementation.
  - GREEN: model, route, and template changes satisfy both tests.
  - Evaluation: implementation follows existing package boundaries and shared UI
    partial conventions.
  - Files: `models/repo/release.go`, `routers/web/repo/release.go`,
    `templates/repo/tag/list.tmpl`.

- [x] Task: Evaluate the slice.
  - SDIE mapping: Evaluation / Validation and Verification.
  - Acceptance: targeted tests and build pass, and unrelated host caveats are
    recorded.
  - RED: full backend suite is not a reliable local quality gate on this host
    because generated helper binaries can fail with `dyld: missing LC_UUID load
    command`.
  - GREEN: narrow behavior tests and build pass.
  - Evaluation: verification evidence is recorded below; validation maps back to
    the BDD criteria in `tasks/spec.md`.
  - Files: `tasks/todo.md`.

## Evaluation Results

- Verification:
  - UT targeted: `TAGS="sqlite sqlite_unlock_notify" go test -tags "sqlite sqlite_unlock_notify" ./models/repo -run TestFindReleasesOptionsKeyword -count=1`
    passed.
  - UT package: `TAGS="sqlite sqlite_unlock_notify" go test -tags "sqlite sqlite_unlock_notify" ./models/repo -count=1`
    passed.
  - IT targeted: `TAGS="sqlite sqlite_unlock_notify" make test-sqlite#TestViewTagsListSearch`
    passed, including filtered results, no-results state, retained input value,
    and pagination links preserving `q`.
  - E2E targeted: `PATH="/Users/summer/.local/bin:/Users/summer/.asdf/shims:$PATH" GITEA_E2E_CHROME_PATH="/Applications/Google Chrome.app/Contents/MacOS/Google Chrome" TAGS="sqlite sqlite_unlock_notify" make test-e2e-sqlite#tag-search`
    passed with `1 passed (2.5s)`.
  - E2E test file lint: `npx eslint tests/e2e/tag-search.test.e2e.js` passed.
  - Build: `TAGS="sqlite sqlite_unlock_notify" make build` passed.
- Validation:
  - The implemented slice solves the original user pain: quickly finding a tag
    on a repository with many tags.
  - Search remains shareable/bookmarkable through `q`.
  - Existing tags page links, permissions, and actions remain owned by the
    original template and route logic.
- Caveats:
  - Full `TAGS="sqlite sqlite_unlock_notify" make test-backend` still fails on
    this macOS host with unrelated `dyld: missing LC_UUID load command` errors
    in multiple packages.
  - Bindata generation can hit the same host/toolchain issue; dynamic-template
    local verification is the reliable local path here.
  - E2E should use Node 22 from `/Users/summer/.local/bin` on this workstation.
    Node 26 from Homebrew hung during Playwright test discovery.
  - Playwright-managed Chromium build v1112 was not downloadable reliably in
    this run, so the passing E2E path uses local Google Chrome instead.
- Harness improvements:
  - SDIE from the workshop material has been folded into the repo harness.
  - Repo-local skill has been normalized to
    `.codex/skills/sdie-harness-loop/SKILL.md`, with references in
    `.codex/skills/sdie-harness-loop/references/` and templates in
    `.codex/skills/sdie-harness-loop/assets/templates/`.
  - The current requirement has been upgraded from a simple SDD/TDD task list to
    Specification, Design, Implementation, and Evaluation artifacts with six
    cross-checks.
  - The local skill now requires UT, IT, and E2E classification before
    verification is reported.
