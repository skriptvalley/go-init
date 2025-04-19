run:
	go run ./cmd/server/main.go

build:
	go build -o bin/server ./cmd/server

test:
	go test ./...

lint:
	golangci-lint run

fmt:
	gofmt -w .

