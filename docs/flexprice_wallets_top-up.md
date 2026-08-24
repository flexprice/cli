## flexprice wallets top-up

Top up wallet

### Synopsis

Top up wallet

POST /wallets/{id}/top-up

Fields you can set with flags:
  --amount  [string]
  --bonus_credits_expiry_date_utc  [string]
  --bonus_credits_to_add  [string]
  --credits_to_add  [string]
  --description  [string]
  --expiry_date_utc  [string]
  --force_sync_invoice  [boolean]
  --idempotency_key  [string]
  --priority  [integer]
  --transaction_reason  (required)  [string]

Nested fields — these cannot be set with flags:
  checkout  [object]
  metadata  [object]

Use --edit to fill in a pre-built request body, or --data @file.json.


```
flexprice wallets top-up [flags]
```

### Options

```
      --data string   request body: @file.json, - for stdin, or a JSON literal
      --edit          open $EDITOR with a pre-filled request body
      --force         skip the confirmation prompt on destructive actions
  -h, --help          help for top-up
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

