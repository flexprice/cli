## flexprice subscription-line-items update

Update subscription line item

### Synopsis

Update subscription line item

PUT /subscriptions/lineitems/{id}

Fields you can set with flags:
  --amount  [string]
  --billing_model  [string]
  --commitment_amount  [number]
  --commitment_duration  [string]
  --commitment_overage_factor  [number]
  --commitment_quantity  [number]
  --commitment_true_up_enabled  [boolean]
  --commitment_type  [string]
  --commitment_windowed  [boolean]
  --effective_from  [string]
  --tier_mode  [string]

Nested fields — these cannot be set with flags:
  commitment_time_buckets  [array]
  metadata  [object]
  tiers  [array]
  transform_quantity  [object]

Use --edit to fill in a pre-built request body, or --data @file.json.


```
flexprice subscription-line-items update [flags]
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

* [flexprice subscription-line-items](flexprice_subscription-line-items.md)	 - Individual charges on a subscription

