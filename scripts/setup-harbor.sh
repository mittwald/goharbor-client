#!/usr/bin/env bash

CLUSTER_NAME="goharbor-client-integration-tests"
HARBOR_VERSION="${1}"
HARBOR_CHART_VERSION=""
REGISTRY_IMAGE_TAG="2.7.1"

echo "Check for existence of necessary tools..."

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
if [[ -z "${HARBOR_VERSION}" ]]; then
    >&2 echo "Harbor version as first argument not provided, aborting."
    exit 1
fi

# Map Goharbor versions to their corresponding helmchart version
while read CHART HARBOR; do
    if [[ "${HARBOR_VERSION}" == "${HARBOR}" ]]; then
        HARBOR_CHART_VERSION="${CHART}"
    fi
done <<< "1.5.0 2.1.0
1.4.3 2.0.3
1.4.2 2.0.2
1.4.1 2.0.1
1.4.0 2.0.0
1.3.5 1.10.5
1.3.4 1.10.4
1.3.2 1.10.3
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

if [[ -z "${HARBOR_CHART_VERSION}" ]]; then
    >&2 echo "Unsupported Harbor version, aborting."
    exit 1
fi

echo "Creating a new kind cluster to deploy Harbor into..."
kind create cluster --config testdata/kind-config.yaml --name "${CLUSTER_NAME}"
if [[ "$?" -ne "0" ]]; then
    >&2 echo "Could not create kind cluster, aborting."
    exit 1
fi

echo "Installing Harbor via Helm..."
helm repo add harbor https://helm.goharbor.io && helm repo update
helm install harbor harbor/harbor \
    --set expose.type=nodePort,expose.tls.enabled=false,externalURL=http://localhost \
    --namespace default \
    --kube-context kind-"${CLUSTER_NAME}" \
    --version="${HARBOR_CHART_VERSION}"
if [[ "$?" -ne "0" ]]; then
    >&2 echo "Could not install Harbor, aborting."
    exit 1
fi

echo "Installing seperate docker registry for integration tests..."
helm repo add stable https://charts.helm.sh/stable && helm repo update
helm install registry stable/docker-registry \
    --set service.port=5000,image.tag=${REGISTRY_IMAGE_TAG}
if [[ "$?" -ne "0" ]]; then
    >&2 echo "Could not install Registry, aborting."
    exit 1
fi

echo "Waiting for Harbor to become ready..."
API_URL_PREFIX="http://localhost:30002/api"
if [[ "${HARBOR_VERSION}" =~ ^2 ]]; then
    API_URL_PREFIX="http://localhost:30002/api/v2.0"
fi

for i in {1..100}; do
    echo "Pinging Harbor instance ($i/100)..."
    STATUS="$(curl -s -X GET --connect-timeout 3 "${API_URL_PREFIX}/health" | jq '.status' 2>/dev/null)"
    if [[ "${STATUS}" == "\"healthy\"" ]]; then
        echo "Harbor installation finished successfully. Visit at http://localhost:30002"
        exit 0
    fi
    sleep 5
done

echo -e "Timeout while waiting for the Harbor installation to finish."
exit 1