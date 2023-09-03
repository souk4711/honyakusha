GIT_COMMIT := $(shell git rev-list -1 HEAD)
BUILT_TIME := $(shell date +%s)
LDFLAGS := -X 'main.GitCommit=${GIT_COMMIT}' -X 'main.BuiltTime=${BUILT_TIME}'

.PHONY: default fmt lint generate tidy test build clean

default: fmt lint generate tidy test build

fmt:
	gofmt -s -w .

lint:
	golangci-lint run

generate:
	go generate ./...

tidy:
	go mod tidy

test:
	go test -v ./... -count=1

build:
	go build -ldflags "$(LDFLAGS)" -o ./bin/honyakusha ./cmd/honyakusha/

clean:
	go clean
	rm -f ./bin/honyakusha
