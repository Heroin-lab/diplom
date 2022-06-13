.PHONY: build
build:
	   go build -v ./cmd/taxi_service && ./taxi_service

.DEFAULT_GOAL := build