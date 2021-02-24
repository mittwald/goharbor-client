.PHONY: generate swagger-generate mock harbor-v1 harbor-v2 harbor-teardown test swagger-cleanup \
mock-cleanup integration-test-v1 integration-test-v2

V1_VERSION = v1.10.6
V2_VERSION = v2.1.3
MOCKERY_VERSION = v2.5.1
GOSWAGGER_VERSION = v0.26.1

# Run all code generation targets
generate: swagger-generate mock-generate

# Run go-swagger code generation
swagger-generate: swagger-cleanup
	scripts/swagger-gen.sh $(V1_VERSION) $(GOSWAGGER_VERSION)
	scripts/swagger-gen.sh $(V2_VERSION) $(GOSWAGGER_VERSION)

# Run mockery
mock-generate: mock-cleanup
	scripts/gen-mock.sh v1 $(MOCKERY_VERSION)
	scripts/gen-mock.sh v2 $(MOCKERY_VERSION)

# Create a Harbor instance as a docker container via Kind.
# Delete cluster via scripts/teardown-harbor.sh
harbor-v1:
	scripts/setup-harbor.sh $(V1_VERSION)

harbor-v2:
	scripts/setup-harbor.sh $(V2_VERSION)

harbor-teardown:
	scripts/teardown-harbor.sh

test:
	go test -v ./...

swagger-cleanup:
	rm -rf ./apiv*/internal ./apiv*/model/

# Delete all auto-generated mock files
mock-cleanup:
	rm -rf ./apiv*/mocks/*

# Integration testing on Harbor v1
integration-test-v1: harbor-v1
	CGO_ENABLED=0 go test -p 1 -count 1 -v github.com/mittwald/goharbor-client/v3/apiv1/... -tags integration

# Integration testing on Harbor v2
integration-test-v2: harbor-v2
	CGO_ENABLED=0 go test -p 1 -count 1 -v github.com/mittwald/goharbor-client/v3/apiv2/... -tags integration

# Exclude auto-generated code to be formatted by gofmt
gofmt:
	find . \( -path "./apiv*/internal" -o -path "./apiv*/mocks" -o -path "./apiv*/model" \)  \
	-prune -false -o -name '*.go' -exec gofmt -l -w {} \;
