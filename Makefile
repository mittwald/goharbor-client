.PHONY: swagger swaggerclientcleanup swaggermodelcleanup harbor-%

swagger:
	scripts/swagger-gen.sh

swaggercleanup: swaggerclientcleanup swaggermodelcleanup

swaggerclientcleanup:
	rm -rf ./client

swaggermodelcleanup:
	rm -rf ./model

integration-test:
	go test -v ./... -integration -version=1.10.2

# Creates a Harbor instance as a docker container via Kind.
# Specify a version after "-", i.e.
# make harbor-2.0.0
# make harbor-1.10.2
# Delete cluster via scripts/teardown-harbor.sh
harbor-%:
	scripts/setup-harbor.sh $$(echo $@ | sed 's/harbor-//')

