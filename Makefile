.PHONY: cli
cli:
	@go run ./cmd/cli/main.go

.PHONY: server
server:
	@go run ./cmd/server/main.go

