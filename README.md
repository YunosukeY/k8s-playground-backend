# K8s Playground's Backend App

[![build backend](https://github.com/YunosukeY/k8s-playground-backend/actions/workflows/build-backend.yaml/badge.svg?branch=master&event=push)](https://github.com/YunosukeY/k8s-playground-backend/actions/workflows/build-backend.yaml)
[![golangci-lint](https://github.com/YunosukeY/k8s-playground-backend/actions/workflows/golangci-lint.yml/badge.svg?branch=master&event=push)](https://github.com/YunosukeY/k8s-playground-backend/actions/workflows/golangci-lint.yml)
[![e2e](https://github.com/YunosukeY/k8s-playground-backend/actions/workflows/e2e.yaml/badge.svg?branch=master&event=push)](https://github.com/YunosukeY/k8s-playground-backend/actions/workflows/e2e.yaml)

<!--
## Local Usage

1. Create `.env` file.

```bash
cat <<EOF > .env
MYSQL_ROOT_PASSWORD={ROOT_PASSWORD}
MYSQL_DATABASE={DATABASE_NAME}
MYSQL_USER={USER}
MYSQL_PASSWORD={PASSWORD}
REDIS_PASSWORD={PASSWORD}
EOF
```

2. Start apps.

```bash
./e2e.sh up
```

## Preparation for Kind Sample

0. Register to DockerHub if you have never used.
1. Create a repository "kind-backend" at DockerHub
2. Fork this repository.
3. Add your DockerHub username and password as `DOCKERHUB_USERNAME` and `DOCKERHUB_PASSWORD` to the repository secret.
4. Run `build backend` action. -->
