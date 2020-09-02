#!/bin/bash
BUILD_FLAG="// +build !integration"
MOCK_FILE=mocks/client_service.go

if ! mockery --version; then
    >&2 echo "Mockery not installed, aborting."
    exit 1
fi

if ! mockery --name ClientService --dir ./internal/api/"${1}"/client/products/ \
  --filename client_service.go --structname MockClientService --log-level debug; then
  >&2 echo "Mockery command failed."
  exit 1
fi

printf "%s\n" 1 i "${BUILD_FLAG}" . w | ed -s "${MOCK_FILE}" &>/dev/null
