up: ## start server
	@docker-compose -f docker-compose.yaml up -d --build --force-recreate
.PHONY: up

down: ## shut down server
	@docker-compose -f docker-compose.yaml down --remove-orphans
.PHONY: down


init: ## shut down server
	@docker-compose -f docker-compose-init.yaml up -d --build --force-recreate
	@sleep 5
	@docker-compose -f docker-compose-init.yaml down --remove-orphans
.PHONY: down
