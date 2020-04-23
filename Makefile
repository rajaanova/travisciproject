.PHONY: build
build:
	go build ./main.go

.PHONY: tests
tests:
	go test ./... -v
	
	
	