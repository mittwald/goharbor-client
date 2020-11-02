.PHONY: swagger-generate mock harbor-v1 harbor-v2 teardown-harbor test swagger-cleanup \
mock-cleanup integration-test-v1 integration-test-v2

V1_VERSION = v1.10.5
V2_VERSION = v2.1.1

swagger-generate: swagger-cleanup
	scripts/swagger-gen.sh $(V1_VERSION)
	scripts/swagger-gen.sh $(V2_VERSION)

mock: mock-cleanup
	scripts/gen-mock.sh v1
	scripts/gen-mock.sh v2

# Creates a Harbor instance as a docker container via Kind.
# Delete cluster via scripts/teardown-harbor.sh
harbor-v1:
	scripts/setup-harbor.sh $(V1_VERSION)

harbor-v2:
	scripts/setup-harbor.sh $(V2_VERSION)

teardown-harbor:
	scripts/teardown-harbor.sh

test:
	go test -v ./...

swagger-cleanup:
	rm -rf ./apiv*/internal ./apiv*/model/

mock-cleanup:
	rm -rf ./apiv*/mocks/*

# Testing on Harbor v1
integration-test-v1:
	CGO_ENABLED=0 go test -p 1 -count 1 -v github.com/mittwald/goharbor-client/v3/apiv1/... -version=$(V1_VERSION) -tags integration

# Testing on Harbor v2
integration-test-v2:
	CGO_ENABLED=0 go test -p 1 -count 1 -v github.com/mittwald/goharbor-client/v3/apiv2/... -version=$(V2_VERSION) -tags integration

gofmt:
	find . \( -path "./apiv*/internal" -o -path "./apiv*/mocks" -o -path "./apiv*/model" \)  \
	-prune -false -o -name '*.go' -exec gofmt -l -w {} \;
