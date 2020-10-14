.PHONY: swagger-v1 swagger-v2 mock-v1 mock-v2 harbor-1.10.4 harbor-2.0.2 teardown-harbor test swagger-cleanup \
integration-test-v1.10.3 integration-test-v1.10.4 integration-test-v2.0.2 gofmt

swagger-v1:
	scripts/swagger-gen.sh v1.10.5

swagger-v2:
	scripts/swagger-gen.sh v2.1.0

mock-v1:
	scripts/gen-mock.sh v1

mock-v2:
	scripts/gen-mock.sh v2

# Creates a Harbor instance as a docker container via Kind.
# Delete cluster via scripts/teardown-harbor.sh
harbor-1.10.5:
	scripts/setup-harbor.sh 1.10.5

harbor-2.1.0:
	scripts/setup-harbor.sh 2.1.0

teardown-harbor:
	scripts/teardown-harbor.sh

test:
	go test -v ./...

swagger-cleanup:
	rm -rf ./apiv*/internal ./apiv*/model/

# Testing on Harbor 1.10.5
integration-test-v1.10.5:
	CGO_ENABLED=0 go test -p 1 -count 1 -v github.com/mittwald/goharbor-client/v2/apiv1/... -version=1.10.5 -tags integration

# Testing on Harbor 2.1.0
integration-test-v2.1.0:
	CGO_ENABLED=0 go test -p 1 -count 1 -v github.com/mittwald/goharbor-client/v2/apiv2/... -version=2.1.0 -tags integration

gofmt:
	find . \( -path "./apiv*/internal" -o -path "./apiv*/mocks" -o -path "./apiv*/model" \)  \
	-prune -false -o -name '*.go' -exec gofmt -l -w {} \;
