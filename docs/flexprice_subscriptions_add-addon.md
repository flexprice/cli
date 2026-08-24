## flexprice subscriptions add-addon

Add addon to subscription

### Synopsis

Add addon to subscription

POST /subscriptions/addon

Fields you can set with flags:
  --addon_id  (required)  [string]
  --cadence  [string]
  --proration_behavior  [string]
  --start_date  [string]
  --subscription_id  (required)  [string]

Nested fields — these cannot be set with flags:
  checkout  [object]
  line_item_commitments  [object]
  metadata  [object]
  override_line_items  [array]

Use --edit to fill in a pre-built request body, or --data @file.json.


```
flexprice subscriptions add-addon [flags]
```

### Options

```
      --data string   request body: @file.json, - for stdin, or a JSON literal
      --edit          open $EDITOR with a pre-filled request body
      --force         skip the confirmation prompt on destructive actions
  -h, --help          help for add-addon
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

