## flexprice subscriptions change-preview

Preview subscription plan change

### Synopsis

Preview subscription plan change

POST /subscriptions/{id}/change/preview

Fields you can set with flags:
  --billing_cadence  (required)  [string]
  --billing_cycle  (required)  [string]
  --billing_period  (required)  [string]
  --billing_period_count  [integer]
  --change_at  [string]
  --proration_behavior  (required)  [string]
  --target_plan_id  (required)  [string]

Nested fields — these cannot be set with flags:
  metadata  [object]

Use --edit to fill in a pre-built request body, or --data @file.json.


```
flexprice subscriptions change-preview [flags]
```

### Options

```
      --data string   request body: @file.json, - for stdin, or a JSON literal
      --edit          open $EDITOR with a pre-filled request body
      --force         skip the confirmation prompt on destructive actions
  -h, --help          help for change-preview
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

* [flexprice subscriptions](flexprice_subscriptions.md)	 - Active plan assignments and their lifecycle

