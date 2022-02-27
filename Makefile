APP_NAME = ms-api
EXTERNAL_APPS = postgresql ms-recommendation-api

init: ## Initialize app
	go mod download
	go mod tidy
	mkdir -p ./logs/

run: ## Run on local
	swag init
	APP_ENV=development go run main.go

docker-build: ## Build on docker
	docker build -t $(APP_NAME) .

docker-run: ## Run on docker
	docker run --rm \
		--add-host=localhost:host-gateway \
		-p 8080:8080 \
		-e APP_ENV=development \
		-v assets:/go/src/$(APP_NAME)/assets \
		-v logs:/go/src/$(APP_NAME)/logs \
		--name $(APP_NAME) \
		$(APP_NAME):latest

external-init: ## Initialize external apps
	rm -rf ./external-apps/logs

external-run: ## Run external apps
	docker-compose up -d $(EXTERNAL_APPS)

external-end: ## End external apps
	docker-compose down

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
