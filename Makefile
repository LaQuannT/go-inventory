test:
	@go test -v ./...

build:
	@go build -o go-inventory && go install

docker-compose:
	@docker-compose up -d


