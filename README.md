# Go on Porter

A minimal Go web service using only the standard library, ready to deploy on Porter.

## What's inside

| Component | Version | Notes |
| --- | --- | --- |
| Go | 1.27.3 | `net/http` routing with method patterns, `log/slog` JSON logs |
| Builder image | `golang:1.27-trixie` | Static binary built with `CGO_ENABLED=0` |
| Runtime image | `gcr.io/distroless/static-debian13:nonroot` | Runs as a non-root user |

| Method | Path | Response |
| --- | --- | --- |
| GET | `/` | `{"message":"Hello from Go on Porter!"}` |
| GET | `/healthz` | `{"status":"ok"}` |

The server listens on `0.0.0.0:$PORT` (default `8080`) and shuts down gracefully on `SIGTERM`.

## Run locally

```sh
go run .
go test ./...
```

With Docker:

```sh
docker build -t golang-getting-started . && docker run -p 8080:8080 golang-getting-started
```

Then open http://localhost:8080.

## Deploy on Porter

### From the dashboard

1. Fork this repo.
2. In https://dashboard.porter.run, create a new app from your GitHub repo.
3. Porter reads `porter.yaml`, builds the Dockerfile, and opens a PR adding a GitHub Actions workflow.
4. Merge the PR. Every push to `main` redeploys.

See https://docs.porter.run/applications/deploy/deploy-from-github-repo for details.

### From the CLI

1. Install the CLI: https://docs.porter.run/cli/installation
2. Log in with `porter auth login`.
3. Deploy with `porter apply -f porter.yaml`.

## Configuration

`porter.yaml` builds the image from `./Dockerfile` and runs one `web` service:

| Field | Value | Purpose |
| --- | --- | --- |
| `run` | `/app/server` | Start command, same as the Dockerfile `CMD` |
| `port` | `8080` | Port the service receives traffic on |
| `cpuCores` | `0.25` | CPU allocated per instance |
| `ramMegabytes` | `128` | Memory allocated per instance |
| `healthCheck.httpPath` | `/healthz` | Path Porter polls to confirm the service is healthy |

| Name | Required | Description |
| --- | --- | --- |
| `PORT` | No | Port to listen on. Defaults to `8080`, which matches `porter.yaml`. |

Full reference: https://docs.porter.run/applications/configuration-as-code/reference

## Learn more

- https://docs.porter.run/getting-started/quickstart
- https://docs.porter.run/applications/configuration-as-code/overview
- https://docs.porter.run/applications/configure/environment-groups
- https://docs.porter.run/addons/datastores
- https://docs.porter.run/preview-environments/overview
- https://github.com/porter-dev/starters
