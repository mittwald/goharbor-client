#!/usr/bin/env bash
GOSWAGGER_IMAGE="quay.io/goswagger/swagger:${2}"
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
swagger_operations+=("PutUsersUserIDPassword")
swagger_operations+=("DeleteUsersUserID")

swagger_operations+=("PostReplicationPolicies")
swagger_operations+=("GetReplicationPolicies")
swagger_operations+=("PutReplicationPoliciesID")
swagger_operations+=("GetReplicationPoliciesID")
swagger_operations+=("DeleteReplicationPoliciesID")
swagger_operations+=("PostReplicationExecutions")
swagger_operations+=("GetReplicationExecutions")
swagger_operations+=("GetReplicationExecutionsID")

swagger_operations+=("PostSystemGcSchedule")
swagger_operations+=("GetSystemGcSchedule")
swagger_operations+=("PutSystemGcSchedule")

swagger_operations+=("GetRetentionsMetadatas")
swagger_operations+=("PostRetentions")
swagger_operations+=("GetRetentionsID")


for i in "${swagger_operations[@]}"; do
  operation_flags+="--operation=${i} "
done

if [[ "${1}" = *"v1"* ]]; then
  SWAGGER_FILE="https://raw.githubusercontent.com/goharbor/harbor/${1}/api/harbor/swagger.yaml"
  echo "generating client API using the v1 swagger file (${1})"
  docker run --rm -e GOPATH="${HOME}/go:/go" -v "${HOME}:${HOME}" -w "$(pwd)" "${GOSWAGGER_IMAGE}" \
  generate client \
  -q \
  --skip-validation \
  --model-package"=apiv1/model/"\
  --name="harbor" \
  --client-package="apiv1/internal/api/client" \
  --spec="${SWAGGER_FILE}" \
  ${operation_flags}
fi

if [[ "${1}" = *"v2"* ]]; then
  LEGACY_SWAGGER_FILE="https://raw.githubusercontent.com/goharbor/harbor/${1}/api/v2.0/legacy_swagger.yaml"
  SWAGGER_FILE="https://raw.githubusercontent.com/goharbor/harbor/${1}/api/v2.0/swagger.yaml"
  echo "generating client API using the v2 swagger files (${1})"
  # Generate client using the Harbor v2 legacy API
  docker run --rm -e GOPATH="${HOME}/go:/go" -v "${HOME}:${HOME}" -w "$(pwd)" "${GOSWAGGER_IMAGE}" \
  generate client \
  -q \
  --skip-validation \
  --model-package="apiv2/model/legacy" \
  --name="harbor" \
  --client-package="apiv2/internal/legacyapi/client" \
  --spec="${LEGACY_SWAGGER_FILE}"
  # Generate client using the new / changed Harbor v2 API
  docker run --rm -e GOPATH="${HOME}/go:/go" -v "${HOME}:${HOME}" -w "$(pwd)" "${GOSWAGGER_IMAGE}" \
  generate client \
  -q \
  --skip-validation \
  --model-package="apiv2/model/" \
  --name="harbor" \
  --client-package="apiv2/internal/api/client" \
  --spec="${SWAGGER_FILE}"
fi
