test:
	go test -v -race ./...

lint:
	go get -u github.com/golangci/golangci-lint/cmd/golangci-lint
	golangci-lint run

check: lint test

cover:
	go test -race -cover -coverprofile=cover.out ./...
	go tool cover -html=cover.out

build:
	export GO111MODULE=off && \
	go mod download && \
	go build -v .