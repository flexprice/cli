## flexprice subscriptions create

Create subscription

### Synopsis

Create subscription

POST /subscriptions

Fields you can set with flags:
  --auto_invoice_threshold  [string]
  --billing_anchor  [string]
  --billing_cycle  [string]
  --billing_period  (required)  [string]
  --billing_period_count  [integer]
  --collection_method  [string]
  --commitment_amount  [string]
  --commitment_duration  [string]
  --currency  (required)  [string]
  --customer_id  [string]
  --enable_true_up  [boolean]
  --end_date  [string]
  --external_customer_id  [string]
  --gateway_payment_method_id  [string]
  --lookup_key  [string]
  --overage_factor  [string]
  --payment_behavior  [string]
  --payment_terms  [string]
  --plan_id  (required)  [string]
  --proration_behavior  [string]
  --start_date  [string]
  --subscription_status  [string]
  --timezone  [string]
  --trial_period_days  [integer]

Nested fields — these cannot be set with flags:
  addons  [array]
  checkout  [object]
  coupons  [array]
  credit_grants  [array]
  inheritance  [object]
  line_item_commitments  [object]
  line_item_coupons  [object]
  line_items  [array]
  metadata  [object]
  override_entitlements  [array]
  override_line_items  [array]
  phases  [array]
  subscription_coupons  [array]
  tax_rate_overrides  [array]

Use --edit to fill in a pre-built request body, or --data @file.json.


```
flexprice subscriptions create [flags]
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

* [flexprice subscriptions](flexprice_subscriptions.md)	 - Active plan assignments and their lifecycle

