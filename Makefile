.PHONY: default
default: usage

.PHONY: tidy
tidy: ## Run go mod tidy
	@go mod tidy -v

.PHONY: build
build: | tidy ## Build app
	go build -ldflags '-s -w' -o bin/sonoff-diy ./cmd/sonoff-diy/main.go

.PHONY: outdated-deps
outdated-deps: ## List outdated dependencies
	@go list -u -m -f '{{if not .Indirect}}{{if .Update}}{{.}}{{end}}{{end}}' all

.PHONY: usage
usage: ## List available targets
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL := usage
