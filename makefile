.PHONY: build
build:
	go build -v ./cmd/diplom  && ./diplom

.DEFAULT_GOAL := build