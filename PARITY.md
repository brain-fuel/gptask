# gptask — parity ledger

Go+ rewrite of [`go-task/task`](https://github.com/go-task/task), pinned to
**v3.52.0** (MIT). Module: `goforge.dev/gptask`. Wave 2, candidate 2.

## Status

**Phase 1 — parity foundation: DONE.** go-task's own source (~13k LOC across the
root engine, `taskfile`/`taskfile/ast`, `internal/*`, `args`, `errors`,
`experiments`, `taskrc`) is vendored under `goforge.dev/gptask` (import-path
surgery only, third-party libs reused: `mvdan.cc/sh`, `fsnotify`, `slim-sprig`,
`pflag`, yaml, …). It **builds**, the `task` binary reports **3.52.0**, and:

- **CLI differential** vs the pinned upstream binary: default task, `--list`,
  and var-templated tasks are **byte-identical**.
- **Full test suite passes** — 13 packages green (the gitignore/fsext walk tests
  require gptask to be its own git boundary, which it now is; the lone
  `experiments` `go vet` catch — `%q` on an int — is pre-existing in upstream
  too, verified).

**Phase 2 — Go+ authorship (in progress).** Convert the vendored `.go` sources to
Go+ (`.gp`) with enum idiomata for the closed sums (task/step lifecycle states,
output styles, sort/precondition/status outcomes, experiment states), pin
released goplus v0.138.0, keep `gen -check` clean and parity intact. Wire
formats (the Taskfile YAML schema, CLI flags) must round-trip byte-identically.

**Phase 3 — broaden the differential + release.** Subprocess-diff the deeper
surface (deps/parallelism ordering, up-to-date fingerprinting, includes, env,
preconditions, status, watch, dotenv, `--summary`, remote taskfiles), then
release `goforge.dev/gptask`.

## Provenance

Pinned upstream MIT (© 2016 Andrey Nering). Third-party deps unchanged.
Repo: `git@github.com:brain-fuel/gptask.git`.
