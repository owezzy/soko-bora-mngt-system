# Include variables from the .envrc file
include .envrc

# ==================================================================================== #
# HELPERS
# ==================================================================================== #

## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /'
# Create the new confirm target.
.PHONY: confirm
confirm:
	@echo -n 'Are you sure? [y/N] ' && read ans && [ $${ans:-N} = y ]


# ==================================================================================== # 
# DEVELOPMENT
# ==================================================================================== #
## run/api: run the cmd/api application
.PHONY: run/api
run/api:
	JWT_SECRET=${JWT_SECRET} MONGO_URI=${MONGO_URI} MONGO_DATABASE=${MONGO_DATABASE} go run ./cmd/api/*.go

generate-open-api-doc:
	swagger generate spec -o ./swagger.json

serve-open-api-doc:
	swagger serve --flavor=swagger ./swagger.json


.PHONY: run/install-tools
run/install-tools:
	@echo installing tools && \
	@go install \
	github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
	github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
	google.golang.org/protobuf/cmd/protoc-gen-go \
	google.golang.org/grpc/cmd/protoc-gen-go-grpc
	@echo done

generate:
	@echo running code generation
	@go generate ./...
	@echo done
