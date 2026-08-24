## flexprice invoices create

Create one-off invoice

### Synopsis

Create one-off invoice

POST /invoices

Fields you can set with flags:
  --amount_due  (required)  [string]
  --amount_paid  [string]
  --billing_period  [string]
  --billing_reason  [string]
  --currency  (required)  [string]
  --customer_id  (required)  [string]
  --description  [string]
  --due_date  [string]
  --force_sync_invoice  [boolean]
  --idempotency_key  [string]
  --invoice_number  [string]
  --invoice_pdf_url  [string]
  --invoice_status  [string]
  --invoice_type  [string]
  --issue_date  [string]
  --payment_status  [string]
  --period_end  [string]
  --period_start  [string]
  --subscription_id  [string]
  --subtotal  (required)  [string]
  --total  (required)  [string]
  --total_prepaid_applied  [string]

Nested fields — these cannot be set with flags:
  coupons  [array]
  invoice_coupons  [array]
  line_item_coupons  [array]
  line_items  [array]
  metadata  [object]
  prepared_tax_rates  [array]
  tax_rate_overrides  [array]
  tax_rates  [array]

Use --edit to fill in a pre-built request body, or --data @file.json.


```
flexprice invoices create [flags]
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

* [flexprice invoices](flexprice_invoices.md)	 - Draft, finalize and void billing documents

