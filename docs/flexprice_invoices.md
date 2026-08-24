## flexprice invoices

Draft, finalize and void billing documents

### Options

```
  -h, --help   help for invoices
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

* [flexprice](flexprice.md)	 - Flexprice CLI — usage-based billing from your terminal
* [flexprice invoices create](flexprice_invoices_create.md)	 - Create one-off invoice
* [flexprice invoices customer-summary](flexprice_invoices_customer-summary.md)	 - Get customer invoice summary
* [flexprice invoices finalize](flexprice_invoices_finalize.md)	 - Finalize invoice
* [flexprice invoices list](flexprice_invoices_list.md)	 - Query invoices
* [flexprice invoices payment-attempt](flexprice_invoices_payment-attempt.md)	 - Attempt invoice payment
* [flexprice invoices payment-status](flexprice_invoices_payment-status.md)	 - Update invoice payment status
* [flexprice invoices pdf](flexprice_invoices_pdf.md)	 - Get invoice PDF
* [flexprice invoices preview](flexprice_invoices_preview.md)	 - Get invoice preview
* [flexprice invoices recalculate](flexprice_invoices_recalculate.md)	 - Recalculate draft invoice (v2)
* [flexprice invoices retrieve](flexprice_invoices_retrieve.md)	 - Get invoice
* [flexprice invoices trigger-comms-webhook](flexprice_invoices_trigger-comms-webhook.md)	 - Trigger invoice communication webhook
* [flexprice invoices update](flexprice_invoices_update.md)	 - Update invoice
* [flexprice invoices void](flexprice_invoices_void.md)	 - Void invoice

