tools: ## Install go tools
	go install golang.org/x/tools/cmd/goimports@latest
	go install mvdan.cc/gofumpt@latest

imports: ## Format imports
	goimports -l -w .

fmt: ## Format Code
	gofumpt -l -w .

start: ## Run application by docker-compose
	docker compose up -d --build

stop: ## Stop application by docker-compose
	docker compose down -v