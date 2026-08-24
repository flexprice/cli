## flexprice events huggingface-inference

Get Hugging Face inference data

### Synopsis

Get Hugging Face inference data

POST /events/huggingface-inference

Nested fields — these cannot be set with flags:
  requestIds  (required)  [array]

Use --edit to fill in a pre-built request body, or --data @file.json.


```
flexprice events huggingface-inference [flags]
```

### Options

```
      --data string   request body: @file.json, - for stdin, or a JSON literal
      --edit          open $EDITOR with a pre-filled request body
      --force         skip the confirmation prompt on destructive actions
  -h, --help          help for huggingface-inference
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

