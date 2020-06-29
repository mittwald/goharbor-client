.PHONY: swagger-v1 swagger-v2 swaggerclientcleanup swaggermodelcleanup harbor-1.10.2 teardown-harbor integration-test-v1.10.1 integration-test-v1.10.0 mock test

swagger-v1.10.0:
	scripts/swagger-gen.sh v1.10.0
swagger-v2.0.0:
	scripts/swagger-gen.sh v2.0.0

swaggercleanup: swaggerclientcleanup swaggermodelcleanup

swaggerclientcleanup:
	rm -rf ./internal/api/*/client

swaggermodelcleanup:
	rm -rf ./internal/api/*/model

# Creates a Harbor instance as a docker container via Kind.
# Delete cluster via scripts/teardown-harbor.sh
harbor-1.10.2:
	scripts/setup-harbor.sh 1.10.2

harbor-1.10.1:
	scripts/setup-harbor.sh 1.10.1

harbor-1.10.0:
	scripts/setup-harbor.sh 1.10.0

teardown-harbor:
	scripts/teardown-harbor.sh

test:
	go test -v ./...

mock:
	mockery -name ClientService -dir internal/api/v1.10.0/client/products/. -filename client_service.go -structname MockClientService

# Testing on Harbor 1.10.2
integration-test-v1.10.2:
	CGO_ENABLED=0 go test -p 1 -count 1 -v ./... github.com/mittwald/goharbor-client -version=1.10.2 -tags integration

# Testing on Harbor 1.10.1
integration-test-v1.10.1:
	CGO_ENABLED=0 go test -p 1 -count 1 -v ./... github.com/mittwald/goharbor-client -version=1.10.1 -tags integration

# Testing on Harbor 1.10.0
integration-test-v1.10.0:
	CGO_ENABLED=0 go test -p 1 -count 1 -v ./... github.com/mittwald/goharbor-client -version=1.10.0 -tags integration
