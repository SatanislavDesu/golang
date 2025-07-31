.DEFAULT_GOAL := clean

.PHONY:fmt vet build run clean
fmt:
	go fmt ./...

vet: fmt
	go vet ./...

build: vet
	go build -o hello

run: build
	./hello

clean: run
	go clean -i

# run: build
# 	go run ./hello
# Хз почему но ран не работает