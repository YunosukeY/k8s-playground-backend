#!/bin/bash
set -eu

repo_dir="$(git rev-parse --show-toplevel)"

sh "${repo_dir}/script/wait-db.sh"

go run "${repo_dir}/cmd/migration/main.go" -o up -t schema
go run "${repo_dir}/cmd/migration/main.go" -o down -t record
go run "${repo_dir}/cmd/migration/main.go" -o up -t record

go test "${repo_dir}/internal/..."
