## flexprice invoices list

Query invoices

### Synopsis

Query invoices

POST /invoices/search

Fields you can set with flags:
  --amount_due_gt  [number]
  --amount_remaining_gt  [number]
  --billing_reason  [string]
  --currency  [string]
  --customer_id  [string]
  --end_time  [string]
  --expand  [string]
  --external_customer_id  [string]
  --invoice_type  [string]
  --limit  [integer]
  --offset  [integer]
  --order  [string]
  --period_end_gte  [string]
  --period_end_lte  [string]
  --period_start_gte  [string]
  --period_start_lte  [string]
  --skip_line_items  [boolean]
  --start_time  [string]
  --status  [string]
  --subscription_id  [string]

Nested fields — these cannot be set with flags:
  filters  [array]
  invoice_ids  [array]
  invoice_status  [array]
  payment_status  [array]
  sort  [array]
  subscription_customer_id  [array]

Use --edit to fill in a pre-built request body, or --data @file.json.


```
flexprice invoices list [flags]
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

* [flexprice invoices](flexprice_invoices.md)	 - Draft, finalize and void billing documents

