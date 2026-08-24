## flexprice credit-grants create

Create credit grant

### Synopsis

Create credit grant

POST /creditgrants

Fields you can set with flags:
  --addon_id  [string]
  --cadence  (required)  [string]
  --conversion_rate  [string]
  --credits  (required)  [string]
  --end_date  [string]
  --expiration_duration  [integer]
  --expiration_duration_unit  [string]
  --expiration_type  [string]
  --name  (required)  [string]
  --period  [string]
  --period_count  [integer]
  --plan_id  [string]
  --priority  [integer]
  --scope  (required)  [string]
  --start_date  [string]
  --subscription_id  [string]
  --topup_conversion_rate  [string]

Nested fields — these cannot be set with flags:
  metadata  [object]

Use --edit to fill in a pre-built request body, or --data @file.json.


```
flexprice credit-grants create [flags]
```

### Options

```
      --data string   request body: @file.json, - for stdin, or a JSON literal
      --edit          open $EDITOR with a pre-filled request body
      --force         skip the confirmation prompt on destructive actions
  -h, --help          help for create
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

* [flexprice credit-grants](flexprice_credit-grants.md)	 - Prepaid and promotional credit allocations

