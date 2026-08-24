# Standalone Flexprice CLI repo Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Turn the empty `/Users/omkar/Developer/source-code/cli` repo into the standalone, CI/CD-equipped home for the Flexprice CLI, with its full commit history imported from `flexprice/flexprice`'s `cli/` subtree.

**Architecture:** Extract `cli/`'s history from the monorepo with `git filter-repo`, merge it into this repo (which already has one commit — the design spec) with `--allow-unrelated-histories`, flatten the layout to the repo root, then add GitHub Actions workflows for validation, linting, release, CodeQL, and docs generation. Local verification (build/test/lint) gates every step; the final GitHub push is a separate, explicitly-confirmed step.

**Tech Stack:** Go 1.25, cobra, golangci-lint v1.64.8, goreleaser, GitHub Actions.

---

### Task 1: Extract CLI history from the monorepo

**Files:** none in the target repo — this operates entirely in a scratch clone at `/Users/omkar/Developer/source-code/cli-history-extract`.

- [ ] **Step 1: Record the expected commit count**

Run: `cd /Users/omkar/Developer/source-code/flexprice/flexprice && git log --oneline -- cli/ | wc -l`
Expected: `74`

- [ ] **Step 2: Clone the monorepo into a scratch directory**

Run:
```bash
rm -rf /Users/omkar/Developer/source-code/cli-history-extract
git clone /Users/omkar/Developer/source-code/flexprice/flexprice /Users/omkar/Developer/source-code/cli-history-extract
```
Expected: clone completes, prints `done.` and a remote-tracking summary.

- [ ] **Step 3: Filter the clone down to the `cli/` subtree**

Run:
```bash
cd /Users/omkar/Developer/source-code/cli-history-extract
git filter-repo --subdirectory-filter cli
```
Expected: output ending in `New size ...` with no errors. `git-filter-repo` also strips the `origin` remote as a safety measure — that's expected.

- [ ] **Step 4: Verify the filtered history**

Run:
```bash
cd /Users/omkar/Developer/source-code/cli-history-extract
git log --oneline | wc -l
ls
```
Expected: the count matches Step 1 (`74`), and `ls` shows `go.mod main.go internal spec tools guides integration scripts ARCHITECTURE.md AGENTS.md README.md LICENSE .goreleaser.yaml .gitignore` at the top level — **not** any monorepo backend directories (`internal/domain`, `ent/`, etc. must be absent).

No commit in this task — it only prepares the scratch repo that Task 2 pulls from.

---

### Task 2: Merge the extracted history into the target repo and flatten references

**Files:**
- Modify: `AGENTS.md:2-4` and the 6 files matching `internal/*/AGENTS.md` (strip the `cli/` prefix from `owns:` globs)

- [ ] **Step 1: Add the scratch repo as a remote and fetch it**

Run:
```bash
cd /Users/omkar/Developer/source-code/cli
git remote add cli-history /Users/omkar/Developer/source-code/cli-history-extract
git fetch cli-history
```
Expected: fetch succeeds, prints a new branch ref like `cli-history/main`.

- [ ] **Step 2: Merge with unrelated histories allowed**

Run:
```bash
cd /Users/omkar/Developer/source-code/cli
git merge cli-history/main --allow-unrelated-histories -m "merge: import Flexprice CLI history from flexprice/flexprice cli/"
```
Expected: merge succeeds with no conflicts (the only pre-existing file, `docs/superpowers/specs/2026-08-24-standalone-cli-repo-design.md`, doesn't exist in the incoming history, so there's nothing to conflict with).

- [ ] **Step 3: Verify the merge**

Run:
```bash
cd /Users/omkar/Developer/source-code/cli
ls
git log --oneline | wc -l
```
Expected: `ls` shows both the CLI files (`go.mod`, `main.go`, `internal/`, `spec/`, `tools/`, `guides/`, `integration/`, `scripts/`, `ARCHITECTURE.md`, `AGENTS.md`, `README.md`, `LICENSE`, `.goreleaser.yaml`, `.gitignore`) and `docs/` (containing `superpowers/specs/...`). Commit count is `76` (74 imported + 1 spec commit + 1 merge commit).

- [ ] **Step 4: Remove the scratch remote**

Run:
```bash
cd /Users/omkar/Developer/source-code/cli
git remote remove cli-history
```
Expected: no output, `git remote -v` no longer lists `cli-history`.

- [ ] **Step 5: Fix the `owns:` globs left stale by flattening**

The 7 `AGENTS.md` files still say `owns: ["cli/**"]` (or a `cli/internal/...` variant) from when they lived inside the monorepo's `cli/` subdirectory. Now that this repo's root **is** that directory, the `cli/` prefix is wrong. Fix all seven with:

```bash
cd /Users/omkar/Developer/source-code/cli
grep -rl '"cli/' --include=AGENTS.md . | xargs sed -i '' 's|"cli/|"|g'
```

- [ ] **Step 6: Verify no stale references remain**

Run: `cd /Users/omkar/Developer/source-code/cli && grep -rn '"cli/' --include=AGENTS.md .`
Expected: no output (no matches).

- [ ] **Step 7: Commit**

```bash
cd /Users/omkar/Developer/source-code/cli
git add AGENTS.md internal/cmd/AGENTS.md internal/config/AGENTS.md internal/spec/AGENTS.md internal/output/AGENTS.md internal/keyring/AGENTS.md internal/client/AGENTS.md
git commit -m "chore: drop stale cli/ prefix from AGENTS.md ownership globs"
```

---

### Task 3: Verify the imported code builds and tests cleanly

**Files:** none — verification only.

- [ ] **Step 1: Build**

Run: `cd /Users/omkar/Developer/source-code/cli && go build ./...`
Expected: exits 0, no output.

- [ ] **Step 2: Test**

Run: `cd /Users/omkar/Developer/source-code/cli && go test -race ./...`
Expected: exits 0, all packages report `ok`.

If either fails, stop and diagnose before continuing — the merge should be a pure history/reference change with no functional impact, so a failure here means Task 2 went wrong.

---

### Task 4: Add golangci-lint config

**Files:**
- Create: `.golangci.yml`

- [ ] **Step 1: Write the config**

```yaml
run:
  timeout: 5m

linters:
  disable-all: true
  enable:
    - govet
    - staticcheck
    - errcheck
    - unused
    - ineffassign
    - gofmt
    - goimports
```

- [ ] **Step 2: Run it locally**

Run: `cd /Users/omkar/Developer/source-code/cli && golangci-lint run ./...`
Expected: exits 0. If it reports findings, fix them in the flagged files before continuing (do not weaken the config to silence them — these are pre-existing files that should already be clean Go).

- [ ] **Step 3: Commit**

```bash
cd /Users/omkar/Developer/source-code/cli
git add .golangci.yml
git commit -m "ci: add golangci-lint config"
```

---

### Task 5: Add PR/push build-and-test workflow

**Files:**
- Create: `.github/workflows/validate.yml`

- [ ] **Step 1: Write the workflow**

```yaml
name: Validate

on:
  pull_request:
  push:
    branches: [main]

jobs:
  build-test:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4

      - uses: actions/setup-go@v5
        with:
          go-version-file: go.mod

      - name: Build
        run: go build ./...

      - name: Vet
        run: go vet ./...

      - name: Test
        run: go test -race ./...
```

- [ ] **Step 2: Validate the YAML parses**

Run: `cd /Users/omkar/Developer/source-code/cli && python3 -c "import yaml,sys; yaml.safe_load(open('.github/workflows/validate.yml'))" && echo OK`
Expected: `OK`

- [ ] **Step 3: Commit**

```bash
cd /Users/omkar/Developer/source-code/cli
git add .github/workflows/validate.yml
git commit -m "ci: add build/vet/test workflow"
```

---

### Task 6: Add lint workflow

**Files:**
- Create: `.github/workflows/lint.yml`

- [ ] **Step 1: Write the workflow**

```yaml
name: Lint

on:
  pull_request:
  push:
    branches: [main]

jobs:
  golangci-lint:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4

      - uses: actions/setup-go@v5
        with:
          go-version-file: go.mod

      - uses: golangci/golangci-lint-action@v6
        with:
          version: v1.64.8
```

- [ ] **Step 2: Validate the YAML parses**

Run: `cd /Users/omkar/Developer/source-code/cli && python3 -c "import yaml,sys; yaml.safe_load(open('.github/workflows/lint.yml'))" && echo OK`
Expected: `OK`

- [ ] **Step 3: Commit**

```bash
cd /Users/omkar/Developer/source-code/cli
git add .github/workflows/lint.yml
git commit -m "ci: add golangci-lint workflow"
```

---

### Task 7: Update goreleaser config and add the release workflow

**Files:**
- Modify: `.goreleaser.yaml` (full contents below)
- Create: `.github/workflows/release.yml`

- [ ] **Step 1: Rewrite `.goreleaser.yaml`**

Replace the entire file with:

```yaml
version: 2
project_name: flexprice

before:
  hooks:
    - go mod tidy

builds:
  - id: flexprice
    main: .
    binary: flexprice
    env:
      - CGO_ENABLED=0
    goos: [darwin, linux, windows]
    goarch: [amd64, arm64]
    ldflags:
      - -s -w -X main.version={{.Version}}

archives:
  - formats: [tar.gz]
    format_overrides:
      - goos: windows
        formats: [zip]

checksum:
  name_template: checksums.txt

# Homebrew publishing is disabled until github.com/flexprice/homebrew-tap
# exists. Once it does, uncomment this block — no other changes needed,
# goreleaser already targets the repo it runs in for the GitHub release.
# brews:
#   - repository:
#       owner: flexprice
#       name: homebrew-tap
#       token: "{{ .Env.SDK_DEPLOY_GIT_TOKEN }}"
#     homepage: https://flexprice.io
#     description: Usage-based billing from your terminal
#     license: Apache-2.0
```

- [ ] **Step 2: Verify the YAML parses**

Run: `cd /Users/omkar/Developer/source-code/cli && python3 -c "import yaml; yaml.safe_load(open('.goreleaser.yaml'))" && echo OK`
Expected: `OK`

- [ ] **Step 3: Write the release workflow**

```yaml
name: Release

on:
  push:
    tags: ['v*']

permissions:
  contents: write

jobs:
  release:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
        with:
          fetch-depth: 0

      - uses: actions/setup-go@v5
        with:
          go-version-file: go.mod

      - name: Test
        run: go test -race ./...

      - name: Release
        uses: goreleaser/goreleaser-action@v6
        with:
          args: release --clean
        env:
          GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}
```

- [ ] **Step 4: Validate the YAML parses**

Run: `cd /Users/omkar/Developer/source-code/cli && python3 -c "import yaml; yaml.safe_load(open('.github/workflows/release.yml'))" && echo OK`
Expected: `OK`

- [ ] **Step 5: Commit**

```bash
cd /Users/omkar/Developer/source-code/cli
git add .goreleaser.yaml .github/workflows/release.yml
git commit -m "ci: add release workflow, disable homebrew publish until tap repo exists"
```

---

### Task 8: Add CodeQL workflow

**Files:**
- Create: `.github/workflows/codeql.yml`

- [ ] **Step 1: Write the workflow**

```yaml
name: CodeQL

on:
  push:
    branches: [main]
  pull_request:
    branches: [main]
  schedule:
    - cron: '0 6 * * 1'

jobs:
  analyze:
    runs-on: ubuntu-latest
    permissions:
      security-events: write
      contents: read
    steps:
      - uses: actions/checkout@v4

      - uses: github/codeql-action/init@v3
        with:
          languages: go

      - uses: github/codeql-action/autobuild@v3

      - uses: github/codeql-action/analyze@v3
```

- [ ] **Step 2: Validate the YAML parses**

Run: `cd /Users/omkar/Developer/source-code/cli && python3 -c "import yaml; yaml.safe_load(open('.github/workflows/codeql.yml'))" && echo OK`
Expected: `OK`

- [ ] **Step 3: Commit**

```bash
cd /Users/omkar/Developer/source-code/cli
git add .github/workflows/codeql.yml
git commit -m "ci: add CodeQL analysis workflow"
```

---

### Task 9: Wire up the docs generation workflow

**Files:**
- Create: `.github/workflows/docs.yml`
- Modify: `tools/gendocs/main.go:2` (drop the now-stale TODO comment)

- [ ] **Step 1: Drop the stale TODO comment**

`tools/gendocs/main.go` currently starts with:
```go
// Command gendocs writes the CLI command reference as Markdown.
// TODO: not wired into cli-release.yml; runs only via `make cli-docs`.
package main
```

This workflow is what wires it up, so the TODO is no longer true. Replace those two lines with:
```go
// Command gendocs writes the CLI command reference as Markdown.
package main
```

- [ ] **Step 2: Confirm `tools/gendocs` still builds**

Run: `cd /Users/omkar/Developer/source-code/cli && go run ./tools/gendocs && ls docs/*.md | head -5`
Expected: exits 0, and `docs/*.md` lists generated command reference files (e.g. `docs/flexprice.md`, `docs/flexprice_customers.md`, ...).

- [ ] **Step 3: Leave the locally-generated docs uncommitted**

Run: `cd /Users/omkar/Developer/source-code/cli && git status --porcelain docs/`
Expected: untracked `docs/*.md` files from Step 2. Leave them untracked for now — Step 6 commits the workflow file and the `main.go` comment fix only, not these generated docs (the workflow itself creates and commits them on its first run against `main`).

- [ ] **Step 4: Write the workflow**

```yaml
name: Docs

on:
  push:
    branches: [main]
    paths:
      - 'internal/cmd/**'
      - 'spec/**'

permissions:
  contents: write

jobs:
  gendocs:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4

      - uses: actions/setup-go@v5
        with:
          go-version-file: go.mod

      - name: Generate command reference
        run: go run ./tools/gendocs

      - uses: stefanzweifel/git-auto-commit-action@v5
        with:
          commit_message: 'docs: regenerate command reference [skip ci]'
          file_pattern: 'docs/*.md'
```

- [ ] **Step 5: Validate the YAML parses**

Run: `cd /Users/omkar/Developer/source-code/cli && python3 -c "import yaml; yaml.safe_load(open('.github/workflows/docs.yml'))" && echo OK`
Expected: `OK`

- [ ] **Step 6: Commit**

```bash
cd /Users/omkar/Developer/source-code/cli
git add .github/workflows/docs.yml tools/gendocs/main.go
git commit -m "ci: wire up gendocs into a docs-publish workflow"
```

---

### Task 10: Add Dependabot config

**Files:**
- Create: `.github/dependabot.yml`

- [ ] **Step 1: Write the config**

```yaml
version: 2
updates:
  - package-ecosystem: "gomod"
    directory: "/"
    schedule:
      interval: "weekly"

  - package-ecosystem: "github-actions"
    directory: "/"
    schedule:
      interval: "weekly"
```

- [ ] **Step 2: Validate the YAML parses**

Run: `cd /Users/omkar/Developer/source-code/cli && python3 -c "import yaml; yaml.safe_load(open('.github/dependabot.yml'))" && echo OK`
Expected: `OK`

- [ ] **Step 3: Commit**

```bash
cd /Users/omkar/Developer/source-code/cli
git add .github/dependabot.yml
git commit -m "ci: add dependabot config for gomod and github-actions"
```

---

### Task 11: Update README for standalone contribution model

**Files:**
- Modify: `README.md`

- [ ] **Step 1: Replace the mirror-language Contributing section**

Find:
```markdown
## Contributing

**Source of truth is `flexprice/flexprice` at `cli/`.** This repository is a
release mirror — please open pull requests against the monorepo. Issues here
are welcome.
```

Replace with:
```markdown
## Contributing

Pull requests and issues are welcome directly against this repository.
```

- [ ] **Step 2: Fix the now-stale "build from inside `cli/`" instruction**

Find:
```markdown
Build and test locally with the standard Go toolchain from inside `cli/`:

    go build ./...
    go test -race ./...
```

Replace with:
```markdown
Build and test locally with the standard Go toolchain:

    go build ./...
    go test -race ./...
```

- [ ] **Step 3: Verify no stale references remain**

Run: `cd /Users/omkar/Developer/source-code/cli && grep -n "release mirror\|inside \`cli/\`\|monorepo" README.md`
Expected: no output (no matches).

- [ ] **Step 4: Commit**

```bash
cd /Users/omkar/Developer/source-code/cli
git add README.md
git commit -m "docs: update README for standalone contribution model"
```

---

### Task 12: Add repo-root CODEOWNERS

**Files:**
- Create: `.github/CODEOWNERS`

- [ ] **Step 1: Write the file**

```
* @flexprice/cli-maintainers
```

- [ ] **Step 2: Commit**

```bash
cd /Users/omkar/Developer/source-code/cli
git add .github/CODEOWNERS
git commit -m "chore: add CODEOWNERS"
```

---

### Task 13: Final local verification

**Files:** none — verification only.

- [ ] **Step 1: Full build**

Run: `cd /Users/omkar/Developer/source-code/cli && go build ./...`
Expected: exits 0.

- [ ] **Step 2: Full test suite**

Run: `cd /Users/omkar/Developer/source-code/cli && go test -race ./...`
Expected: exits 0, all `ok`.

- [ ] **Step 3: Lint**

Run: `cd /Users/omkar/Developer/source-code/cli && golangci-lint run ./...`
Expected: exits 0.

- [ ] **Step 4: Review the commit log**

Run: `cd /Users/omkar/Developer/source-code/cli && git log --oneline -12`
Expected: the merge commit from Task 2, followed by one commit per task (AGENTS.md fix, golangci config, validate workflow, lint workflow, release workflow, CodeQL workflow, docs workflow, dependabot, README, CODEOWNERS).

- [ ] **Step 5: Clean up the scratch clone**

Run: `rm -rf /Users/omkar/Developer/source-code/cli-history-extract`

---

### Task 14: Push to `github.com/flexprice/cli` (requires explicit confirmation)

**Files:** none — this is a remote-repository push.

`github.com/flexprice/cli` currently has one unrelated commit (a placeholder `# cli` README). This repo's history is unrelated to it, so publishing requires a **force push**, which permanently discards that placeholder commit from `main`'s history on GitHub. Per the design spec, this step requires explicit user confirmation immediately before running — do not run Step 2 without it, even if every prior task succeeded.

- [ ] **Step 1: Add the GitHub remote**

Run:
```bash
cd /Users/omkar/Developer/source-code/cli
git remote add origin https://github.com/flexprice/cli.git
git remote -v
```
Expected: `origin` listed with fetch/push URLs pointing at `github.com/flexprice/cli`.

- [ ] **Step 2: Confirm with the user, then force-push**

Ask: "Ready to force-push this repo's history to github.com/flexprice/cli, replacing its current placeholder commit. Confirm?" Only after an explicit yes, run:

```bash
cd /Users/omkar/Developer/source-code/cli
git push --force-with-lease origin main
```
Expected: push succeeds, GitHub's `main` now matches local `main`.

- [ ] **Step 3: Verify on GitHub**

Run: `gh repo view flexprice/cli --json defaultBranchRef,pushedAt`
Expected: `pushedAt` reflects the just-completed push.

- [ ] **Step 4: Verify CI runs**

Run: `gh run list -R flexprice/cli --limit 5`
Expected: `Validate`, `Lint`, and `CodeQL` workflows listed as running or completed (triggered by the push to `main`). Do not proceed to check "success" status immediately — GitHub Actions takes a minute or two to complete; check back after a short wait with `gh run list -R flexprice/cli --limit 5` again, and inspect any failure with `gh run view <run-id> -R flexprice/cli --log-failed` before considering this task done.

No release tag is pushed as part of this plan — `release.yml` is verified by YAML validity only (Task 7, Step 4), not by an actual run.
