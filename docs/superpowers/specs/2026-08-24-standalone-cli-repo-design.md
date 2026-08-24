# Standalone Flexprice CLI repo — design

Date: 2026-08-24

## Context

The Flexprice CLI currently lives at `cli/` inside the `flexprice/flexprice`
monorepo, with its own `go.mod` (`github.com/flexprice/cli`). The monorepo
already has a "mirror" pattern wired up in `.github/workflows/`:

- `cli-release.yml` — on `cli/v*` tags, pushes the `cli/` subtree into
  `flexprice/cli` and runs `goreleaser` targeting that repo for GitHub
  releases + a Homebrew tap.
- `cli-validate.yml` — runs tests/spec-drift checks on PRs to the monorepo
  that touch `cli/`.
- `cli-spec-sync.yml` — auto-opens a PR in the monorepo when the OpenAPI spec
  changes.

Right now all CI/CD and actual development happen in the monorepo;
`flexprice/cli` is a passive push target, and its README tells contributors
to open PRs against the monorepo instead.

`github.com/flexprice/cli` already exists on GitHub (private), but is just a
placeholder — a one-line `# cli` README, no releases, no branches beyond
`main`.

`flexprice/homebrew-tap` does **not** exist yet, so the existing
`.goreleaser.yaml`'s `brews:` publish step would fail if run today.

## Goal

Turn `github.com/flexprice/cli` into the real, standalone home for CLI
development: source of truth, its own CI/CD, no more dependency on the
monorepo's push mechanism.

Explicitly out of scope for this pass (deferred, deliberately):

- Cleaning up the monorepo side (`cli/` directory, `cli-release.yml`,
  `cli-validate.yml`, `cli-spec-sync.yml` stay as-is for now).
- Automated spec sync from the monorepo (manual for now — someone re-copies
  `docs/swagger/swagger-3-0.json` by hand when it changes).
- Cutting an actual release tag, or creating `flexprice/homebrew-tap`.

## 1. Repository layout

Flatten `cli/*` to the repo root — the nesting existed only to coexist with
the backend inside the monorepo. `go.mod`, `main.go`, `internal/`, `tools/`,
`spec/`, `guides/`, `integration/`, `scripts/`, `ARCHITECTURE.md`,
`AGENTS.md`, `README.md`, `LICENSE`, `.gitignore`, `.goreleaser.yaml` all
move up one level. Module path stays `github.com/flexprice/cli` — already
correct for a standalone repo, no rename needed.

## 2. History migration

Use `git filter-repo --subdirectory-filter cli` on a disposable clone of the
monorepo to extract just the CLI's commit history (original authors, dates,
messages intact), rewritten as if `cli/` had always been the repo root. That
becomes this repo's history.

`github.com/flexprice/cli` already has one unrelated commit (the placeholder
README), so publishing the filtered history there means force-pushing over
unrelated history. This requires explicit confirmation immediately before
the push — it is not assumed here.

## 3. CI/CD pipelines (`.github/workflows/`)

- **`validate.yml`** — on PR + push to `main`: `go build ./...`,
  `go vet ./...`, `go test -race ./...`.
- **`lint.yml`** — `golangci-lint` on PRs. No lint config exists today for
  the CLI; add a new `.golangci.yml` at the repo root with a sane default
  set (govet, staticcheck, errcheck, unused, ineffassign, gofmt/goimports).
- **`release.yml`** — on `v*` tag push: run tests, then
  `goreleaser release --clean`. Builds darwin/linux/windows × amd64/arm64
  archives and a GitHub release using the default `GITHUB_TOKEN` (same-repo
  release now — no cross-repo secret needed). Tag scheme drops the `cli/`
  prefix used in the monorepo (`v1.2.3`, not `cli/v1.2.3`) since there's no
  backend release to disambiguate from anymore.
- **`codeql.yml`** — CodeQL analysis for Go: on push to `main`, on PRs, and
  weekly on a schedule.
- **`docs.yml`** — on push to `main` touching `internal/cmd/**` or
  `spec/**`: run `tools/gendocs`, auto-commit the regenerated `docs/`
  markdown back to `main` as a bot commit with `[skip ci]` in the message to
  avoid a workflow loop. This is the first time `gendocs` gets wired into
  any pipeline — today it only runs via `make cli-docs` by hand.
- **`dependabot.yml`** — weekly update PRs for the `gomod` and
  `github-actions` ecosystems.

## 4. Config/doc updates

- `.goreleaser.yaml`:
  - Drop the header comment about goreleaser running from inside a `cli/`
    subdirectory of a monorepo checkout — no longer true, it runs from the
    repo root now.
  - Comment out the `brews:` section, with a note that it's disabled until
    `flexprice/homebrew-tap` exists. Re-enabling it later is a one-line
    uncomment plus creating the tap repo.
  - Drop the explicit `release.github: {owner: flexprice, name: cli}`
    override — goreleaser defaults to releasing against the repo it runs in,
    which is now correct without an override.
- `README.md`: remove the "this repo is a release mirror, open PRs against
  the monorepo" language — this repo is now where contributions happen.
  Keep the rest of the user-facing content (install, quickstart, commands,
  scripting) as-is; it already describes the CLI correctly regardless of
  where its source lives.
- `CODEOWNERS`: `* @flexprice/cli-maintainers` at the repo root (was scoped
  to `/cli/` inside the monorepo).

## 5. Secrets / prerequisites

None required for `validate`, `lint`, `codeql`, or `docs` — all same-repo,
default `GITHUB_TOKEN` with `contents: write` (docs auto-commit) and
`security-events: write` (CodeQL) is enough. `release.yml` also only needs
the default `GITHUB_TOKEN` now that Homebrew publishing is disabled and the
release targets the repo it runs in.

## 6. Rollout order

1. Extract history via `git filter-repo` into a local working copy.
2. Restructure into this repo locally (flatten layout, add the new
   workflows/config, apply the doc updates above).
3. Verify locally: `go build ./...`, `go test -race ./...`,
   `golangci-lint run`.
4. Confirm with the user, then force-push to `github.com/flexprice/cli`
   `main`.
5. Verify CI goes green on GitHub.
6. Do **not** cut a release tag as part of this work — separate, deliberate
   step later, after revisiting the Homebrew tap decision.

## Testing

- `go build ./...` and `go test -race ./...` must pass locally before any
  push.
- `golangci-lint run` must pass (or findings triaged) before any push.
- After pushing, all GitHub Actions workflows (`validate`, `lint`, `codeql`)
  must run green on the resulting `main` branch. `release.yml` is not
  exercised (no tag is pushed in this pass), but its YAML should be
  reviewed for correctness.
