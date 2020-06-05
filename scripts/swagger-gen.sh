#!/usr/bin/env bash

SWAGGER_IMAGE="quay.io/goswagger/swagger"
API_VERSION="v1.10.0"
SWAGGER_FILE="https://raw.githubusercontent.com/goharbor/harbor/${API_VERSION}/api/harbor/swagger.yaml"

# go-swaggers documentation on swagger operations at the moment is really sparse,
# so here is a bit of explanation from code observations.
# go-swagger accepts operation names from command line "--operation" flag, to filter which operations to generate.
# It tries to match by operationId of a swagger path, if set.
# If this doesn't match, it tries a match on a generated name using the operation method + path.
# This dynamically constructed name is a golint-compatible method name, which is best explained
# by looking at some examples:
#
# GET  /projects                --> GetProjects
# PUT  /projects/{project_id}   --> PutProjectsProjectID
# POST /chartrepo/{repo}/charts --> PostChartrepoRepoCharts
#
# WARNING: When adding new operations, make sure they really get created, since go-swagger DOES NOT complain about
# not finding the operation.
# Also it is not possible at the moment to list dynamically generated names from a swagger-file with go-swagger.
#
# see also: https://swagger.io/docs/specification/paths-and-operations/
declare -a swagger_operations
swagger_operations+=("GetHealth")
swagger_operations+=("GetProjects")
swagger_operations+=("PostProjects")
swagger_operations+=("GetProjectsProjectID")
swagger_operations+=("PutProjectsProjectID")
swagger_operations+=("DeleteProjectsProjectID")

for i in "${swagger_operations[@]}"; do
  operation_flags+="--operation=${i} "
done

docker run --rm -it -e GOPATH="${HOME}/go:/go" -v "${HOME}:${HOME}" -w "$(pwd)" ${SWAGGER_IMAGE} \
  generate client \
  --skip-validation \
  --model-package="api/${API_VERSION}/model" \
  --name="harbor" \
  --client-package="api/${API_VERSION}/client" \
  --spec="${SWAGGER_FILE}" \
  ${operation_flags}
