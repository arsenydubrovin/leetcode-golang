.PHONY: help run deps fmt

all: help

# Run
run:
	go run .

# Update dependencies
deps:
	go mod tidy -v

# Format all files
fmt:
	go fmt .

# Show this help
help:
	@awk '/^#/{c=substr($$0,3);next}c&&/^[[:alpha:]][[:alnum:]_-]+:/{print substr($$1,1,index($$1,":")),c}1{c=0}' $(MAKEFILE_LIST) | column -s: -t