#project name
APP_NAME = cli-todo

.PHONY: fmt vet build run clean

fmt: 
	go fmt ./...

vet: fmt
	go vet ./...

build: vet
	go build -o $(APP_NAME) main.go

run: build
	./$(APP_NAME)

clean: 
	go clean
	rm -f $(APP_NAME)


