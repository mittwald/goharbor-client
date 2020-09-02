.PHONY: swagger-v1 swagger-v2 swaggerclientcleanup swaggermodelcleanup harbor-1.10.2 \
	teardown-harbor integration-test-v1.10.1 integration-test-v1.10.0 mock test

swagger-v1.10.4:
	scripts/swagger-gen.sh v1.10.4
swagger-v2.0.2:
	scripts/swagger-gen.sh v2.0.2

swaggercleanup: swaggerclientcleanup swaggermodelcleanup

swaggerclientcleanup:
	rm -rf ./internal/api/*

swaggermodelcleanup:
	rm -rf ./model/*/

# Creates a Harbor instance as a docker container via Kind.
# Delete cluster via scripts/teardown-harbor.sh
harbor-1.10.0:
	scripts/setup-harbor.sh 1.10.0

harbor-1.10.1:
	scripts/setup-harbor.sh 1.10.1

harbor-1.10.2:
	scripts/setup-harbor.sh 1.10.2

harbor-1.10.3:
	scripts/setup-harbor.sh 1.10.3

harbor-1.10.4:
	scripts/setup-harbor.sh 1.10.4

teardown-harbor:
	scripts/teardown-harbor.sh

test:
	go test -v ./...

mock-v1.10.4:
	scripts/gen-mock.sh v1_10_4

mock-v2.0.2:
	scripts/gen-mock.sh v2_0_2

# Testing on Harbor 1.10.0
integration-test-v1.10.0:
	CGO_ENABLED=0 go test -p 1 -count 1 -v ./... github.com/mittwald/goharbor-client -version=1.10.0 -tags integration

# Testing on Harbor 1.10.1
integration-test-v1.10.1:
	CGO_ENABLED=0 go test -p 1 -count 1 -v ./... github.com/mittwald/goharbor-client -version=1.10.1 -tags integration

# Testing on Harbor 1.10.2
integration-test-v1.10.2:
	CGO_ENABLED=0 go test -p 1 -count 1 -v ./... github.com/mittwald/goharbor-client -version=1.10.2 -tags integration

# Testing on Harbor 1.10.3
integration-test-v1.10.3:
	CGO_ENABLED=0 go test -p 1 -count 1 -v ./... github.com/mittwald/goharbor-client -version=1.10.2 -tags integration

# Testing on Harbor 1.10.4
integration-test-v1.10.4:
	CGO_ENABLED=0 go test -p 1 -count 1 -v ./... github.com/mittwald/goharbor-client -version=1.10.2 -tags integration
