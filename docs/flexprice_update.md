## flexprice update

Update the CLI to the latest release

### Synopsis

Check GitHub for a newer release and, if there is one, run the install script (https://cli.flexprice.io) to replace this binary in place.

Every other command checks once a day and prints a notice when a newer release exists; set FLEXPRICE_NO_UPDATE=1 to turn that off.

```
flexprice update [flags]
```

### Options

```
      --check   report whether a newer release exists without installing it
  -h, --help    help for update
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
