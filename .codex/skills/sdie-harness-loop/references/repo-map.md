# Repo Map

## Architecture

- `main.go`: process entrypoint.
- `cmd/web.go`: web command bootstrap and installer switch.
- `routers/web`: server-rendered web routes.
- `routers/api`: REST API routes.
- `models`: database models and query options.
- `services`: business workflows that coordinate models, git, and side effects.
- `templates`: Go templates for HTML views.
- `web_src`: JavaScript and CSS source compiled into public assets.
- `tests/integration`: real HTTP integration tests.

## Tags Page Slice

- `routers/web/repo/release.go`: `TagsList` handles `/{owner}/{repo}/tags`.
- `models/repo/release.go`: `FindReleasesOptions` builds release/tag SQL
  conditions used by the tags page.
- `templates/repo/tag/list.tmpl`: renders the tags table.
- `templates/shared/search/*`: existing search input and button partials.
- `services/context/pagination.go`: `SetDefaultParams` preserves `Keyword` as
  the `q` query parameter during pagination.
- `tests/integration/release_test.go`: web behavior tests for releases and tags.
- `models/repo/release_test.go`: model-level release query tests.

## Pattern Notes

- URL search state uses `q` and is exposed to templates as `Keyword`.
- Pagination links preserve `Keyword` automatically.
- Release/tag rows store both `TagName` and `LowerTagName`; use the lower-case
  column for case-insensitive tag search.
