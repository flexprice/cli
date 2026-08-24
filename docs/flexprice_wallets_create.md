## flexprice wallets create

Create a new wallet

### Synopsis

Create a new wallet

POST /wallets

Fields you can set with flags:
  --conversion_rate  [string]
  --currency  (required)  [string]
  --customer_id  [string]
  --description  [string]
  --external_customer_id  [string]
  --initial_credits_expiry_date_utc  [string]
  --initial_credits_to_load  [string]
  --initial_credits_to_load_expiry_date  [integer]
  --price_unit  [string]
  --topup_conversion_rate  [string]
  --wallet_type  [string]

Nested fields — these cannot be set with flags:
  alert_settings  [object]
  auto_topup  [object]
  metadata  [object]

Use --edit to fill in a pre-built request body, or --data @file.json.


```
flexprice wallets create [flags]
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

* [flexprice wallets](flexprice_wallets.md)	 - Prepaid credit balances held by a customer

