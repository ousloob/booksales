SHELL := /bin/bash

# ==============================================================================
# Running the API locally

.PHONY: run
run:
	go run api/sales/main.go

# ==============================================================================
# Running tests locally

.PHONY: lint
lint:
	CGO_ENABLED=0 go vet ./...
	staticcheck -checks=all ./...
	
.PHONY: test
test:
	go test -count=1 ./...

.PHONY: vulncheck
vulncheck:
	govulncheck ./...

# ==============================================================================	
# Module support

.PHONY: tidy
tidy:
	go mod tidy
	go mod vendor

# ==============================================================================