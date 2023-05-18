SHELL := /usr/bin/env bash

up:
	docker-compose up -d

relay:
	go run ./cmd/relay/main.go

cli:
	go run ./cmd/cli/*

fmt:
	go mod tidy -compat=1.17
	gofmt -l -s -w .
