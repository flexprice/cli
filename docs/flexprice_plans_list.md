## flexprice plans list

Query plans

### Synopsis

Query plans

POST /plans/search

Fields you can set with flags:
  --end_time  [string]
  --expand  [string]
  --limit  [integer]
  --lookup_key  [string]
  --offset  [integer]
  --order  [string]
  --start_time  [string]
  --status  [string]

Nested fields — these cannot be set with flags:
  filters  [array]
  metadata  [object]
  plan_ids  [array]
  sort  [array]

Use --edit to fill in a pre-built request body, or --data @file.json.


```
flexprice plans list [flags]
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

* [flexprice plans](flexprice_plans.md)	 - Pricing models customers can subscribe to

