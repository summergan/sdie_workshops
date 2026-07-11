# Testing Reference

## Test Layers

- UT: package-level Go tests that verify isolated model/query behavior without a
  browser. Prefer these for SQL condition semantics and data contracts.
- IT: Go integration tests under `tests/integration` that exercise HTTP routes,
  database fixtures, permissions, and rendered templates.
- E2E: Playwright tests under `tests/e2e` that exercise real browser behavior,
  form submission, navigation, and DOM-visible outcomes.

## Narrow Commands

- UT, model search behavior:
  `TAGS="sqlite sqlite_unlock_notify" go test -tags "sqlite sqlite_unlock_notify" ./models/repo -run TestFindReleasesOptionsKeyword -count=1`
- IT, tags page behavior:
  `TAGS="sqlite sqlite_unlock_notify" make test-sqlite#TestViewTagsListSearch`
- E2E, tags browser flow:
  `PATH="/Users/summer/.local/bin:/Users/summer/.asdf/shims:$PATH" GITEA_E2E_CHROME_PATH="/Applications/Google Chrome.app/Contents/MacOS/Google Chrome" TAGS="sqlite sqlite_unlock_notify" make test-e2e-sqlite#tag-search`

## Broader Commands

- UT, affected model package:
  `TAGS="sqlite sqlite_unlock_notify" go test -tags "sqlite sqlite_unlock_notify" ./models/repo -count=1`
- IT, full SQLite integration suite:
  `TAGS="sqlite sqlite_unlock_notify" make test-sqlite`
- E2E, full SQLite browser suite:
  `TAGS="sqlite sqlite_unlock_notify" make test-e2e-sqlite`
- Build:
  `TAGS="sqlite sqlite_unlock_notify" make build`
- Full backend:
  `TAGS="sqlite sqlite_unlock_notify" make test-backend`

## Caveats

- On this macOS setup, the full backend suite has previously failed with a
  `dyld: missing LC_UUID load command` environment issue. Treat that as an
  environment caveat only after the targeted package and integration tests pass.
- The same dyld issue can prevent `go generate -tags bindata ./modules/templates`
  from refreshing embedded templates. Use dynamic-template tests for local
  verification, then regenerate bindata with a compatible Go/macOS toolchain
  before a release build that requires embedded templates.
- Integration tests compile a test binary first, so the first run can be slow.
- E2E tests run Playwright through `npx`; verify `npx` exists before running
  them and expect the first run to install browser dependencies.
- If `go` is installed through asdf but is not on the shell PATH, prefix test
  commands with `PATH="/Users/summer/.asdf/shims:$PATH"` on this workstation.
- To avoid downloading Playwright-managed browsers, set
  `GITEA_E2E_CHROME_PATH="/Applications/Google Chrome.app/Contents/MacOS/Google Chrome"`.
  The Makefile skips `npx playwright install` when this variable is set, and
  `playwright.config.js` restricts the run to the local Chrome-backed chromium
  project.
- Prefer `PATH="/Users/summer/.local/bin:/Users/summer/.asdf/shims:$PATH"` for
  E2E on this workstation. `/Users/summer/.local/bin/node` is Node 22.22.3 and
  works with Playwright 1.43.1; `/opt/homebrew/bin/node` is Node 26.0.0 and can
  hang during Playwright test discovery.
- If Playwright browser download stalls on Azure CDN, retry once with
  `PLAYWRIGHT_DOWNLOAD_HOST="https://npmmirror.com/mirrors/playwright"`. If it
  still stalls before tests start, record the E2E layer as dependency-download
  blocked rather than as a product failure.
