run:
	@go run cmd/main.go -FILE=$(NMAX)

test:
	@go test ./...

