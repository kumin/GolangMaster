.PHONY: init
init:
	go install github.com/google/wire/cmd/wire@latest

wire: init
	wire ./...

build: wire $(DES)
	env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $(DES) ./cmd/server-ctl

