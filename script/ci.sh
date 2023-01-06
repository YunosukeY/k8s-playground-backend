#!/bin/bash
set -eu

repo_dir="$(git rev-parse --show-toplevel)"

sh "${repo_dir}/script/wait-db.sh"

go test "${repo_dir}/internal/..."
