.PHONY: swagger-v1.10.4 swagger-v2.0.2 swagger-cleanup

swagger-v1.10.4:
	scripts/swagger-gen.sh v1.10.4

swagger-v2.0.2:
	scripts/swagger-gen.sh v2.0.2

mock-v1.10.4:
	scripts/gen-mock.sh v1_10_4

mock-v2.0.2:
	scripts/gen-mock.sh v2_0_2

# Creates a Harbor instance as a docker container via Kind.
# Delete cluster via scripts/teardown-harbor.sh
harbor-1.10.4:
	scripts/setup-harbor.sh 1.10.4

teardown-harbor:
	scripts/teardown-harbor.sh

test:
	go test -v ./...

swagger-cleanup:
	rm -rf ./v*/internal
	rm -rf ./v*/model/

# Testing on Harbor 1.10.3
integration-test-v1.10.3:
	CGO_ENABLED=0 go test -p 1 -count 1 -v ./... github.com/mittwald/goharbor-client -version=1.10.3 -tags integration

# Testing on Harbor 1.10.4
integration-test-v1.10.4:
	CGO_ENABLED=0 go test -p 1 -count 1 -v ./... github.com/mittwald/goharbor-client -version=1.10.4 -tags integration
