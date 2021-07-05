.PHONY: build
build:
	go build -v ./cmd/mdlpservice

.DEFAULT_GOAL := build