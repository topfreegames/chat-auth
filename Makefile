.PHONY: mocks

run-example:
	@go run examples/complete-flow/main.go

deps:
	@cd examples && docker-compose up -d

mocks:
	@mockgen github.com/topfreegames/chat-auth Storage | sed 's/mock_chat_auth/mocks/' > mocks/storage.go
	@mockgen github.com/topfreegames/chat-auth Password | sed 's/mock_chat_auth/mocks/' > mocks/password.go

test:
	@go test ./...
