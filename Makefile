LDFLAGS :=

race-flag:
	@$(eval LDFLAGS=$(LDFLAGS) -ldflags -race)

.PHONY: default
default: usage

.PHONY: tidy
tidy: ## Run go mod tidy
	@go mod tidy -v

#.PHONY: pack
#pack: ## Pack app resources
#	statik -f -src ./web/build

.PHONY: build
build: | tidy ## Build app
	go build -ldflags '-s -w' -o bin/sonoff-diy ./cmd/sonoff-diy/main.go

.PHONY: usage
usage: ## List available targets
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL := usage