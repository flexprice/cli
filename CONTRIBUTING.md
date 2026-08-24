# Contributing to the Flexprice CLI

## Development setup

    git clone https://github.com/flexprice/cli.git
    cd cli
    make build

`make build` puts a `flexprice` binary at `bin/flexprice`. Run it directly,
or use `make run ARGS="customers list"` to build-and-run in one step.

## Testing and linting

    make test    # go test -race ./...
    make lint    # golangci-lint run ./...
    make check   # build + vet + test + lint — what CI runs on every PR

`make help` lists every available target, including the smoke (`make smoke`)
and end-to-end (`make e2e`) suites under `scripts/`.

## Adding a command

This CLI dispatches commands at runtime from an embedded OpenAPI spec rather
than generating Go source per command. [ARCHITECTURE.md](ARCHITECTURE.md)
walks through the request lifecycle end to end. The two most common
maintenance tasks are walked through step by step in `guides/`:

- [Adding a generated command](guides/adding-a-command.md) — the common
  case, for any operation already in the OpenAPI spec.
- [Adding a hand-written command](guides/adding-a-hand-written-command.md) —
  for commands like `init` or `whoami` that don't map to a single API
  operation.

## Regenerating the command reference

    make docs

writes the Markdown command reference to `docs/`. CI regenerates and commits
this automatically on merges that touch command or spec code, so you don't
need to run it by hand before opening a PR — it's here for local iteration.

## Code of conduct

Participation in this project is governed by our
[Code of Conduct](CODE_OF_CONDUCT.md).
