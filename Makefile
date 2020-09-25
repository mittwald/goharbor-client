.PHONY: swagger-v1 swagger-v2 mock-v1 mock-v2 harbor-1.10.4 harbor-2.0.2 teardown-harbor test swagger-cleanup \
integration-test-v1.10.3 integration-test-v1.10.4 integration-test-v2.0.2 gofumpt

swagger-v1:
	scripts/swagger-gen.sh v1.10.4

swagger-v2:
	scripts/swagger-gen.sh v2.0.2

mock-v1:
	scripts/gen-mock.sh v1

mock-v2:
	scripts/gen-mock.sh v2

# Creates a Harbor instance as a docker container via Kind.
# Delete cluster via scripts/teardown-harbor.sh
harbor-1.10.4:
	scripts/setup-harbor.sh 1.10.4

harbor-2.0.2:
	scripts/setup-harbor.sh 2.0.2

teardown-harbor:
	scripts/teardown-harbor.sh

test:
	go test -v ./...

swagger-cleanup:
	rm -rf ./apiv*/internal ./apiv*/model/

# Testing on Harbor 1.10.3
integration-test-v1.10.3:
	CGO_ENABLED=0 go test -p 1 -count 1 -v github.com/mittwald/goharbor-client/apiv1/... -version=1.10.3 -tags integration

# Testing on Harbor 1.10.4
integration-test-v1.10.4:
	CGO_ENABLED=0 go test -p 1 -count 1 -v github.com/mittwald/goharbor-client/apiv1/... -version=1.10.4 -tags integration

# Testing on Harbor 2.0.2
integration-test-v2.0.2:
	CGO_ENABLED=0 go test -p 1 -count 1 -v github.com/mittwald/goharbor-client/apiv2/... -version=2.0.2 -tags integration

gofmt:
	find . \( -path "./apiv*/internal" -o -path "./apiv*/mocks" -o -path "./apiv*/model" \)  \
	-prune -false -o -name '*.go' -exec gofmt -l -w {} \;