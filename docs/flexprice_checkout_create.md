## flexprice checkout create

Create checkout session

### Synopsis

Create checkout session

POST /checkout/sessions

Fields you can set with flags:
  --action  (required)  [string]
  --cancel_url  [string]
  --customer_external_id  (required)  [string]
  --failure_url  [string]
  --idempotency_key  [string]
  --payment_provider  (required)  [string]
  --success_url  [string]

Nested fields — these cannot be set with flags:
  configuration  [object]
  metadata  [object]
  payment_provider_config  [object]

Use --edit to fill in a pre-built request body, or --data @file.json.


```
flexprice checkout create [flags]
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

* [flexprice checkout](flexprice_checkout.md)	 - Hosted checkout sessions

