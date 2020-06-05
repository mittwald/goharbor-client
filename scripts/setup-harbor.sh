#!/usr/bin/env bash

CLUSTER_NAME="goharbor-client-integration-tests"
HARBOR_VERSION="${1}"
CHART_VERSION=""


echo "Check for existence of neccessary tools..."

docker --version &>/dev/null
if [[ $? -ne "0" ]]; then
    >&2 echo "Docker not installed, aborting."
    exit 1
fi

kind version &>/dev/null
if [[ $? -ne "0" ]]; then
    >&2 echo "kind not installed, aborting."
    exit 1
fi

helm_version="$(helm version --short)"
if ! [[ ${helm_version} =~ ^v3. ]]; then
    >&2 echo "Helm not installed or not v3, aborting."
    exit 1
fi

jq --version &>/dev/null
if [[ $? -ne "0" ]]; then
    >&2 echo "jq not installed, aborting."
    exit 1
fi

echo "Check needed program arguments..."
found="false"
if [[ -z "${HARBOR_VERSION}" ]]; then
    >&2 echo "Harbor version as first argument not provided, aborting."
    exit 1
fi

# map Harbor version to it's helm chart version
while read CHART HARBOR; do
    if [[ "${HARBOR_VERSION}" == "${HARBOR}" ]]; then
        CHART_VERSION="${CHART}"
    fi
done <<< "1.4.0 2.0.0
1.3.2 1.10.2
1.3.1 1.10.1
1.3.0 1.10.0
1.2.4 1.9.4
1.2.3 1.9.3
1.2.2 1.9.2
1.2.1 1.9.1
1.2.0 1.9.0
1.1.6 1.8.6
1.1.5 1.8.5
1.1.4 1.8.4
1.1.3 1.8.3
1.1.2 1.8.2
1.1.1 1.8.1
1.1.0 1.8.0
1.0.1 1.7.5
1.0.0 1.7.0"
if [[ -z "${CHART_VERSION}" ]]; then
    >&2 echo "Unsupported Harbor version, aborting."
    exit 1
fi

echo "Create new kind cluster to deploy Harbor into..."
kind create cluster --config scripts/kind-config.yaml --name "${CLUSTER_NAME}"
if [[ "$?" -ne "0" ]]; then
    >&2 echo "Could not create kind cluster, aborting."
    exit 1
fi

echo "Install Harbor via helm..."
helm repo add harbor https://helm.goharbor.io
helm install harbor harbor/harbor \
    --set expose.type=nodePort,expose.tls.enabled=false,externalURL=http://localhost \
    --namespace default \
    --kube-context kind-"${CLUSTER_NAME}" \
    --version="${CHART_VERSION}"
if [[ "$?" -ne "0" ]]; then
    >&2 echo "Could not install Harbor, cleaning up and aborting."
    kind delete cluster --name "${CLUSTER_NAME}"
    exit 1
fi

echo "Waiting for Harbor to become ready..."
API_URL_PREFIX="http://localhost:30002/api"
if [[ "${HARBOR_VERSION}" =~ ^2 ]]; then
    API_URL_PREFIX="http://localhost:30002/api/v2.0"
fi

for i in {1..100}; do
    echo "Ping Harbor instance ($i/100)..."
    status="$(curl -s -X GET --connect-timeout 3 ${API_URL_PREFIX}/health | jq '.status' 2>/dev/null)"
    if [[ "${status}" == "\"healthy\"" ]]; then
        echo "Harbor installation finished sucessfully. Visit at http://localhost:30002"
        exit 0
    fi
    sleep 5
done

echo "Timeout waiting for Harbor installation, aborting."
exit 1