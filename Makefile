.PHONY: mocks

help: Makefile
	@echo " Choose a command run in "$(PROJECTNAME)":"
	@echo ""
	@find . -maxdepth 1 -type f \( -name Makefile -or -name "*.mk" \) -exec cat {} \+ | sed -n 's/^##//p' | column -t -s ':' |  sed -e 's/^/ /'

## run-example: runs simple complete flow example
run-example:
	@go run examples/complete-flow/main.go

## deps: start example deps
deps:
	@cd examples && docker-compose up -d

## mocks: generate mocks
mocks:
	@mockgen github.com/topfreegames/chat-auth Storage | sed 's/mock_chat_auth/mocks/' > mocks/storage.go
	@mockgen github.com/topfreegames/chat-auth Password | sed 's/mock_chat_auth/mocks/' > mocks/password.go

## test: run unit tests
test:
	@go test ./... -tags=unit
