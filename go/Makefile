# Self documented Makefile
# http://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
help: ## Show list of make targets and their description
	@grep -E '^[/%.a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) \
		| awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL:= help

.PHONY: setup
setup: ## Run setup scripts to prepare development environment
	@scripts/setup.sh

.PHONY: test
test: ## Generate mock and run all test. To run specified tests, use `./scripts/test.sh <pattern>`)
	@scripts/test.sh $*

.PHONY: gen
gen: ## Show gen.sh help
	@scripts/gen.sh

.PHONY: build
build: ## Show build.sh help for building binnary package under cmd
	@scripts/build.sh

build/%: ## Build artifact defined by '%', e.g: 'make build.server` will trigger ./scripts/build.sh server
	@scripts/build.sh $*

all: setup gen build/all test ## Clean, setup, generate and then build all the binaries.

