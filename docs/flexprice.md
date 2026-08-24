## flexprice

Flexprice CLI — usage-based billing from your terminal

### Synopsis

Send events, inspect how they metered, and drive the Flexprice API from your terminal.

Start with: flexprice init

```
flexprice [flags]
```

### Options

```
      --all               page through every record (prints the last page; use --output json with --limit for bulk export)
      --api-key string    API key (CI use; prefer flexprice login)
      --base-url string   override the API base URL
      --columns strings   columns to show in table output
      --debug             dump requests and responses, secrets redacted
  -h, --help              help for flexprice
      --limit int         maximum records to return (default 20)
      --no-color          disable coloured output
      --no-input          never prompt; fail instead of asking
      --output string     output format: table, json, yaml (default "table")
  -p, --profile string    profile to use for this command
      --quiet             suppress progress output
      --region string     region key, e.g. us or in
```

### SEE ALSO

* [flexprice addons](flexprice_addons.md)	 - Optional extras attachable to a plan
* [flexprice alert-settings](flexprice_alert-settings.md)	 - How and when alerts fire
* [flexprice alerts](flexprice_alerts.md)	 - Threshold and anomaly notifications
* [flexprice checkout](flexprice_checkout.md)	 - Hosted checkout sessions
* [flexprice config](flexprice_config.md)	 - Manage profiles
* [flexprice costs](flexprice_costs.md)	 - Cost sheets derived from usage
* [flexprice coupon-associations](flexprice_coupon-associations.md)	 - Which coupons apply to which subscriptions
* [flexprice coupons](flexprice_coupons.md)	 - Discount codes and their rules
* [flexprice credit-grants](flexprice_credit-grants.md)	 - Prepaid and promotional credit allocations
* [flexprice credit-notes](flexprice_credit-notes.md)	 - Refunds and credit memos against invoices
* [flexprice customers](flexprice_customers.md)	 - Manage the people and organisations you bill
* [flexprice delete](flexprice_delete.md)	 - Issue a raw DELETE against the API
* [flexprice entitlements](flexprice_entitlements.md)	 - What a customer's plan grants them access to
* [flexprice env](flexprice_env.md)	 - Inspect environments
* [flexprice environments](flexprice_environments.md)	 - Isolated spaces within your tenant
* [flexprice events](flexprice_events.md)	 - Raw usage events you send in for metering
* [flexprice features](flexprice_features.md)	 - Capabilities that can be metered or gated
* [flexprice get](flexprice_get.md)	 - Issue a raw GET against the API
* [flexprice groups](flexprice_groups.md)	 - Collections of users or entities
* [flexprice init](flexprice_init.md)	 - Set up the CLI (guided)
* [flexprice integrations](flexprice_integrations.md)	 - Connections to Stripe, HubSpot and others
* [flexprice invoices](flexprice_invoices.md)	 - Draft, finalize and void billing documents
* [flexprice login](flexprice_login.md)	 - Store credentials for a region and environment
* [flexprice logout](flexprice_logout.md)	 - Remove a stored profile and its key
* [flexprice open](flexprice_open.md)	 - Open Flexprice in your browser
* [flexprice payments](flexprice_payments.md)	 - Payment attempts and their outcomes
* [flexprice plans](flexprice_plans.md)	 - Pricing models customers can subscribe to
* [flexprice post](flexprice_post.md)	 - Issue a raw POST against the API
* [flexprice price-units](flexprice_price-units.md)	 - Units of measurement used by prices
* [flexprice prices](flexprice_prices.md)	 - Individual pricing units within a plan
* [flexprice rbac](flexprice_rbac.md)	 - Roles and permissions
* [flexprice resources](flexprice_resources.md)	 - List every resource this CLI can act on
* [flexprice scheduled-tasks](flexprice_scheduled-tasks.md)	 - Work queued to run later
* [flexprice secrets](flexprice_secrets.md)	 - API keys and integration credentials
* [flexprice subscription-line-items](flexprice_subscription-line-items.md)	 - Individual charges on a subscription
* [flexprice subscription-schedules](flexprice_subscription-schedules.md)	 - Planned future changes to a subscription
* [flexprice subscriptions](flexprice_subscriptions.md)	 - Active plan assignments and their lifecycle
* [flexprice tasks](flexprice_tasks.md)	 - Background jobs and their status
* [flexprice tax-associations](flexprice_tax-associations.md)	 - Which tax rates apply to which entities
* [flexprice tax-rates](flexprice_tax-rates.md)	 - Tax rates available to apply
* [flexprice tenants](flexprice_tenants.md)	 - Your top-level account
* [flexprice users](flexprice_users.md)	 - People with access to your tenant
* [flexprice version](flexprice_version.md)	 - Print the CLI version and embedded spec build
* [flexprice wallets](flexprice_wallets.md)	 - Prepaid credit balances held by a customer
* [flexprice whoami](flexprice_whoami.md)	 - Show the active profile, environment and key backend
* [flexprice workflows](flexprice_workflows.md)	 - Long-running billing processes

