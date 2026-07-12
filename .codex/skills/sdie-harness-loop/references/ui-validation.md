# UI Validation Reference

Use this reference whenever a change affects a rendered page, template,
browser-visible data, icon, status, filter, form, menu, navigation, or layout.

## UI Impact Triage

Classify a change as UI-impacting when any of these are true:

- A route, model, API, or query changes data shown in a browser page.
- A template, locale string, icon partial, CSS class, frontend script, or E2E
  selector changes.
- The user problem mentions finding, filtering, seeing, clicking, scanning, or
  recognizing something on a page.
- A backend state maps to browser-visible text, icon, color, menu option, badge,
  empty state, error state, or disabled/enabled control.

Do not call a change "logic only" until the spec names the affected pages and
explains why no browser-visible behavior changes.

## Required UI Contract

Before implementation, write the UI contract in `tasks/plan.md`:

- Surface: exact route/page and component or template.
- Existing pattern: nearby Gitea page, template partial, icon, class, or locale
  key to reuse.
- Visual states: every state the user can see, including empty, error,
  disabled, loading, selected, and filtered variants when relevant.
- Copy/localization: visible labels and locale keys.
- Interaction: click, keyboard, focus, filter, sort, navigation, and persistence
  behavior when relevant.
- Responsive/accessibility: whether layout, icon-only controls, aria labels, or
  title/tooltips need coverage.

## Required Verification

For every UI-impacting change:

- Add at least one browser-visible assertion in IT or E2E.
- Add UI smoke evidence from Chrome or Playwright when the change affects visual
  composition, status/icon/color conventions, or template layout.
- Compare against the existing local UI pattern before declaring completion.
- Record the screenshot path, DOM assertion, or reason the UI smoke layer is
  blocked.

## Anti-Regression Checks

Use this checklist before finalizing:

- Status-like values use the same icon, color, label, and localization pattern
  as adjacent Gitea status UI.
- Menu/filter options are not text-only when the existing page uses icons or
  visual status affordances.
- New states are visible in all places where users search, filter, summarize, or
  scan the affected entity.
- E2E asserts visible affordances, not only URL/query/data changes.
- Screenshots show the changed UI, not just a successful page load.
- The final answer separates UT, IT, E2E, UI smoke, build, and caveats.

## Completion Rule

If a user-visible surface changed and there is no UI contract plus UI smoke
evidence, mark validation incomplete and loop back to Design. Do not claim the
requirement is fully complete.
