# Project Agent Guide

## Project Context

This repository is a Gitea fork used for SDIE workshop work. It is a Go web
application with server-rendered templates, a small JavaScript frontend layer,
and integration tests that exercise real HTTP routes.

## Required SDIE Harness Loop

Use the local SDIE harness for any non-trivial change. SDIE means
Specification, Design, Implementation, and Evaluation. It is the controlling
method for changing mature software; TDD is one verification mechanism inside
Implementation and Evaluation.

1. Load `.codex/skills/sdie-harness-loop/SKILL.md`.
2. Load `.codex/skills/sdie-harness-loop/references/sdie.md` and the relevant
   repo reference.
3. Specification: update `tasks/spec.md` with structured requirements, business
   context, acceptance criteria, and BDD-style validation notes.
4. Design: update `tasks/plan.md` with architecture, contracts, task split,
   risks, and the six SDIE cross-checks that apply.
5. Implementation: write the RED test first, prove it fails, then implement the
   smallest GREEN change.
6. Evaluation: run verification and validation checks, record results in
   `tasks/todo.md`, and loop back to Spec or Design when a cross-check fails.
7. Steering loop: when the same issue appears twice, improve the harness before
   continuing the same style of work.

## Commands

- Build: `TAGS="sqlite sqlite_unlock_notify" make build`
- Backend tests: `TAGS="sqlite sqlite_unlock_notify" make test-backend`
- Model test slice: `TAGS="sqlite sqlite_unlock_notify" go test -tags "sqlite sqlite_unlock_notify" ./models/repo -run TestFindReleasesOptionsKeyword -count=1`
- Tags integration slice: `TAGS="sqlite sqlite_unlock_notify" make test-sqlite#TestViewTagsListSearch`
- Tags E2E slice with local Chrome:
  `PATH="/Users/summer/.local/bin:/Users/summer/.asdf/shims:$PATH" GITEA_E2E_CHROME_PATH="/Applications/Google Chrome.app/Contents/MacOS/Google Chrome" TAGS="sqlite sqlite_unlock_notify" make test-e2e-sqlite#tag-search`
- Actions status UT slice:
  `PATH="/Users/summer/.asdf/shims:$PATH" TAGS="sqlite sqlite_unlock_notify" go test -tags "sqlite sqlite_unlock_notify" ./models/actions -run 'TestAggregateJobStatus|TestGetStatusInfoListIncludesAggregateStatuses' -count=1`
- Actions status IT slice:
  `PATH="/Users/summer/.asdf/shims:$PATH" TAGS="sqlite sqlite_unlock_notify" make test-sqlite#TestActionsRunStatusFilterOptions`
- Actions status E2E slice with local Chrome:
  `PATH="/Users/summer/.local/bin:/Users/summer/.asdf/shims:$PATH" GITEA_E2E_CHROME_PATH="/Applications/Google Chrome.app/Contents/MacOS/Google Chrome" TAGS="sqlite sqlite_unlock_notify" make test-e2e-sqlite#actions-status-filter`
- Dev server: `./gitea web -c custom/conf/app.ini`

Use the local Go version declared in `.tool-versions` when available.
On this macOS host, Go 1.22.12-built helper binaries can fail with
`dyld: missing LC_UUID load command`; prefer dynamic-template local tests unless
the bindata generator is known to run.
For E2E, prefer the local Node 22 binary in `/Users/summer/.local/bin` plus
local Chrome via `GITEA_E2E_CHROME_PATH`; Homebrew Node 26 can hang Playwright
1.43's test runner on this workstation.

## Code Conventions

- Follow existing package boundaries: `models` for database query behavior,
  `routers` for HTTP request handling, `templates` for server-rendered UI.
- Prefer existing template partials such as `shared/search/*` over new markup.
- Keep search/filter state in URL query parameters so pages are shareable.
- Keep tests state-based: assert visible page output or returned model rows.
- Classify behavior checks as UT, IT, and E2E before execution; record each
  layer separately in `tasks/todo.md`.
- Split work MECE: each task should have one responsibility, clear ownership,
  and no overlap with neighboring tasks.
- Treat humans as reviewers and AI as producer: AI may draft artifacts and code,
  but the harness must expose assumptions, contracts, and verification evidence
  for human judgment.

## Boundaries

- Always: write a spec and a failing test before behavior changes.
- Always: keep SDIE artifacts current before implementation continues.
- Always: define Done through cross-check evidence, not through code completion
  alone.
- Always: preserve user work and inspect `git status --short` before editing.
- Always: run the narrow test that protects the changed behavior.
- Ask first: database schema changes, new dependencies, CI changes.
- Never: commit secrets, delete fixtures to make tests pass, or skip failing tests.
