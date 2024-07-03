build:
    @go build -o bin/chirpx cmd/main.go

test:
    @go test -v ./...

run: build
    @./bin/chirpx
