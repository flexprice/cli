## flexprice login

Store credentials for a region and environment

### Synopsis

Verifies your API key, resolves the tenant and environment it is scoped to,
and stores it in your OS keychain.

An API key belongs to exactly one environment, so use one profile per environment.

```
flexprice login [flags]
```

### Options

```
  -h, --help                  help for login
      --label string          free-text note shown by whoami, e.g. "sandbox"
      --profile-name string   name for the stored profile (default: "default")
```

### Options inherited from parent commands

```
      --all               page through every record (prints the last page; use --output json with --limit for bulk export)
      --api-key string    API key (CI use; prefer flexprice login)
      --base-url string   override the API base URL
      --columns strings   columns to show in table output
      --debug             dump requests and responses, secrets redacted
      --limit int         maximum records to return (default 20)
      --no-color          disable coloured output
      --no-input          never prompt; fail instead of asking
      --output string     output format: table, json, yaml (default "table")
  -p, --profile string    profile to use for this command
      --quiet             suppress progress output
      --region string     region key, e.g. us or in
```

### SEE ALSO

* [flexprice](flexprice.md)	 - Flexprice CLI — usage-based billing from your terminal

