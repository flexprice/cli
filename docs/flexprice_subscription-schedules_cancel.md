## flexprice subscription-schedules cancel

Cancel subscription schedule

### Synopsis

Cancel subscription schedule

POST /v1/subscriptions/schedules/{schedule_id}/cancel

Fields you can set with flags:
  --schedule_id  [string]
  --schedule_type  [string]
  --subscription_id  [string]


```
flexprice subscription-schedules cancel [flags]
```

### Options

```
      --data string   request body: @file.json, - for stdin, or a JSON literal
      --edit          open $EDITOR with a pre-filled request body
      --force         skip the confirmation prompt on destructive actions
  -h, --help          help for cancel
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

* [flexprice subscription-schedules](flexprice_subscription-schedules.md)	 - Planned future changes to a subscription

