include docker/.env

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

## soko-bora: run the cmd/mallbots application
.PHONY: soko-bora
soko-bora:
	docker compose --profile monolith up


# Stop everything and remove volumes
.PHONY: clean-up-monolith
clean-up-soko-bora:
	docker-compose down -v

.PHONY: run/install-tools
install-tools:
	@echo installing tools && \
	go install \
	github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
	github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
	google.golang.org/protobuf/cmd/protoc-gen-go \
	google.golang.org/grpc/cmd/protoc-gen-go-grpc
	@echo done

generate:
	@echo running code generation
	@go generate ./...
	@echo done

