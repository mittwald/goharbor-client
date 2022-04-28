.PHONY: generate swagger-generate swagger-cleanup mock-generate mock-cleanup setup-harbor-v1 setup-harbor-v2 \
harbor-teardown test integration-test-v1-ci integration-test-v2-ci integration-test-v1 integration-test-v2 \
fmt gofmt gofumpt goimports lint uninstall-harbor-v2 uninstall-harbor-v1

V1_VERSION = v1.10.9
V2_VERSION = v2.4.0
MOCKERY_VERSION = v2.9.4
GOSWAGGER_VERSION = v0.25.0
GOLANGCI_LINT_VERSION = v1.42.1

# Run all code generation targets
generate: swagger-generate mock-generate

# Run go-swagger code generation
swagger-generate: swagger-cleanup
	scripts/swagger-gen.sh $(V1_VERSION) $(GOSWAGGER_VERSION)
	scripts/swagger-gen.sh $(V2_VERSION) $(GOSWAGGER_VERSION)

# Delete all auto-generated API files
swagger-cleanup:
	rm -rf ./apiv*/internal ./apiv*/model/

# Run mockery
mock-generate: mock-cleanup
	scripts/gen-mock.sh v1 $(MOCKERY_VERSION)
	scripts/gen-mock.sh v2 $(MOCKERY_VERSION)

# Delete all auto-generated mock files
mock-cleanup:
	rm -rf ./apiv*/mocks/*

# Create a Harbor instance as a docker container via Kind.
# Delete cluster via scripts/teardown-harbor.sh
setup-harbor-v1:
	scripts/setup-harbor.sh $(V1_VERSION)

setup-harbor-v2:
	scripts/setup-harbor.sh $(V2_VERSION)

uninstall-harbor-v2:
	kind delete clusters "goharbor-client-integration-tests-$(V2_VERSION)"

uninstall-harbor-v1:
	kind delete clusters "goharbor-client-integration-tests-$(V1_VERSION)"

test:
	go test -v ./... -tags !integration

INTREGRATION_V1 = CGO_ENABLED=0 go test -p 1 -count 1 -v github.com/mittwald/goharbor-client/v5/apiv1/... -tags integration
INTEGRATION_V2 = CGO_ENABLED=0 go test -p 1 -count 1 -v github.com/mittwald/goharbor-client/v5/apiv2/... -tags integration

# Integration testing (CI Jobs)
integration-test-v1-ci: setup-harbor-v1 integration-test-v1

integration-test-v2-ci: setup-harbor-v2 integration-test-v2

# Integration testing (local execution)
integration-test-v1:
	$(INTEGRATION_V1)

integration-test-v2: upload-test-image
	$(INTEGRATION_V2)

upload-test-image:
	@echo Building and uploading test image
	docker login -u admin -p Harbor12345 core.harbor.domain
	docker build -t core.harbor.domain/library/image:test ./testdata
	docker push core.harbor.domain/library/image:test

# Exclude auto-generated code to be formatted by gofmt, gofumpt & goimports.
FIND=find . \( -path "./apiv*/internal" -o -path "./apiv*/mocks" -o -path "./apiv*/model" \) -prune -false -o -name '*.go'

fmt: gofmt gofumpt goimports tidy

tidy:
	go mod tidy

gofmt:
	$(FIND) -exec gofmt -l -w {} \;

gofumpt:
	$(FIND) -exec gofumpt -w {} \;

goimports:
	$(FIND) -exec goimports -w {} \;

lint:
	docker run --rm -v $(shell pwd):/goharbor-client -w /goharbor-client/. \
	golangci/golangci-lint:v1.45.2 golangci-lint run --sort-results
