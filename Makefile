default: generate dependency

# Removes currently generated client files
cleanup:
	@echo "Removing current client files"
	rm -rf client/*


# Validating Swagger spec
#
# https://openapi-generator.tech/docs/usage/#validate
validate:
	@echo "Validating Swagger spec"
	openapi-generator validate -i swagger.json

# Generating API client
# 
# https://openapi-generator.tech/docs/generators/go
generate: cleanup validate
	@echo "Generating API client"
	openapi-generator generate \
		-i swagger.json \
		-g go \
		-o client \
		--package-name client \
		--global-property skipFormModel=false \
		-p enumClassPrefix=true
	@echo "Removing generated dependencies file"
	rm -rf client/go.*

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