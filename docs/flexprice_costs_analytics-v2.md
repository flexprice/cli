## flexprice costs analytics-v2

Get combined revenue and cost analytics (V2)

### Synopsis

Get combined revenue and cost analytics (V2)

POST /costs/analytics-v2

Fields you can set with flags:
  --end_time  [string]
  --external_customer_id  [string]
  --include_children  [boolean]
  --limit  [integer]
  --offset  [integer]
  --start_time  [string]

Nested fields — these cannot be set with flags:
  expand  [array]
  feature_ids  [array]
  property_filters  [object]

Use --edit to fill in a pre-built request body, or --data @file.json.


```
flexprice costs analytics-v2 [flags]
```

### Options

```
      --data string   request body: @file.json, - for stdin, or a JSON literal
      --edit          open $EDITOR with a pre-filled request body
      --force         skip the confirmation prompt on destructive actions
  -h, --help          help for analytics-v2
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

* [flexprice costs](flexprice_costs.md)	 - Cost sheets derived from usage

