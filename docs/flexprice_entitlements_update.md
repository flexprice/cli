## flexprice entitlements update

Update entitlement

### Synopsis

Update entitlement

PUT /entitlements/{id}

Fields you can set with flags:
  --aggregation_mode  [string]
  --clear_grant_config  [boolean]
  --grant_allocation_behavior  [string]
  --grant_duration_unit  [string]
  --grant_duration_value  [integer]
  --grant_measure  [string]
  --grant_quota  [string]
  --is_enabled  [boolean]
  --is_soft_limit  [boolean]
  --static_value  [string]
  --usage_limit  [integer]
  --usage_reset_period  [string]

Nested fields — these cannot be set with flags:
  config_value  [object]

Use --edit to fill in a pre-built request body, or --data @file.json.


```
flexprice entitlements update [flags]
```

### Options

```
      --data string   request body: @file.json, - for stdin, or a JSON literal
      --edit          open $EDITOR with a pre-filled request body
      --force         skip the confirmation prompt on destructive actions
  -h, --help          help for update
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

* [flexprice entitlements](flexprice_entitlements.md)	 - What a customer's plan grants them access to

