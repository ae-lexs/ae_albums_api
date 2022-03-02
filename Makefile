coverage:
	go test ./... -coverprofile=coverage.txt -covermode=atomic

format:
	gofmt -w ./..

install:
	go get ./...

lint:
	gofmt -d ./..

test:
	go test -v ./...
