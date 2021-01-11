#!/usr/bin/env bash
# This scripts provide some helpful method and predefined shell variables. This
# is not intended to be used by its own. Use should source this script into
# other scripts.

# Backup working dir when the script is sourced. This value is helpful to change
# the shell dir back to the original working dir, as some command might need to
# be run in other dir.
WORKING_DIR=$PWD

# Detect scripts dir and go up one level to be at root project dir
SCRIPTS_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
# SCRIPTS_DIR=$(cwd=$(pwd);ewd=$(dirname "$0");cd "${ewd}";pwd;cd "${cwd}")

# Project dir is one level up from SCRIPT_DIR
PROJECT_DIR="$(cd "$SCRIPTS_DIR/.." && pwd)"

echo_info() {
  color=$(tput setaf 2)
  reset=$(tput sgr0)
  echo "${color}[INFO] $*${reset}"
}

echo_warn() {
  color=$(tput setaf 3)
  reset=$(tput sgr0)
  echo "${color}[WARN] $*${reset}"
}

echo_error() {
  # Visit this page for tput look up color code
  # https://en.wikipedia.org/wiki/ANSI_escape_code#8-bit
  color=$(tput setaf 208)
  reset=$(tput sgr0)
  echo "${color}[ERROR] $*${reset}"
}

exit_script() {
  echo_warn "Cancelling current execution on user input"
  cd "$WORKING_DIR" || exit 1
  exit
}

has() {
  command -v "$1" >/dev/null 2>&1
}

# Force go module enable, despite whether the project folder locates
export GO111MODULE=on

# Make sure that the script that source this script will stop everything and
# exit on Ctrl-C. (SIGINT)
trap exit_script INT
