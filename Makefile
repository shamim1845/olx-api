.PHONY: build run clean runws

build:
	@go build -o bin/api/main cmd/api/main.go
	@go build -o bin/ws/main cmd/ws/main.go

run: build
	@./bin/api/main

runws: build
	@./bin/ws/main

clean:
	@rm -rf bin