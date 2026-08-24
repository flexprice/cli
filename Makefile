.DEFAULT_GOAL := help
BIN := bin/flexprice

.PHONY: build install run test lint lint-fix fmt vet tidy docs bootstrap-commands \
        smoke e2e release-snapshot release clean check help

build: ## Build the flexprice binary into bin/
	go build -o $(BIN) .

install: ## Install flexprice to $GOBIN (or $GOPATH/bin)
	go install .

run: build ## Build and run the binary, e.g. make run ARGS="customers list"
	$(BIN) $(ARGS)

test: ## Run the test suite with the race detector
	go test -race ./...

lint: ## Run golangci-lint
	golangci-lint run ./...

lint-fix: ## Run golangci-lint and apply auto-fixes where possible
	golangci-lint run --fix ./...

fmt: ## Format all Go source with gofmt
	gofmt -l -w .

vet: ## Run go vet
	go vet ./...

tidy: ## Tidy go.mod/go.sum
	go mod tidy

docs: ## Regenerate the command reference under docs/
	go run ./tools/gendocs

bootstrap-commands: ## Print operations from the embedded spec missing a name in spec/commands.yaml
	go run ./tools/bootstrap-commands

smoke: build ## Run the smoke suite (no API key needed; set FLEXPRICE_API_KEY for live phases)
	./scripts/smoke.sh

e2e: build ## Run the full lifecycle suite against the India region (needs FLEXPRICE_API_KEY)
	./scripts/e2e.sh

release-snapshot: ## Build release binaries locally without publishing (needs goreleaser)
	goreleaser release --snapshot --clean

release: ## Cut a real release — run from a v* tag, needs goreleaser and GITHUB_TOKEN
	goreleaser release --clean

clean: ## Remove build artifacts (bin/, dist/, generated docs)
	rm -rf bin dist
	find docs -maxdepth 1 -name '*.md' -delete

check: build vet test lint ## Run everything CI runs: build, vet, test, lint

help: ## Show this help
	@grep -E '^[a-zA-Z_-]+:.*## ' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*## "}; {printf "  \033[36m%-20s\033[0m %s\n", $$1, $$2}'
