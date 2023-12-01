golang_spec = spec/openapi_spec.yaml
vbr_spec = spec/swagger.json

default: generate 
cleanup:
	@echo "Cleaning up..."
	rm -f ./pkg/client/types.go
	rm -f ./pkg/client/client.go

generate: cleanup
	@echo "Generating types..."
	oapi-codegen -generate types -o ./pkg/client/types.go -package client ${golang_spec}
	@echo "Generating client..."
	oapi-codegen -generate client -o ./pkg/client/client.go -package client ${golang_spec}

convert:
	@echo "Converting spec..."
	cd tools/oapifixer && go build
	./tools/oapifixer/oapifixer -input ${vbr_spec} -output ${golang_spec}