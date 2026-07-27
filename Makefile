.PHONY: fmt test vet build bench lint clean


fmt:
	go fmt ./...


test:
	go test ./...


vet:
	go vet ./...


build:
	go build ./...


bench:
	go test -bench=. ./tests/...


lint:
	golangci-lint run


clean:
	go clean
