# Learning ledger

One row per 2-hour block. Each block has a task in my Todoist (`Builder · GO-N: <outcome>`). It counts as done only when it has a link: a commit, a tag or a demo. A block closes as **PASS** (shipped, link below), **LOSS** (no link by Sunday) or **OVERRIDE** (the week was declared off). A missed block stays next. Nothing doubles up.

Blocks run every other Wednesday.

## Plan

| Block | Planned | Ships | Go concepts |
|---|---|---|---|
| GO-1 | Wed Oct 28 | `shelfpace pace --pages 320 --at 45 --by 2026-11-12`: pages/day and a dated page table ending the day before the meeting | modules, `package main`, `flag`, `time` and dates, integer math and remainders, returning `error` values |
| GO-2 | Wed Nov 11 | Plan logic moved into a pure `internal/plan` package, with golden-file tests and CI (vet, test, `-race`) | packages and `internal/`, table-driven tests, `t.Helper`, golden files, GitHub Actions |
| GO-3 | Wed Nov 25 | `shelfpace import goodreads\|storygraph\|md <file>` and `shelfpace ls` grouped by theme | a small consumer-side interface (`Importer`), `io.Reader`, `encoding/csv`, `bufio.Scanner`, wrapping errors with `%w` |
| GO-4 | Wed Dec 9 | `shelfpace plan --days 14`: multi-book plan with deadline alerts (EDF, stride across themes, WIP limit, author rule) | `slices`/`maps`/`cmp`, `container/heap`, `Stringer`, deterministic tie-breaks, property-style tests |
| GO-5 | Wed Dec 23 | **MVP:** `log`, `today`, a subscribed `.ics` feed on my phone, one morning push | `encoding/json`, `os.UserConfigDir`, writing RFC 5545 by hand, `net/http` client |
| GO-6 | | SQLite persistence and learned pace (EWMA) | `database/sql`, embedded migrations with `go:embed`, transactions, `sql.ErrNoRows` |
| GO-7 | | `shelfpace serve`: today/14-day/alerts pages, log form, feed endpoint with ETag | `ServeMux` method + wildcard patterns, `html/template`, middleware, `httptest` |
| GO-8 | | `shelfpace enrich`: fills missing page counts and subjects from Open Library | `errgroup.SetLimit`, `x/time/rate`, `context` cancellation, caching, backoff |
| GO-9 | | Always-on service: 05:00 nightly re-plan, clean shutdown | goroutines, `time.Timer`, `select`, `signal.NotifyContext`, `http.Server.Shutdown` |
| GO-10 | | Vault I/O: read my reading notes through `fs.FS`, write a plan note back between marker comments | `io/fs`, `testing/fstest`, `os.Root`, atomic writes (temp file + rename), idempotence fuzz test |
| GO-11 | | Hardening: structured logs, profiling endpoint, fuzz corpus, benchmark table | `log/slog`, `net/http/pprof`, native fuzzing, `testing.B` with `b.Loop` |
| GO-12 | | Generics refactor: one `PriorityQueue[T]` serving both EDF and stride | type parameters, constraints, iterators (`iter.Seq`) |
| GO-13 | | What-if and blackout ranges (e.g. exam week = 0 minutes) | interface seams for clock and calendar, more property tests |
| GO-14 | | v1.0.0: release builds, architecture diagram, "learning Go by building shelfpace" write-up | release workflow, `-trimpath`, `runtime/debug.ReadBuildInfo` |

## Log

| Date | Block | Status | Evidence | What I learned |
|---|---|---|---|---|
| 2026-10-24 | GO-0 | PASS | [repo created](https://github.com/munadm/shelfpace) | — |
| 2026-11-21 | GO-1 | PASS | [PR #1](https://github.com/munadm/shelfpace/pull/1) | learned basic variable initialization, how errors work, how there's only for loops and basic if statements, how falsy , truthy things are not accepted, how panic is not an error and should never happen, how to intialize using var, const, := if inside a function, ExitOnError returns before the parse command can run |
