## flexprice tax-associations create

Create Tax Association

### Synopsis

Create Tax Association

POST /taxes/associations

Fields you can set with flags:
  --auto_apply  [boolean]
  --currency  [string]
  --end_date  [string]
  --entity_id  [string]
  --entity_type  [string]
  --external_customer_id  [string]
  --priority  [integer]
  --start_date  [string]
  --tax_rate_code  (required)  [string]

Nested fields — these cannot be set with flags:
  metadata  [object]

Use --edit to fill in a pre-built request body, or --data @file.json.


```
flexprice tax-associations create [flags]
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

* [flexprice tax-associations](flexprice_tax-associations.md)	 - Which tax rates apply to which entities

