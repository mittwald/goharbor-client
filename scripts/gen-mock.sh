#!/bin/bash
BUILD_FLAG="// +build !integration"
MOCKERY_IMAGE="vektra/mockery:v2.2.1"

if [[ "${1}" = *"v1"* ]]; then
  MOCK_FILE=apiv1/mocks/client_service.go
  echo "generating mocks using the v1 API"
  docker run --rm -e GOPATH="${HOME}/go:/go" -v "${HOME}:${HOME}" -w "$(pwd)" ${MOCKERY_IMAGE} \
  --name ClientService --dir apiv1/internal/api/client/products/ \
  --output ./apiv1/mocks --filename client_service.go --structname MockClientService
  printf "%s\n" 1 i "${BUILD_FLAG}" . w | ed -s "${MOCK_FILE}" &>/dev/null
fi

if [[ "${1}" = *"v2"* ]]; then
  MOCK_FILE=apiv2/mocks/client_service.go
  echo "generating mocks using the v2 API"
  docker run --rm -e GOPATH="${HOME}/go:/go" -v "${HOME}:${HOME}" -w "$(pwd)" ${MOCKERY_IMAGE} \
  --name ClientService --dir apiv2/internal/api/client/products/ \
  --output ./apiv2/mocks --filename client_service.go --structname MockClientService
  printf "%s\n" 1 i "${BUILD_FLAG}" . w | ed -s "${MOCK_FILE}" &>/dev/null
fi
