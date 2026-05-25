# AGENTS.md

## Creating New LeetCode Problem Files (Rust)

When creating a new Rust file for a LeetCode problem:

1. File naming: `p_{number}_{snake_case_title}.rs` ( e.g.,
   `p_2733_neither_minimum_nor_maximum.rs`). The file should contain `struct
   Solution` and the function should belong to an `impl Solution` block.
2. Register the module `src/lib.rs`.
3. Function body: Use `todo!()` — do NOT include the solution.
4. Problem description: Add the problem description as Rust doc comments (
   `///`) with Markdown formatting above the function signature.
5. Tests: Include a `#[cfg(test)]` module with sample test cases from the
   problem.
6. Tests should be parameterized via `#[test_case(...)`
