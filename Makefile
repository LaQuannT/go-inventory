run:
	@go build -o bin/go-inventory cmd/cli/main.go && bin/go-inventory

build:
	@go build -o bin/go-inventory cmd/cli/main.go

test:
	@go test ./...

docker-compose:
	@docker-compose up -d
