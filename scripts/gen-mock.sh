#!/bin/bash
BUILD_FLAG="// +build !integration"

if ! mockery --version; then
    >&2 echo "Mockery not installed, aborting."
    exit 1
fi

if [[ "${1}" = *"v1"* ]]; then
  MOCK_FILE=v1/mocks/client_service.go
  if ! mockery --name ClientService --dir apiv1/internal/api/client/products/ \
  --output ./apiv1/mocks --filename client_service.go --structname MockClientService; then
    >&2 echo "Mockery command failed."
    exit 1
  fi
  printf "%s\n" 1 i "${BUILD_FLAG}" . w | ed -s "${MOCK_FILE}" &>/dev/null
fi

if [[ "${1}" = *"v2"* ]]; then
  MOCK_FILE=apiv2/mocks/client_service.go
  if ! mockery --name ClientService --dir ./apiv2/internal/api/"${1}"/client/products/ \
  --output ./apiv2/mocks/ --filename client_service.go --structname MockClientService; then
    >&2 echo "Mockery command failed."
    exit 1
  fi
    printf "%s\n" 1 i "${BUILD_FLAG}" . w | ed -s "${MOCK_FILE}" &>/dev/null
fi
