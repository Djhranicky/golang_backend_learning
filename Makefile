build:
	@go build -o bin/GO_BACKEND cmd/main.go

test:
	@go test -v ./...

run: build
	@./bin/GO_BACKEND