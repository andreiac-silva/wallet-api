tools: ## Install go tools.
	go install golang.org/x/tools/cmd/goimports@latest
	go install mvdan.cc/gofumpt@latest

imports: ## Format imports.
	goimports -l -w .

fmt: ## Format Code.
	gofumpt -l -w .

start: ## Run application by docker-compose.
	docker compose up -d --build

stop: ## Stop application by docker-compose.
	docker compose down -v

db-migrate: ## Run db migrations present in /migrations directory.
	docker run -ti --rm \
    	--name wallet-migrate \
    	--network wallet \
    	-v $(PWD)/migrations:/migrations \
    	migrate/migrate:v4.14.1 \
    	-path=/migrations/ -database $(MONGO_URL) up