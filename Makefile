.PHONY: build run clean runws

build:
	@go build -o bin/api/main.exe cmd/api/main.go
	@go build -o bin/ws/main.exe cmd/ws/main.go

run: build
	@./bin/api/main.exe

runws: build
	@./bin/ws/main.exe

clean:
	@rmdir /s /q bin