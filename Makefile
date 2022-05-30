.PHONY:
build:
	go build -v -race ./cmd/apiserver

run: build
	./apiserver

.DEFAULT_GOAL := build