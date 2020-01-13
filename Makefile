PACKAGES := $(shell go list ./... | grep -v cmd | grep -v test_files)

.PHONY: test
test:
	go test -v -failfast -race -coverpkg=./... -covermode=atomic -coverprofile=coverage.out $(PACKAGES)