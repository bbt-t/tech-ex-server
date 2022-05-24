.PHONY:
build:
	go build -v ./cmd/apiserver

run: build
	./apiserver

.DEFAULT_GOAL := build