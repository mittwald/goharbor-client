.PHONY: swagger swaggerclientcleanup swaggermodelcleanup harbor-1.10.2 integration-test-v1.10.0 teardown

swagger:
	scripts/swagger-gen.sh

swaggercleanup: swaggerclientcleanup swaggermodelcleanup

swaggerclientcleanup:
	rm -rf ./client

swaggermodelcleanup:
	rm -rf ./model

# Creates a Harbor instance as a docker container via Kind.
# Delete cluster via scripts/teardown-harbor.sh
harbor-1.10.2:
	scripts/setup-harbor.sh 1.10.2

harbor-1.10.1:
	scripts/setup-harbor.sh 1.10.1

harbor-1.10.0:
	scripts/setup-harbor.sh 1.10.0

# Testing on Harbor 1.10.2
integration-test-v1.10.2:
	CGO_ENABLED=0 go test -v github.com/mittwald/goharbor-client -integration -version=1.10.2

# Testing on Harbor 1.10.1
integration-test-v1.10.1:
	CGO_ENABLED=0 go test -v github.com/mittwald/goharbor-client -integration -version=1.10.1

# Testing on Harbor 1.10.0
integration-test-v1.10.0:
	CGO_ENABLED=0 go test -v github.com/mittwald/goharbor-client -integration -version=1.10.0

teardown:
	scripts/teardown-harbor.sh
