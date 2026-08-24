## flexprice wallets search-transactions

Query wallet transactions

### Synopsis

Query wallet transactions

POST /wallets/transactions/search

Fields you can set with flags:
  --created_by  [string]
  --credits_available_gt  [number]
  --end_time  [string]
  --expand  [string]
  --expiry_date_after  [string]
  --expiry_date_before  [string]
  --id  [string]
  --limit  [integer]
  --offset  [integer]
  --order  [string]
  --priority  [integer]
  --reference_id  [string]
  --reference_type  [string]
  --start_time  [string]
  --status  [string]
  --transaction_reason  [string]
  --transaction_status  [string]
  --type  [string]

Nested fields — these cannot be set with flags:
  filters  [array]
  sort  [array]

Use --edit to fill in a pre-built request body, or --data @file.json.


```
flexprice wallets search-transactions [flags]
```

### Options

```
      --data string   request body: @file.json, - for stdin, or a JSON literal
      --edit          open $EDITOR with a pre-filled request body
      --force         skip the confirmation prompt on destructive actions
  -h, --help          help for search-transactions
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

* [flexprice wallets](flexprice_wallets.md)	 - Prepaid credit balances held by a customer

