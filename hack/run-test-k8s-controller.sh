#!/usr/bin/env bash
#
# This script provisions testing environment using 'kind'(kubernetes-in-docker)
# and execute end-to-end Voltron tests.
#
# It requires Docker to be installed.
# Dependencies such as kubebuilder can be installed on demand.

# standard bash error handling
set -o nounset # treat unset variables as an error and exit immediately.
set -o errexit # exit immediately when a command fails.
set -E         # needs to be set if we want the ERR trap

readonly CURRENT_DIR=$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )
readonly REPO_ROOT_DIR=${CURRENT_DIR}/..
readonly TMP_DIR=$(mktemp -d)

# shellcheck source=./hack/lib/utilities.sh
source "${CURRENT_DIR}/lib/utilities.sh" || { echo 'Cannot load CI utilities.'; exit 1; }
# shellcheck source=./hack/lib/deps_ver.sh
source "${CURRENT_DIR}/lib/deps_ver.sh" || { echo 'Cannot load dependencies versions.'; exit 1; }

SKIP_DEPS_INSTALLATION=${SKIP_DEPS_INSTALLATION:-true}

cleanup() {
    rm -rf "${TMP_DIR}"
}

trap cleanup EXIT

main() {
    if [[ "${SKIP_DEPS_INSTALLATION}" == "false" ]]; then
        export INSTALL_DIR=${TMP_DIR} KUBEBUILDER_VERSION=${STABLE_KUBEBUILDER_VERSION}
        host::install::kubebuilder
    else
        echo "Skipping kubebuilder installation cause SKIP_DEPS_INSTALLATION is set to true."
    fi

    shout "Starting k8s controller tests..."

    go test -v --tags=controllertests "${REPO_ROOT_DIR}/internal/k8s-engine/controller/..."

    shout "K8s controller tests completed successfully."
}

main