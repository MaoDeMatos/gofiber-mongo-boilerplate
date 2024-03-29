# A Self-Documenting Makefile: http://marmelab.com/blog/2016/02/29/auto-documented-makefile.html

# Variables
BUILD_PATH = ./tmp
MONGO_CONTAINER_NAME = mongodb

.PHONY: install
install: ## Install dependencies and setup .env file
	@echo ⬇️ Download project dependencies
	go mod download
	@echo ⬇️ Download air
	go install github.com/cosmtrek/air@latest
	@echo 📄 Create .env
	@cp .env.sample .env

.PHONY: dev
dev: ## Run in Development mode
	air

.PHONY: build
build: ## Build for Production
	go build -o ${BUILD_PATH}/main ./cmd/api

.PHONY: run
run: ## Run Production build
	${BUILD_PATH}/main

.PHONY: clear
clear: ## Clear generated files
	rm -rf ${BUILD_PATH}

.PHONY: db-create
db-create: ## Create MongdoDB Docker container
	docker run -d --name ${MONGO_CONTAINER_NAME} -v ${MONGO_CONTAINER_NAME}-volume:/data/db -v ${MONGO_CONTAINER_NAME}-volume:/data/configdb -p 27017:27017 mongo:latest

.PHONY: db-start
db-start: ## Start MongdoDB container
	docker start ${MONGO_CONTAINER_NAME}

.PHONY: db-stop
db-stop: ## Stop MongdoDB container
	docker stop ${MONGO_CONTAINER_NAME}

# Self-Documenting part
.PHONY: help
.DEFAULT_GOAL := help
help:
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "-> \033[36m%-20s\033[0m %s\n", $$1, $$2}'
