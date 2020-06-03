#!/usr/bin/env bash

CLUSTER_NAME="goharbor-client-integration-tests"

kind version &>/dev/null
if [[ $? -ne "0" ]]; then
    &>2 echo "kind not installed, aborting."
    exit 1
fi

echo "Delete existing kind cluster..."
kind delete cluster --name "${CLUSTER_NAME}"
if [[ "$?" -ne "0" ]]; then
    &>2 echo "Could not delete kind cluster, aborting."
    exit 1
fi