#!/bin/bash
MOCKERY_IMAGE="vektra/mockery:${2}"

# v1 API
if [[ "${1}" = *"v1"* ]]; then
  MOCK_FILE=apiv1/mocks/client_service.go
  echo "generating mocks using the v1 API"
  docker run --rm -e GOPATH="${HOME}/go:/go" -v "${HOME}:${HOME}" -w "$(pwd)" "${MOCKERY_IMAGE}" \
  --name ClientService --dir apiv1/internal/api/client/products/ \
  --output ./apiv1/mocks --filename client_service.go --structname MockClientService
fi

if [[ "${1}" = *"v2"* ]]; then
  # v2 API
  for CLIENT in artifact auditlog icon preheat project repository scan; do
    MOCK_FILE=apiv2/mocks/${CLIENT}_client_service.go
    echo "generating mocks for the '${CLIENT}' client using the v2 API"
    docker run --rm -e GOPATH="${HOME}/go:/go" -v "${HOME}:${HOME}" -w "$(pwd)" "${MOCKERY_IMAGE}" \
    --name ClientService --dir apiv2/internal/api/client/${CLIENT} -r \
    --output ./apiv2/mocks --filename ${CLIENT}_client_service.go --structname Mock${CLIENT^}ClientService
  done
  # v2 legacy API
  for CLIENT in products scanners; do
    MOCK_FILE=apiv2/mocks/${CLIENT}_client_service.go
    echo "generating mocks for the '${CLIENT}' client using the v2 legacy API"
    docker run --rm -e GOPATH="${HOME}/go:/go" -v "${HOME}:${HOME}" -w "$(pwd)" "${MOCKERY_IMAGE}" \
    --name ClientService --dir apiv2/internal/legacyapi/client/${CLIENT} -r \
    --output ./apiv2/mocks --filename ${CLIENT}_client_service.go --structname Mock${CLIENT^}ClientService
  done
fi
