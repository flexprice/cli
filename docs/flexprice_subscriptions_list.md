## flexprice subscriptions list

Query subscriptions

### Synopsis

Query subscriptions

POST /subscriptions/search

Fields you can set with flags:
  --active_at  [string]
  --customer_id  [string]
  --effective_date_for_update  [string]
  --end_time  [string]
  --expand  [string]
  --external_customer_id  [string]
  --limit  [integer]
  --offset  [integer]
  --order  [string]
  --plan_id  [string]
  --start_time  [string]
  --status  [string]
  --trial_end_due_lte  [string]
  --with_coupon_associations  [boolean]
  --with_line_items  [boolean]

Nested fields — these cannot be set with flags:
  billing_cadence  [array]
  billing_period  [array]
  customer_ids  [array]
  filters  [array]
  invoicing_customer_ids  [array]
  parent_subscription_ids  [array]
  sort  [array]
  subscription_ids  [array]
  subscription_status  [array]
  subscription_type  [array]

Use --edit to fill in a pre-built request body, or --data @file.json.


```
flexprice subscriptions list [flags]
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

* [flexprice subscriptions](flexprice_subscriptions.md)	 - Active plan assignments and their lifecycle

