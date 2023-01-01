#!/usr/bin/env bash

set -eu

usage() {
  cat <<USAGE
  Usage:
  - e2e.sh up
  - e2e.sh run
  - e2e.sh down
USAGE
}

if [ "$#" != 1 ]; then
  usage
  exit 1
fi

command="$1"
repo_dir="$(git rev-parse --show-toplevel)"

up () {
  docker compose up -d db cache zookeeper queue mailhog app auth mail
}

run () {
  go test "${repo_dir}/cmd/e2e/main_test.go"
}

if [ "$command" == "up" ]; then
  up
elif [ "$command" == "run" ]; then
  up
  run
elif [ "$command" == "down" ]; then
  docker compose down
else
  usage
  exit 1
fi
