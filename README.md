# K8s Playground's Backend App

[![ci](https://github.com/YunosukeY/k8s-playground-backend/actions/workflows/ci.yaml/badge.svg?branch=master&event=push)](https://github.com/YunosukeY/k8s-playground-backend/actions/workflows/ci.yaml)
[![e2e](https://github.com/YunosukeY/k8s-playground-backend/actions/workflows/e2e.yaml/badge.svg?branch=master&event=push)](https://github.com/YunosukeY/k8s-playground-backend/actions/workflows/e2e.yaml)
[![build backend](https://github.com/YunosukeY/k8s-playground-backend/actions/workflows/build-backend.yaml/badge.svg?branch=master&event=push)](https://github.com/YunosukeY/k8s-playground-backend/actions/workflows/build-backend.yaml)
[![golangci-lint](https://github.com/YunosukeY/k8s-playground-backend/actions/workflows/golangci-lint.yml/badge.svg?branch=master&event=push)](https://github.com/YunosukeY/k8s-playground-backend/actions/workflows/golangci-lint.yml)
[![Renovate](https://img.shields.io/badge/renovate-enabled-brightgreen.svg)](https://renovatebot.com)

## Features

### Observability

- OpenTelemetry Traces
- Exposing Metrics
- Structured Logs

### Container Security

- Distroless Image
- Nonroot User

### API Mode

- REST
- gRPC

## Architecture

This sample application consists of three services: app, auth, and mail.<br>

### Authentication

Sessions are created and checked via the auth service.

The app service has two kinds of endpoints: public endpoints, and private endpoints.<br>
Public endpoints can be accessed without authentication, and private endpoints need authentication.

Since these services depend on Ingress-NGINX's external authentication,<br>
in local use, you just need to add `X-Auth` header to use private endpoints.

### Sending Mails

Sending mails are handled asynchronously.<br>
The app service receives sending requests at `/api/v1/mails`, and enqueues messages.<br>
The mail service dequeues messages, and sends e-mails.

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
./script/e2e.sh up
```

## Preparation for Kind Sample

0. Register to DockerHub if you have never used.
1. Create a repository "kind-backend" at DockerHub
2. Fork this repository.
3. Add your DockerHub username and password as `DOCKERHUB_USERNAME` and `DOCKERHUB_PASSWORD` to the repository secret.
4. Run `build backend` action. -->
