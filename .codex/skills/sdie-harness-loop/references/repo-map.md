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

## Actions Run Status Slice

- `models/actions/status.go`: canonical Actions status enum and helpers.
- `models/actions/run_job.go`: job status updates and workflow-run status
  aggregation from all jobs in a run.
- `models/actions/run_list.go`: run-list query options and status filter menu
  entries for `/{owner}/{repo}/actions`.
- `routers/web/repo/actions/actions.go`: server-rendered Actions run list,
  workflow/actor/status filtering, and pagination.
- `routers/web/repo/actions/view.go`: JSON payload for the run detail page and
  job list consumed by the Vue action view.
- `templates/repo/actions/runs_list.tmpl`: server-rendered run list status
  icon surface.
- `templates/repo/actions/status.tmpl`: server template for status icons.
- `templates/repo/actions/view.tmpl`: injects localized status strings for the
  Vue action view.
- `web_src/js/components/ActionRunStatus.vue`: Vue status icon component that
  mirrors `templates/repo/actions/status.tmpl`.
- `web_src/js/components/RepoActionView.vue`: Vue run detail page that renders
  run, job, and step statuses from `ViewPost`.
- `tests/integration/actions_status_test.go`: web behavior checks for Actions
  status list/filter visibility.
- `models/actions/*_test.go`: model-level status aggregation checks.

## Actions Status Pattern Notes

- `StatusSuccess`, `StatusFailure`, `StatusCancelled`, and `StatusSkipped` are
  final states (`IsDone`).
- `StatusWaiting`, `StatusRunning`, and `StatusBlocked` are pending states.
- Workflow-run status is derived from job statuses; do not duplicate aggregation
  logic in routers or templates.
- The run-list status filter should include every aggregate status that can be
  returned by the aggregator, excluding `StatusUnknown` fallback.
- The Go template and Vue status icon component are intentionally mirrored; when
  changing one status rendering surface, inspect the other.
