## flexprice prices list

Query prices

### Synopsis

Query prices

POST /prices/search

Fields you can set with flags:
  --allow_expired_prices  [boolean]
  --end_time  [string]
  --entity_type  [string]
  --expand  [string]
  --limit  [integer]
  --offset  [integer]
  --order  [string]
  --parent_price_id  [string]
  --sort  [string]
  --start_date_lt  [string]
  --start_time  [string]
  --status  [string]
  --subscription_id  [string]

Nested fields — these cannot be set with flags:
  billing_periods  [array]
  entity_ids  [array]
  filters  [array]
  meter_ids  [array]
  plan_ids  [array]
  price_ids  [array]

Use --edit to fill in a pre-built request body, or --data @file.json.


```
flexprice prices list [flags]
```

### Options

```
      --data string   request body: @file.json, - for stdin, or a JSON literal
      --edit          open $EDITOR with a pre-filled request body
      --force         skip the confirmation prompt on destructive actions
  -h, --help          help for list
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

* [flexprice prices](flexprice_prices.md)	 - Individual pricing units within a plan

