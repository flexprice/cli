## flexprice events usage-by-meter

Get usage by meter

### Synopsis

Get usage by meter

POST /events/usage/meter

Fields you can set with flags:
  --billing_anchor  [string]
  --bucket_size  [string]
  --customer_id  [string]
  --end_time  [string]
  --external_customer_id  [string]
  --meter_id  (required)  [string]
  --start_time  [string]
  --timezone  [string]
  --window_size  [string]

Nested fields — these cannot be set with flags:
  filters  [object]

Use --edit to fill in a pre-built request body, or --data @file.json.


```
flexprice events usage-by-meter [flags]
```

### Options

```
      --data string   request body: @file.json, - for stdin, or a JSON literal
      --edit          open $EDITOR with a pre-filled request body
      --force         skip the confirmation prompt on destructive actions
  -h, --help          help for usage-by-meter
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

* [flexprice events](flexprice_events.md)	 - Raw usage events you send in for metering

