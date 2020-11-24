.PHONY: run-scanner

run: 
	go run -race ./cmd/sectools/main.go

build:
	go build -o bin/main ./cmd/sectools/main.go