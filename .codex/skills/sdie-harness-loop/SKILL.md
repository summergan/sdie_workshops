---
name: sdie-harness-loop
description: Use for non-trivial feature work in this Gitea workshop repo when applying SDIE, maintaining spec/plan/todo artifacts, and verifying changes with UT, IT, E2E, UI smoke, and build evidence.
---

# SDIE Harness Loop

Use this for feature work in this repository.

## References And Assets

- Always load `references/sdie.md` before planning a non-trivial change.
- Load `references/repo-map.md` before editing source files or choosing package
  ownership.
- Load `references/testing.md` before classifying or running UT, IT, E2E, or
  build checks.
- Load `references/ui-validation.md` whenever a change affects a rendered page,
  template, browser-visible text, icon, state, filter, form, navigation, or
  data feeding one of those surfaces.
- Use `assets/templates/sdie-spec.md` for `tasks/spec.md`.
- Use `assets/templates/sdie-plan.md` for `tasks/plan.md`.
- Use `assets/templates/sdie-todo.md` for `tasks/todo.md`.

## Instructions

1. Load `references/sdie.md` before planning a non-trivial change.
2. Start Specification in `tasks/spec.md`: assumptions, business context,
   structured requirements, UI impact, acceptance criteria, and validation
   notes.
3. Continue Design in `tasks/plan.md`: architecture, API/data contracts,
   UI/template contracts, task split, risks, and SDIE cross-checks.
4. Convert design into ordered MECE tasks in `tasks/todo.md`; each task has one
   responsibility and one verification path.
5. Before editing implementation code, add a RED test that fails on current
   behavior.
6. Classify tests before running them:
   - UT: package-level unit/model tests for pure logic or query semantics.
   - IT: integration tests for Go HTTP routes, database fixtures, and rendered
     templates.
   - E2E: Playwright browser tests for critical user flows, form behavior, and
     DOM-visible results.
7. Apply the UI validation gate before implementation when the change has any
   user-visible surface. A logic fix that changes visible status, filtering,
   labels, icons, or layout is a UI-impacting change.
8. Implement the narrowest vertical slice through the relevant repo layers.
9. Evaluate both verification ("did we build it right?") and validation ("did we
   build the right thing?").
10. Update the task list and harness notes as each checkpoint completes.

## Phase Outputs

- Specification outputs: `proposal.md` or `tasks/spec.md`, domain model notes,
  BDD/acceptance criteria.
- Design outputs: `tasks/plan.md`, module boundaries, API/data contracts,
  ordered tasks.
- Implementation outputs: verifiable code changes plus UT, IT, and E2E coverage
  when the feature crosses those layers.
- Evaluation outputs: V&V notes, quality gate results grouped by UT, IT, E2E,
  UI smoke, build, and known caveats.

## Test Layer Policy

- Start with UT when behavior can be proven without a web server or browser.
- Add IT when behavior crosses model, router, database fixture, or template
  boundaries.
- Add E2E when the user's success depends on real browser behavior, form
  submission, navigation, frontend JavaScript, or responsive DOM output.
- Add UI smoke verification whenever a change affects rendered HTML, templates,
  browser-visible data, icons, labels, filters, status colors, or layout. UI
  smoke must compare against the existing local UI pattern and capture evidence
  such as a Chrome screenshot, DOM assertion, or both.
- Run the smallest relevant command at each layer first, then broaden only after
  the targeted command passes.
- Record skipped or blocked layers explicitly; do not imply E2E passed when only
  UT/IT ran.
- Do not claim "logic only" until the spec names the affected user-visible
  pages or explicitly states why no rendered surface changes.

## Stop And Loop Back

- If requirements are ambiguous during Design, return to Specification.
- If implementation exposes an infeasible design, return to Design.
- If verification passes but validation fails, return to Specification.
- If UT/IT/E2E pass but the UI impact was not declared or visually checked,
  return to Design and add the missing UI contract.
- If the same failure mode repeats twice, improve the harness before continuing.

## Output Expectations

- The final answer must list the spec, tests, implementation files, and
  verification commands.
- The final answer must state which SDIE cross-checks were satisfied or which
  remain intentionally out of scope.
- The final answer must group verification evidence by UT, IT, E2E, UI smoke,
  and build.
- For UI-impacting changes, the final answer must include UI smoke evidence or
  explicitly state that UI validation is incomplete.
- Mention any known environment-only failures separately from product failures.
