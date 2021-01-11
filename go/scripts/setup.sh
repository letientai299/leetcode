#!/usr/bin/env bash

# Source the common.sh script
# shellcheck source=./common.sh
. "$(git rev-parse --show-toplevel || echo ".")/go/scripts/common.sh"

cd "$PROJECT_DIR" || exit 1

# Mandatory tools
#-------------------------------------------------------------------------------

echo_info "Download golang dependencies"
go get ./...

if ! has goimports; then
  echo_info "Install goimports"
  go get -v -u golang.org/x/tools/cmd/goimports
fi

# Nice to have tools, should only be installed when not on CI, to save build time
#-------------------------------------------------------------------------------
if [[ -z $CI ]]; then
  # richgo provides colored output when running test.
  if ! has richgo; then
    echo_info "Install richgo for nicer go test output"
    go get -v -u github.com/kyoh86/richgo
  fi

  if ! has packr2; then
    echo_info "Install packr2"
    go get -u github.com/gobuffalo/packr/v2/packr2
  fi
fi

echo_info "Config git hooks push"
git config core.hooksPath "${PROJECT_DIR}/scripts/git-hooks"

cd "$WORKING_DIR" || exit 1
