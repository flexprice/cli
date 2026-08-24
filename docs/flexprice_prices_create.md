## flexprice prices create

Create price

### Synopsis

Create price

POST /prices

Fields you can set with flags:
  --amount  [string]
  --billing_model  (required)  [string]
  --billing_period  (required)  [string]
  --billing_period_count  [integer]
  --currency  (required)  [string]
  --description  [string]
  --display_name  [string]
  --end_date  [string]
  --entity_id  (required)  [string]
  --entity_type  (required)  [string]
  --group_id  [string]
  --invoice_cadence  (required)  [string]
  --lookup_key  [string]
  --meter_id  [string]
  --min_quantity  [integer]
  --price_unit_type  (required)  [string]
  --start_date  [string]
  --tier_mode  [string]
  --trial_period_days  [integer]
  --type  (required)  [string]

Nested fields — these cannot be set with flags:
  filter_values  [object]
  metadata  [object]
  price_unit_config  [object]
  tiers  [array]
  transform_quantity  [object]

Use --edit to fill in a pre-built request body, or --data @file.json.


```
flexprice prices create [flags]
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

* [flexprice prices](flexprice_prices.md)	 - Individual pricing units within a plan

