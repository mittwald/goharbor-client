#!/usr/bin/env bash

CLUSTER_NAME="goharbor-client-integration-tests-${1}"
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

yq --version &>/dev/null
if [[ $? -ne "0" ]]; then
    >&2 echo "yq not installed, aborting."
    exit 1
fi

echo "Check needed program arguments..."
if [[ -z "${HARBOR_VERSION}" ]]; then
    >&2 echo "Harbor version as first argument not provided, aborting."
    exit 1
fi

# Map Goharbor versions to their corresponding helm chart version
while read CHART HARBOR; do
    if [[ "${HARBOR_VERSION#v}" == "${HARBOR}" ]]; then
        HARBOR_CHART_VERSION="${CHART}"
    fi
done <<< $(curl -s https://helm.goharbor.io/index.yaml | yq e '.entries.harbor[] | .version + " " + .appVersion' -)

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
    --set trivy.enabled=false,notary.enabled=false,chartmuseum.enabled=false \
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
if [[ "${HARBOR_VERSION}" =~ ^v2 ]]; then
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

HEALTH_STATUS="$(curl -s -X GET --connect-timeout 3 "${API_URL_PREFIX}/health" | jq)"

echo -e "Timeout while waiting for the Harbor installation to finish. Health status:"
echo "${HEALTH_STATUS}"
exit 1