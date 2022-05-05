#!/bin/bash
# This script is used to ensure that the go.mod file is up to date.

set -euo pipefail

BASE_DIR="$PWD"
TEMP_DIR=$(mktemp -d)
function cleanup() {
  rm -rf "${TEMP_DIR}"
}
trap cleanup EXIT

cp -r . "${TEMP_DIR}/"
cd $TEMP_DIR

for i in $(find $PWD -name go.mod); do
  pushd $(dirname $i)
  gofmt -s -w .
  popd
done

if ! diff -r . "${BASE_DIR}"; then
  echo
  echo "The code is not properly formatted, run gofmt."
  echo "Format them with the 'go fmt .' command."
  exit 1
fi
