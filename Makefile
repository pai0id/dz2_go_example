run:
	@go run cmd/main.go -FILE=$(FILE)

test:
	@go test ./...

