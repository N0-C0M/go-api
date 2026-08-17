.PHONY: build
build:
	go build -v ./cmd/apiserver

.PHONY: apiserver
apiserver:
	go build -v -o apiserver.exe ./cmd/apiserver

.PHONY: test
test:
	go test -v -race -timeout 10s ./...

.DEFAULT_GOAL := build
