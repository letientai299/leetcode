#!/usr/bin/env bash

# Source the common.sh script
# shellcheck source=./common.sh
. "$(git rev-parse --show-toplevel || echo ".")/go/scripts/common.sh"

cd "$PROJECT_DIR" || exit 1

echo_info "Generate packr2 files"
cd ./cmd/create_file && packr2 clean && packr2
cd "$PROJECT_DIR" || exit

echo_info "Go generate"
go generate ./...

echo_info "Reformat everything with goimports"
goimports -w -e $(find . -type f -name '*.go' -not -path ".cache/*")

cd "$WORKING_DIR" || exit 1
