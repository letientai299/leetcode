#!/usr/bin/env bash

# Source the common.sh script
# shellcheck source=./common.sh
. "$(git rev-parse --show-toplevel || echo ".")/go/scripts/common.sh"

build() {
  cmd=$1
  echo_info "Building $cmd"
  target=$(echo $cmd | sed 's/cmd/bin/')
  go build -ldflags="$GO_LDFLAGS" -i -v -o $target $cmd
  ls -lah -d $target
}

build_dir() {
  dir=$1
  # Build
  for cmd_package in $(find $dir -type d); do
    # skip the folder if there's no go file in it, or the package is not "main"
    if ! grep "package main" "$cmd_package"/*.go >/dev/null 2>&1; then
      continue
    fi

    # build the cmd
    build "$cmd_package"
  done
}

check_target() {
  cmd="$1"
  if [ -z $1 ]; then
    usage
    exit
  fi

  target="./cmd/$cmd"
  if [ -d "$target" ]; then
    build_dir "$target"
  else
    echo_warn "$cmd is not exists under 'cmd'! Here's how to use build script"
    echo
    usage
  fi
}

usage() {
  cat <<EOF
Build binary artifacts this service. Build target could be any folder under the
'cmd' one, such as:

  ./script/build.sh server

Or, if you trigger this script via make:

  make build/server

You can also build all packages under 'cmd' into binaries by a special target:
'all'.

  make build/all

To do cross compilation, such as build binary for Linux while working on Mac

  GOOS=linux make build/all

See "make gen" if you are looking for code generation. Note that some build
recipes might fail if the required generated code is missing.
EOF
}

cd "$PROJECT_DIR" || exit 1

case "$1" in
  all)
    build_dir ./cmd
    exit
    ;;
  -h | --help)
    usage
    exit
    ;;
  *)
    check_target "$@"
    exit
    ;;
esac

cd "$WORKING_DIR" || exit 1

