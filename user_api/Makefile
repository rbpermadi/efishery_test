GOOS	 ?= linux
GOARCH  = amd64
ODIR    = _output

export GO111MODULE ?= on

all: compile test

test:
	go test -v ./...

coverage:
	go tool cover -html=cover.out -o cover.html

bin:
	go build -o _output/user_api app/main.go

run: bin
	./_output/user_api

mod:
	go mod tidy

compile:
	GOOS=$(GOOS) GOARCH=$(GOARCH) CGO_ENABLED=0 go build -o $(ODIR)/user_api_docker app/main.go


