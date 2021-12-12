APP_NAME = ms-api
POSTGRESQL = postgresql
DOCKER_STORAGE_PATH = ~/ms-tv

build: ## Build on local
	swag init
	go build main.go

run: ## Run on local
	swag init
	go run main.go

docker-build: ## Build on docker
	docker build -t $(APP_NAME) .

docker-run: ## Run on docker
	docker run --rm \
		-p 8080:8080 \
		-v $(DOCKER_STORAGE_PATH)/assets:/go/src/$(APP_NAME)/assets \
		-v $(DOCKER_STORAGE_PATH)/logs:/go/src/$(APP_NAME)/logs \
		--name $(APP_NAME) \
		$(APP_NAME):latest

external-run: ## Run external apps
	docker run -d --rm \
		-p 5432:5432 \
		-e POSTGRES_DB=video \
		-e POSTGRES_USER=root \
		-e POSTGRES_PASSWORD=root \
		-e POSTGRES_INITDB_ARGS="--encoding=UTF-8" \
		-v $(PWD)/db/:/docker-entrypoint-initdb.d \
		--name $(POSTGRESQL) \
		postgres:latest

external-end: ## End external apps
	docker stop $(POSTGRESQL)

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
