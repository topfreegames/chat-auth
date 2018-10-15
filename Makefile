run-example:
	@go run examples/complete-flow/main.go

deps:
	@cd examples && docker-compose up -d
