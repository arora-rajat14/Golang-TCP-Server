.PHONY: build run test clean

build:
	go build -o bin/server ./cmd/server/main.go

run:
	go run ./cmd/server/main.go 9090

test:
	go test ./... -v

clean:
	rm -rf bin/
