tools: ## Install go tools
	go install golang.org/x/tools/cmd/goimports@latest
	go install mvdan.cc/gofumpt@latest

imports: ## Format imports
	goimports -l -w .

fmt: ## Format Code
	gofumpt -l -w .

lint: ## Run linter
	docker run --rm -v $(PWD):/app -w /app golangci/golangci-lint:v1.49.0 golangci-lint run -v

start: ## Run application by docker-compose
	docker compose up -d --build

stop: ## Stop application by docker-compose
	docker compose down -v