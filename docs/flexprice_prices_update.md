## flexprice prices update

Update price

### Synopsis

Update price

PUT /prices/{id}

Fields you can set with flags:
  --amount  [string]
  --billing_model  [string]
  --description  [string]
  --display_name  [string]
  --effective_from  [string]
  --group_id  [string]
  --lookup_key  [string]
  --price_unit_amount  [string]
  --tier_mode  [string]

Nested fields — these cannot be set with flags:
  metadata  [object]
  price_unit_tiers  [array]
  tiers  [array]
  transform_quantity  [object]

Use --edit to fill in a pre-built request body, or --data @file.json.


```
flexprice prices update [flags]
```

### Options

```
      --data string   request body: @file.json, - for stdin, or a JSON literal
      --edit          open $EDITOR with a pre-filled request body
      --force         skip the confirmation prompt on destructive actions
  -h, --help          help for update
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

