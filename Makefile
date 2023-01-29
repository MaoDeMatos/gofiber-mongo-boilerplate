# A Self-Documenting Makefile: http://marmelab.com/blog/2016/02/29/auto-documented-makefile.html

# Variables
BUILD_PATH = ./tmp
MONGO_CONTAINER_NAME = mongodb

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
	go build -o ${BUILD_PATH}/main ./cmd/api

.PHONY: start
start: ## Run Production build
	${BUILD_PATH}/main

.PHONY: clear
clear: ## Clear generated files
	rm -rf ${BUILD_PATH}

.PHONY: db-start
db-start: ## Start MongdoDB, in a Docker container
	docker run --rm -d --name ${MONGO_CONTAINER_NAME} -v mongodb-volume:/data/db -p 27017:27017 mongo

.PHONY: db-stop
db-stop: ## Stop MongdoDB container
	docker stop /${MONGO_CONTAINER_NAME}

# Self-Documenting part
.PHONY: help
.DEFAULT_GOAL := help
help:
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "-> \033[36m%-20s\033[0m %s\n", $$1, $$2}'
