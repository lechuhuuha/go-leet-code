# AGENTS.md

## Purpose

This repository is a practice and learning workspace. The default collaboration style is guidance-first.

## Collaboration Rules

- Prefer guiding, explaining, reviewing, and suggesting over directly modifying source code.
- Assume the user wants to implement solutions by themselves unless they explicitly ask for code changes.
- When helping on a problem, explain the approach, tradeoffs, algorithm choice, and edge cases to consider.
- Prefer giving small, concrete implementation steps instead of pasting full solutions.
- When reviewing code, focus on correctness, clarity, complexity, naming, and missing test coverage.
- If code examples are needed, keep them minimal and use them to illustrate an idea rather than to complete the task for the user.
- Do not proactively refactor unrelated code.
- Do not change existing files unless the user clearly asks for an edit.

## File Naming

- Use `snake_case` for file names.
- Go source files should follow `snake_case.go`.
- Go test files should follow `snake_case_test.go`.
- Avoid spaces, mixed casing, and names like `copy`, `final`, or `new` unless they describe a real variant with clear intent.
- Keep file names aligned with the problem or feature name so the implementation file and test file are easy to pair.

## Unit Test Conventions

- Prefer table-driven tests for unit tests.
- Prefer the Go standard library for unit tests and supporting code instead of adding new dependencies.
- Each test table should include a `name` field so every case runs as a subtest with `t.Run(...)`.
- Cover at least these categories when relevant:
  - happy path
  - edge cases
  - error or invalid-input cases
  - boundary conditions
  - duplicates, empty input, nil input, single-element input, and large/special values where applicable
- Keep expected outputs explicit in each case.
- Use descriptive test names that explain the behavior being validated.
- Add cases for regressions when a bug is found.
- If order matters, assert exact order.
- If order does not matter, make that explicit in the assertion strategy.
- Include Unicode or rune-focused test cases when string behavior depends on characters rather than bytes.

## Suggested Test Structure

- Arrange tests as a slice of structs.
- Include only fields needed by the function under test plus the expected result.
- Iterate over the table and run each case as a subtest.
- Keep assertions simple and readable.
- Prefer built-in packages such as `testing`, `reflect`, `cmp` alternatives implemented locally when simple, and other standard library helpers before introducing third-party assertions or test frameworks.
- Prefer deterministic tests with no shared mutable state between cases.

## Dependencies

- Prefer the Go standard library before introducing any third-party library.
- For unit tests, use built-in tools like `testing`, `reflect`, slices/maps comparison logic, and helper functions in the test file when enough.
- Add a new dependency only when the standard library would make the solution significantly worse in clarity or correctness.
- If a new dependency is proposed, explain why the standard library is not sufficient first.

## Response Style

- Start with the reasoning or checklist the user should follow.
- Call out edge cases before implementation details.
- When useful, propose a test matrix the user can implement.
- If the user asks for code, confirm whether they want a full implementation or only a hint-level example.
