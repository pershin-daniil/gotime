lint:
	gofumpt -w .
	gci write . --skip-generated -s standard -s default
	golangci-lint run -v --fix --timeout=5m ./...

up:
	docker compose -f ./deploy/local/docker-compose.yml up -d

down:
	docker compose -f ./deploy/local/docker-compose.yml down

dev-tools:
	go install github.com/daixiang0/gci@v0.13.5
	go install mvdan.cc/gofumpt@v0.7.0
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.61.0

