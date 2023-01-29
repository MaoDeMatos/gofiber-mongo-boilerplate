# A Self-Documenting Makefile: http://marmelab.com/blog/2016/02/29/auto-documented-makefile.html

.PHONY: install
install: ## Install dependencies and setup .env file
	@go mod download
	@go install github.com/cosmtrek/air@latest
	@cp .env.sample .env

.PHONY: dev
dev: ## Run in Development mode
	air

.PHONY: build
build: ## Build for Production
	go build -o ./tmp/main ./cmd/api

.PHONY: start
start: ## Run Production build
	./tmp/main

.PHONY: clear
clear: ## Clear generated files
	rm -rf tmp/

# Self-Documenting part
.PHONY: help
.DEFAULT_GOAL := help
help:
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "-> \033[36m%-20s\033[0m %s\n", $$1, $$2}'
