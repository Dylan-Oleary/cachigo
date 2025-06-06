.PHONY: cache
cache:
	@go run ./cmd/cache/main.go

.PHONY: tcp
tcp:
	@go run ./tcp/server/main.go

