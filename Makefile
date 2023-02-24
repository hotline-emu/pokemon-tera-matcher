default:
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

docker-start: ## Starts the docker-compose services.
	docker-compose start

docker-stop: ## Starts the docker-compose services.
	docker-compose stop

api: ## Run the api
	go run server.go
