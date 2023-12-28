# Hatchet: Go SDK Quickstart

This repository contains a simple application which showcases how to use the [Hatchet](https://github.com/hatchet-dev/hatchet) Go SDK to create events and workflows. In [./cmd/server](./cmd/server/main.go), you'll find a simple [`echo`](https://echo.labstack.com/) server which pushes an event to Hatchet every time it receives a request on `/test`. In [./cmd/worker](./cmd/worker/worker.go), we create a worker which runs a simple workflow called `event-test` whenever the `test-called` event is seen.

## Getting Started

### Prerequisites

This quickstart example requires the following tools to work:

- `go 1.18+`
- [`docker-compose`](https://docs.docker.com/compose/install/)
- [`caddy`](https://caddyserver.com/docs/install)

### Get up and running

1. Run `go mod download` to fetch all packages.

2. Run `docker-compose up` to start the Hatchet instance. This will take a few minutes, as the docker-compose services set up the database and generate the required certificates to connect to the Hatchet instance. You can also run `docker-compose up -d` to start this in the background. Once you start to see output from the `engine` and `api` services, you can move on to the next step.

3. Run `caddy start` to get an instance running. You should be able to navigate to app.dev.hatchet-tools.com and use the following credentials to log in:

```
Email: admin@example.com
Password: Admin123!!
```

4. Create the required environment variables via:

```sh
cat > .env <<EOF
HATCHET_CLIENT_TENANT_ID=707d0855-80ab-4e1f-a156-f1c4546cbf52
HATCHET_CLIENT_TLS_ROOT_CA_FILE=./certs/ca.cert
HATCHET_CLIENT_TLS_CERT_FILE=./certs/client-worker.pem
HATCHET_CLIENT_TLS_KEY_FILE=./certs/client-worker.key
HATCHET_CLIENT_TLS_SERVER_NAME=cluster
EOF
```

5. Run the server and worker in two separate shell sessions via: `go run ./cmd/server` and `go run ./cmd/worker`.

6. Run `curl http://localhost:1323/test` to test the endpoint.

You will see events and workflows populated in the Hatchet dashboard:

<img width="1728" alt="image" src="https://github.com/hatchet-dev/hatchet-go-quickstart/assets/25448214/376e4ee8-7233-4a84-85b8-f71ad9e7402e">

