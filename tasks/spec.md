# Spec: Repository Tags Search

## SDIE Phase

Specification.

## Assumptions

1. The requested feature is for the server-rendered repository tags page at
   `/{owner}/{repo}/tags`, not for the REST API, git protocol, or release page.
2. Users search by tag name only. Commit hash, message, author, and release
   metadata search are out of scope for this slice.
3. Matching is case-insensitive and partial, because the existing release/tag
   table stores both `tag_name` and `lower_tag_name`.
4. Search is server-side. Repositories with many tags should not require loading
   the full tag list into browser-side JavaScript before filtering.
5. Search state uses query parameter `q`, matching existing Gitea search and
   pagination conventions.
6. Empty or whitespace-only `q` is equivalent to the current unfiltered tags
   page.

## Business Context

- User/problem: repository users and maintainers need to quickly find a specific
  version tag when a repository contains many tags.
- Current pain: the tags page only lists paginated tags; users must browse page
  by page or use browser find on the current page, which does not work across
  pagination.
- Constraint-heavy brownfield context: this is a mature Gitea fork with
  server-rendered Go templates, XORM model queries, and existing pagination
  contracts. The change should fit those boundaries instead of introducing a new
  client-side search subsystem.
- Success outcome: users can type a tag-name fragment, submit the search, and
  receive a paginated list containing only matching tags while all existing tag
  actions and links still work.

## Structured Requirements

### Added

- Add a search input to the repository tags page.
- Submitting the form sends `GET /{owner}/{repo}/tags?q=<keyword>`.
- Filter tags by case-insensitive partial match against tag name.
- Preserve the submitted keyword in the search input after the page reloads.
- Show an empty-results state when a non-empty keyword has no matches.

### Modified

- Tags pagination must count and paginate the filtered result set when `q` is
  present.
- Tags pagination links must preserve the active `q` value.
- The model query options used by tags listing must accept a keyword filter.

### Removed

- None. Existing tags page behavior remains unchanged when `q` is empty.

## Domain Semantics

- Domain terms:
  - Tag: a repository git tag displayed through Gitea release/tag records.
  - Keyword: user-provided tag-name fragment from query parameter `q`.
  - Match: case-insensitive substring match against the tag name.
- Existing repo concepts involved:
  - `FindReleasesOptions` expresses release/tag query filters.
  - `IsTag: true` identifies tag rows rather than release rows.
  - `LowerTagName` / `lower_tag_name` provides database-neutral
    case-insensitive matching.
  - `Keyword` template data maps to existing shared search partials and
    pagination query preservation.
- Boundary cases:
  - `q` omitted: current page behavior and count are unchanged.
  - `q` is whitespace: trim and treat as empty.
  - `q` differs in case from the tag name: still match.
  - `q` has no matches: show no tag rows and display no-results feedback.
  - User lacks existing permissions for an action: keep existing permission
    checks; search must not widen or narrow authorization.

## Acceptance Criteria / BDD

- Given a repository has tags `v1.0.0`, `delete-tag`, and `signed-tag`, when a
  user opens `/user2/repo1/tags?q=DELETE`, then only `delete-tag` is listed.
- Given a user searches for `DELETE`, when the tags page renders, then the
  search input value remains `DELETE`.
- Given `q` is empty, when a user opens the tags page, then the unfiltered tag
  list and existing actions render as before.
- Given a search matches more tags than one page can show, when the user moves
  through pagination, then pagination links preserve `q` and counts reflect only
  matching tags.
- Given a search has no matching tags, when the page renders, then no tag rows
  are shown and the page displays a no-results state instead of falling back to
  the full list.

## Commands

- Narrow model test:
  `TAGS="sqlite sqlite_unlock_notify" go test -tags "sqlite sqlite_unlock_notify" ./models/repo -run TestFindReleasesOptionsKeyword -count=1`
- Tags integration test:
  `TAGS="sqlite sqlite_unlock_notify" make test-sqlite#TestViewTagsListSearch`
- Tags E2E test:
  `PATH="/Users/summer/.local/bin:/Users/summer/.asdf/shims:$PATH" GITEA_E2E_CHROME_PATH="/Applications/Google Chrome.app/Contents/MacOS/Google Chrome" TAGS="sqlite sqlite_unlock_notify" make test-e2e-sqlite#tag-search`
- Affected model package:
  `TAGS="sqlite sqlite_unlock_notify" go test -tags "sqlite sqlite_unlock_notify" ./models/repo -count=1`
- Build:
  `TAGS="sqlite sqlite_unlock_notify" make build`

## Boundaries

- Always: keep search server-side and encoded in the URL as `q`.
- Always: preserve existing route ownership, permission checks, tag links,
  deletion behavior, and pagination style.
- Always: use existing model/router/template boundaries and shared search
  partials.
- Always: write or update a failing test before changing behavior.
- Ask first: database schema changes, indexes, new dependencies, REST API
  changes, or frontend build pipeline changes.
- Never: hide tags by a new authorization rule, load every tag into JavaScript
  for filtering, or remove existing release/tag tests to make the slice pass.

## Open Questions

- None for this slice.
