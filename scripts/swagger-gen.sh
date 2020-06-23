#!/usr/bin/env bash
SWAGGER_IMAGE="quay.io/goswagger/swagger"
API_VERSION_V1="v1.10.0"
API_VERSION_V2="v2.0.0"
SWAGGER_FILE_V1="https://raw.githubusercontent.com/goharbor/harbor/${API_VERSION_V1}/api/harbor/swagger.yaml"
SWAGGER_FILE_V2=https://raw.githubusercontent.com/goharbor/harbor/${API_VERSION_V2}/api/${API_VERSION_V2%.0}/swagger.yaml
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
swagger_operations+=("PostProjectsProjectIDMembers")
swagger_operations+=("GetProjectsProjectIDMembers")
swagger_operations+=("PutProjectsProjectIDMembersMid")
swagger_operations+=("DeleteProjectsProjectIDMembersMid")
swagger_operations+=("GetProjectsProjectIDMetadatasMetaName")
swagger_operations+=("GetProjectsProjectIDMetadatas")
swagger_operations+=("PutProjectsProjectIDMetadatasMetaName")
swagger_operations+=("PostProjectsProjectIDMetadatas")
swagger_operations+=("DeleteProjectsProjectIDMetadatasMetaName")

swagger_operations+=("GetRegistries")
swagger_operations+=("PostRegistries")
swagger_operations+=("PutRegistriesID")
swagger_operations+=("DeleteRegistriesID")

swagger_operations+=("GetUsers")
swagger_operations+=("PostUsers")
swagger_operations+=("PutUsersUserID")
swagger_operations+=("DeleteUsersUserID")

swagger_operations+=("PostReplicationPolicies")
swagger_operations+=("GetReplicationPolicies")
swagger_operations+=("PutReplicationPoliciesID")
swagger_operations+=("GetReplicationPoliciesID")
swagger_operations+=("DeleteReplicationPoliciesID")

swagger_operations+=("PostSystemGcSchedule")
swagger_operations+=("GetSystemGcSchedule")
swagger_operations+=("PutSystemGcSchedule")

for i in "${swagger_operations[@]}"; do
  operation_flags+="--operation=${i} "
done

if [[ "$1" = "v1" ]]; then
  echo "using the v1 swagger file (${API_VERSION_V1})"
  docker run --rm -e GOPATH="${HOME}/go:/go" -v "${HOME}:${HOME}" -w "$(pwd)" ${SWAGGER_IMAGE} \
  generate client \
  --skip-validation \
  --model-package="api/${API_VERSION_V1}/model" \
  --name="harbor" \
  --client-package="api/${API_VERSION_V1}/client" \
  --spec="${SWAGGER_FILE_V1}" \
  ${operation_flags}
fi

if [[ "$1" = "v2" ]]; then
  set -x
  echo "using the v2 swagger file (${API_VERSION_V2})"
    docker run --rm -e GOPATH="${HOME}/go:/go" -v "${HOME}:${HOME}" -w "$(pwd)" ${SWAGGER_IMAGE} \
  generate client \
  --skip-validation \
  --model-package="api/${API_VERSION_V2%.0}/model" \
  --name="harbor" \
  --client-package="api/${API_VERSION_V2%.0}/client" \
  --spec="${SWAGGER_FILE_V2}"
fi
