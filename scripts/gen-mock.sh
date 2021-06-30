#!/bin/bash
MOCKERY_IMAGE="vektra/mockery:${2}"

# v1 API
if [[ "${1}" = *"v1"* ]]; then
  echo "generating mocks using the v1 API"
  docker run --rm -e GOPATH="${HOME}/go:/go" -v "${HOME}:${HOME}" -w "$(pwd)" "${MOCKERY_IMAGE}" \
  --name ClientService --dir apiv1/internal/api/client/products/ --quiet \
  --output ./apiv1/mocks --filename client_service.go --structname MockClientService
fi

if [[ "${1}" = *"v2"* ]]; then
# v2 API
  while IFS= read -r -d '' DIR
    do
      CLIENT=$(basename "${DIR}")
      echo "generating mocks for the '${CLIENT}' client using the v2 API"
      docker run --rm -e GOPATH="${HOME}/go:/go" -v "${HOME}:${HOME}" -w "$(pwd)" "${MOCKERY_IMAGE}" \
      --name ClientService --dir apiv2/internal/api/client/"${CLIENT}" -r --quiet \
      --output ./apiv2/mocks --filename "${CLIENT}"_client_service.go --structname Mock"${CLIENT^}"ClientService
    done < <(find ./apiv2/internal/api/client/* -type d -print0)
# v2 legacy API
  while IFS= read -r -d '' DIR
    do
      CLIENT=$(basename "${DIR}")
      echo "generating mocks for the '${CLIENT}' client using the v2 legacy API"
      docker run --rm -e GOPATH="${HOME}/go:/go" -v "${HOME}:${HOME}" -w "$(pwd)" "${MOCKERY_IMAGE}" \
      --name ClientService --dir apiv2/internal/legacyapi/client/"${CLIENT}" -r --quiet \
      --output ./apiv2/mocks --filename "${CLIENT}"_client_service.go --structname Mock"${CLIENT^}"ClientService
    done < <(find ./apiv2/internal/legacyapi/client/* -type d -print0)
fi