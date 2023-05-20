gozip/run:
	@go run cmd/gozip/main.go

gozip/build:
	@go build -o bin/gozip cmd/gozip/main.go

gozip/clean:
	@go clean