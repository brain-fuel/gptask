# Changelog

## v1.1.1

- Exclude generated remote-include cache files from module archives so clean
  `go install` works on every platform.

## v1.1.0

- Add reusable Quicken doctor, build, and run task definitions for Web, TUI,
  desktop, iOS, and Android targets.
- Serialize executor output streams through one per-executor lock so parallel
  dependencies remain race-free even when stdout and stderr share a writer.
- Make lazy template caches safe for concurrent output rendering.
- Correct invalid integer formatting in experiment diagnostics.

## v1.0.3

- Update `goforge.dev/goplus/std` to v0.207.0.

## v1.0.0

- Initial Go+ rewrite at parity with `go-task/task` v3.52.0.
