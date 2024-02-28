run:
	@go build -o bin/go-inventory cmd/go-inventory/main.go && bin/go-inventory

build:
	@go build -o bin/go-inventory cmd/go-inventory/main.go

test:
	@go test ./...

docker-compose:
	@docker-compose up -d
