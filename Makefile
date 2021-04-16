all: generate dependency

# Removes currently generated client files
cleanup:
	@echo "Removing current client files"
	rm -rf models client

# Validating Swagger spec
#
# https://goswagger.io/usage/validate.html
validate:
	@echo "Validating Swagger spec"
	swagger validate ./swagger.yml

# Generating API client
# 
# https://goswagger.io/generate/server.html
# https://goswagger.io/generate/client.html
generate: cleanup validate
	@echo "Generating API client"
#	swagger generate server \
		--skip-validation \
		--exclude-main \
		--target=./ \
		--spec=./swagger.yml \
		--name=veeam_vbr
	swagger generate client \
		--skip-validation \
		--spec=./swagger.yml \
		--name=veeam_vbr

# Adding new dependencies
dependency:
	@echo "Adding new dependencies"
	go mod tidy
	go mod vendor

build:
	@echo "Building Go module"
	go build ./...

test:
	@echo "Testing Go module"
	go test ./...

.PHONY: all cleanup validate generate dependency build test