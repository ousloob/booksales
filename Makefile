SHELL := /bin/sh

# ==============================================================================	
# Module support

.PHONY: tidy
tidy:
	go mod tidy
	go mod vendor

# ==============================================================================
# Running the API locally

.PHONY: run
run:
	go run api/sales/main.go

# ==============================================================================
# Code quality check locally

.PHONY: test lint vulncheck check

# Run all checks
check: test lint vulncheck

test:
	go test -count=1 ./...

lint:
	CGO_ENABLED=0 go vet ./...
	staticcheck -checks=all ./...

vulncheck:
	govulncheck ./...

# ==============================================================================