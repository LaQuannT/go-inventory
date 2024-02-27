run:
	@go build -o bin/go-inventory && bin/go-inventory

test:
	@go test ./...
