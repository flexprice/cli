# Flexprice CLI

[![License](https://img.shields.io/badge/license-Apache%202.0-blue.svg)](LICENSE)
[![Release](https://img.shields.io/github/v/release/flexprice/cli)](https://github.com/flexprice/cli/releases)
[![CI](https://github.com/flexprice/cli/actions/workflows/ci.yml/badge.svg)](https://github.com/flexprice/cli/actions/workflows/ci.yml)

Usage-based billing from your terminal. Send events, inspect how they were
metered, and drive the Flexprice API without leaving the command line.

**With the CLI, you can:**

- Send usage events and inspect how they were metered
- Manage customers, subscriptions, invoices, and every other billing resource without leaving the terminal
- Script against a stable JSON output contract and documented exit codes
- Reach any API endpoint through the `get`/`post`/`delete` escape hatch, even ones without a named command

    go install github.com/flexprice/cli@latest
    flexprice init

## Install

**Download a release (macOS, Linux, Windows)**

Grab the archive for your platform from the [latest release](https://github.com/flexprice/cli/releases/latest),
extract it, and put the `flexprice` binary on your `PATH`.

**Go**

    go install github.com/flexprice/cli@latest

**Upgrading:** download the new release and replace the binary, or run
`go install github.com/flexprice/cli@latest` again. Either way, your config
and stored keys are untouched.

## Quickstart

`flexprice init` walks you through picking a data region and pasting an API
key, verifies it against the API, and stores it in your OS keychain (or an
encrypted file when no keychain is available, e.g. in a container or CI).

    $ flexprice init
    ███████╗██╗     ███████╗██╗  ██╗██████╗ ██████╗ ██╗ ██████╗███████╗
    ██╔════╝██║     ██╔════╝╚██╗██╔╝██╔══██╗██╔══██╗██║██╔════╝██╔════╝
    █████╗  ██║     █████╗   ╚███╔╝ ██████╔╝██████╔╝██║██║     █████╗
    ██╔══╝  ██║     ██╔══╝   ██╔██╗ ██╔═══╝ ██╔══██╗██║██║     ██╔══╝
    ██║     ███████╗███████╗██╔╝ ██╗██║     ██║  ██║██║╚██████╗███████╗
    ╚═╝     ╚══════╝╚══════╝╚═╝  ╚═╝╚═╝     ╚═╝  ╚═╝╚═╝ ╚═════╝╚══════╝
    Usage-based billing from your terminal

    Welcome to Flexprice — let's get you set up.
    Your API key is scoped to one environment — you can add more later with `flexprice login`.

    ? Data region
      > us      https://us.api.flexprice.io/v1
        in      https://api.cloud.flexprice.io/v1

    API key: ••••••••••••••••••••
    ⠹ Verifying your key…
    ✓ Verified — stored as profile "default" in encrypted file (~/.flexprice/keys)
    Note: the API does not report which environment a key belongs to, so label your
    profiles yourself (--profile-name, --label) and check with: flexprice whoami

    Here's what to try first:
      flexprice whoami            confirm what you are pointed at
      flexprice resources         see everything you can act on
      flexprice customers list    try a read
      flexprice env list          see your other environments

The region menu is arrow-key driven. The wordmark drops to a compact form on
narrow terminals, and everything above is written to stderr, so it never
pollutes piped output. Pass `--region` and `--api-key` to skip both prompts
entirely in CI.

Confirm what you're pointed at:

    $ flexprice whoami
    Profile:      default
    Label:
    Region:       us
    Base URL:     https://us.api.flexprice.io/v1
    Key backend:  encrypted file (~/.flexprice/keys)
    Key:          sk_test_…4b

See everything you can act on:

    $ flexprice resources
    addons                       create, delete, list, lookup, retrieve, update
    alert-settings               create, delete, list, retrieve, update
    alerts                       list
    checkout                     create, delete, retrieve
    costs                        active, analytics, analytics-v2, create, delete, list, retrieve, update
    coupon-associations          list, retrieve
    coupons                      create, delete, list, lookup, retrieve, update
    credit-grants                create, delete, for-addon, for-plan, retrieve, update
    credit-notes                 create, finalize, retrieve, void
    customers                    by-external-id, create, delete, entitlements, entitlements-by-external-id, list, retrieve, subscriptions, upcoming-grants, update, usage
    ... and 24 more

## Commands at a glance

197 commands across 34 resources are resolved at startup from the embedded
OpenAPI spec. Nothing below is committed generated code (see
[ARCHITECTURE.md](ARCHITECTURE.md)). A handful of commands don't map to a
single API operation and are hand-written instead:

| Command | What it does |
|---|---|
| `flexprice init` | Guided first-run setup |
| `flexprice login` / `flexprice logout` | Add or remove a stored profile |
| `flexprice whoami` | Show the active profile, region, and key backend |
| `flexprice env list` | List environments in your tenant |
| `flexprice config list` / `flexprice config use <profile>` | Manage stored profiles |
| `flexprice resources` | List every resource and its actions |
| `flexprice <resource> <action>` | The generated surface, e.g. `customers list`, `invoices finalize` |
| `flexprice get` / `post` / `delete <path>` | Raw HTTP escape hatch for anything not covered by a resource command |
| `flexprice open dashboard` / `flexprice open webhooks` | Open the web dashboard or webhook portal |
| `flexprice version` | Print the CLI version and embedded spec build |

## What you can do

Every resource in the API is a top-level command, grouped by action:

    flexprice customers list
    flexprice customers retrieve cust_01K...
    flexprice customers create --external_id=acme --email=billing@acme.com
    flexprice invoices finalize inv_01K...
    flexprice subscriptions cancel sub_01K...

For a request body too deep to express as flags (creating a subscription, for
example), open it in your editor with the required fields pre-filled:

    flexprice subscriptions create --edit

or supply it directly:

    flexprice subscriptions create --data @subscription.json

`subscriptions create --help` shows exactly which fields flags can reach and
which cannot:

    Fields you can set with flags:
      --billing_period  (required)  [string]
      --currency  (required)  [string]
      --plan_id  (required)  [string]
      ... (23 more optional fields)

    Nested fields — these cannot be set with flags:
      addons  [array]
      line_items  [array]
      phases  [array]
      ... (10 more)

    Use --edit to fill in a pre-built request body, or --data @file.json.

Destructive actions (`delete`, `void`, `cancel`, `terminate`, `archive`,
`finalize`) ask for confirmation with an arrow-key prompt; `--force` skips it.
Off a terminal, the command **refuses** rather than proceeding (see
[Running non-interactively](#running-non-interactively)).

Reads show a spinner while the request is in flight, and table output carries a
footer naming which profile and region actually served it. That's the fastest
way to catch "am I pointed where I think I am":

    $ flexprice customers list
    ID                NAME          STATUS    CREATED_AT
    cust_01J8XABCDEF  Ada Lovelace  active    2026-01-02
    cust_02           Grace Hopper  archived  2026-03-14

    profile: default · region: us · sandbox · v1.0.0

A write confirms what it did, so you never have to infer success from a table:

    $ flexprice customers create --external_id=acme --email=billing@acme.com
    ✓ Created customer cust_01J8XGHIJKL

An empty list points at the next step rather than saying nothing useful:

    $ flexprice customers list
    No customers yet.
      Create one with: flexprice customers create

The footer, receipt and empty-state lines all go to stderr, so none of them
appear in piped or redirected output.

Anything not covered by a named command is reachable through the raw escape
hatch:

    flexprice get /customers/cust_01K...
    flexprice post /events --data @event.json

## Authentication

An API key belongs to exactly one environment. There is no `--environment`
flag, because the key itself already determines it. Switching environments
means switching profiles:

    flexprice login --label "production"    # stores a second profile
    flexprice config list                    # see every stored profile
    flexprice -p production customers list   # use one for a single command
    flexprice logout -p production           # remove a profile and its key

`flexprice env list` shows every environment in your tenant, but the CLI
cannot tell you which one your active key belongs to: the API itself does not
expose that.

## Output & scripting

    flexprice customers list --output json > customers.json

Data always goes to stdout; progress messages, warnings, and footers always go
to stderr, so redirecting stdout never mixes the two. Exit codes are stable
and safe to script against:

| Code | Meaning |
|---|---|
| 0 | Success |
| 1 | Generic failure |
| 2 | Usage error |
| 3 | Authentication failure |
| 4 | Not found |
| 5 | Rate limited |
| 130 | Interrupted (Ctrl-C) |

### Running non-interactively

The CLI adapts to whether a human is watching. Progress animation is suppressed
automatically when stderr is not a terminal, when `TERM=dumb`, and under
`--quiet`, so CI logs stay clean without any flag. Colour additionally respects
`NO_COLOR` and `--no-color`.

| Flag | Effect |
|---|---|
| `--quiet` | Suppress progress and commentary. Results and errors still print. |
| `--no-color` | Disable colour. Status icons (`✓ ✗ ⚠`) remain. |
| `--no-input` | Never prompt. Fail with a message naming the flag to pass instead. |
| `--force` | Skip the confirmation on destructive commands. |

Destructive commands (`delete`, `void`, `cancel`, `terminate`, `archive`,
`finalize`) confirm before acting. In a script, pass `--force` to proceed:

    flexprice customers delete cust_123 --force

Without it, a non-interactive run **fails rather than proceeding**: the CLI
will not destroy something because nobody could be asked.

## Configuration

Non-secret settings live in `~/.flexprice/config.toml`; API keys live in your
OS keychain (or an encrypted file fallback where no keychain is available).
`FLEXPRICE_API_KEY` and `--api-key` override the stored key for a single
invocation or for CI.

## Full command reference

    flexprice <resource> --help

lists every action for a resource; a generated reference for every command is
published at https://docs.flexprice.io/cli.

## For maintainers

This CLI dispatches commands at runtime from an embedded OpenAPI spec rather
than generating Go source per command; [ARCHITECTURE.md](ARCHITECTURE.md)
walks through the request lifecycle end to end. The two most common
maintenance tasks (adding a generated command and adding a hand-written one)
are walked through in [guides/](guides/).

Build and test locally with the standard Go toolchain, or via `make`.
`make help` lists every target (build, test, lint, docs, smoke/e2e suites,
release):

    go build ./...
    go test -race ./...
    make check   # build, vet, test, lint, same as CI

## Contributing

Pull requests and issues are welcome directly against this repository. See
[CONTRIBUTING.md](CONTRIBUTING.md) for dev setup, testing, and linting, and
[CODE_OF_CONDUCT.md](CODE_OF_CONDUCT.md) for community standards.

## License

Apache-2.0. See [LICENSE](LICENSE).
