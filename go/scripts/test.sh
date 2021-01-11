#!/usr/bin/env bash

# Source the common.sh script
# shellcheck source=./common.sh
. "$(git rev-parse --show-toplevel || echo ".")/go/scripts/common.sh"

cd "$PROJECT_DIR" || exit 1

if has richgo; then
  GO_TEST="richgo"
else
  GO_TEST="go"
fi

TEST_PATTERN="$*"
if [[ -z "$TEST_PATTERN" ]]; then
  echo_info "Test all packages"
  $GO_TEST test -cover -race ./...
else
  echo_info "Test with specified pattern: '$TEST_PATTERN'"
  $GO_TEST test -cover -race ./... -run "$TEST_PATTERN"
fi

EXIT_CODE=$?
cd "$WORKING_DIR" || exit 1
exit $EXIT_CODE
