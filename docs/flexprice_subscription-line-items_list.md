## flexprice subscription-line-items list

Search subscription line items

### Synopsis

Search subscription line items

POST /subscriptions/lineitems/search

Fields you can set with flags:
  --active_filter  [boolean]
  --current_period_start  [string]
  --end_time  [string]
  --entity_type  [string]
  --expand  [string]
  --limit  [integer]
  --offset  [integer]
  --order  [string]
  --start_time  [string]
  --status  [string]

Nested fields — these cannot be set with flags:
  addon_association_ids  [array]
  billing_periods  [array]
  currencies  [array]
  customer_ids  [array]
  entity_ids  [array]
  filters  [array]
  meter_ids  [array]
  price_ids  [array]
  sort  [array]
  subscription_ids  [array]
  subscription_line_item_ids  [array]

Use --edit to fill in a pre-built request body, or --data @file.json.


```
flexprice subscription-line-items list [flags]
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

* [flexprice subscription-line-items](flexprice_subscription-line-items.md)	 - Individual charges on a subscription

