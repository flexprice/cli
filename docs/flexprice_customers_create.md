## flexprice customers create

Create customer

### Synopsis

Create customer

POST /customers

Fields you can set with flags:
  --address_city  [string]
  --address_country  [string]
  --address_line1  [string]
  --address_line2  [string]
  --address_postal_code  [string]
  --address_state  [string]
  --contact  [string]
  --email  [string]
  --external_id  (required)  [string]
  --name  (required)  [string]
  --onboarding_workflow_name  [string]
  --skip_onboarding_workflow  [boolean]
  --timezone  [string]

Nested fields — these cannot be set with flags:
  integration_entity_mapping  [array]
  metadata  [object]
  tax_rate_overrides  [array]

Use --edit to fill in a pre-built request body, or --data @file.json.


```
flexprice customers create [flags]
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

* [flexprice customers](flexprice_customers.md)	 - Manage the people and organisations you bill

