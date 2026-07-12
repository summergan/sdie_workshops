# SDIE Reference

SDIE is the repo's engineering method for mature, brownfield software:
Specification, Design, Implementation, and Evaluation. The goal is to reduce
maintenance cost and lower the uncertainty of agentic engineering in a
constraint-heavy codebase.

## Core Model

- SDIE is not a waterfall. The four domains can overlap, run in parallel, and
  loop back when evidence contradicts an earlier artifact.
- Harness is the execution safety layer across the four domains. It provides
  guidance before an agent acts and sensing after an agent acts.
- Human value is review and taste: judge ambiguity, quality, tradeoffs, and
  long-term fit. AI value is structured production: draft artifacts, code,
  tests, and verification evidence.
- Done is a validation network outcome, not a local task ending.

## Four Domains

| Domain | SDIE phase | Responsibility | Primary output |
| --- | --- | --- | --- |
| Business Analysis | Specification | Understand business context, requirement goals, boundaries, and constraints | Structured spec, domain semantics, BDD criteria |
| Technical Design | Design | Create architecture, module boundaries, data/API contracts, and task plan | Design plan, contracts, tasks |
| Engineering Implementation | Implementation | Produce code, tests, buildable slices, and local runnable artifacts | Verifiable code change plus tests |
| Quality Verification | Evaluation | Verify behavior, contracts, system quality, and delivery risk | V&V notes and quality gate results |

## Phase Components

Each SDIE phase has two components:

- Specification: Specify and Business Context.
- Design: Architecture and Planning.
- Implementation: Development and Testing.
- Evaluation: Validation and Verification.

Validation asks whether the change solves the right user/business problem.
Verification asks whether the implementation satisfies the spec, contracts, and
quality gates.

## Six Cross-Checks

Use these as the definition of done for non-trivial changes:

| Cross-check | Evidence asset | Purpose |
| --- | --- | --- |
| Business Analysis x Technical Design | Domain model or design semantics | Requirements and design structure agree |
| Business Analysis x Engineering Implementation | PoC or vertical slice | Requirement is feasible in this codebase |
| Business Analysis x Quality Verification | BDD criteria or acceptance tests | Business expectation is testable |
| Technical Design x Engineering Implementation | LLD, API contract, or module contract | Implementation obeys design boundaries |
| Technical Design x Quality Verification | Contract tests or design checks | Design can be verified |
| Engineering Implementation x Quality Verification | CI, unit, integration, build checks | Code change remains within quality baseline |

## Work Harness

The smallest work loop is:

1. Input: load the focused spec, design notes, source files, and test errors.
2. Think: identify the next smallest MECE task and expected evidence.
3. Output: produce one artifact or code change.
4. Verify: run the narrowest meaningful check.
5. Feedback: update todo/results and loop back when evidence disagrees.

Work loops must be:

- Closed: every loop has evaluation.
- Feedback-capable: results affect the next action.
- Reusable: artifacts are written into repo files.
- Measurable: commands or explicit review checks define success.

## Steering Loop

Harness accumulates quality over time:

1. Observe: identify recurring agent failure, ambiguity, or waiting.
2. Improve Harness: add an earlier guide, constraint, template, or check.
3. Verify Effect: prove the failure is prevented or detected faster.

## MECE Task Split

Tasks should be mutually exclusive and collectively exhaustive:

- Atomize tasks so each has one responsibility.
- Avoid overlapping ownership between tasks.
- Cover boundary cases and exceptional paths in the spec.
- Make task completion visible through tests, build output, or review evidence.

## SDIE Outputs In This Repo

- `tasks/spec.md`: Specification artifact.
- `tasks/plan.md`: Design artifact.
- `tasks/todo.md`: Implementation and Evaluation tracking.
- `.codex/skills/sdie-harness-loop/references/*`: persistent
  project knowledge.
- Tests and build output: verification evidence.
- Final assistant response: human-readable V&V summary.
