# go-init

🚀 A production-ready Golang microservice starter kit.

## 📁 Project Structure

- `cmd/` - Entrypoint binaries
- `internal/` - Internal application logic
- `pkg/` - Shared packages (e.g., logger)
- `configs/` - Configuration files
- `deployments/` - Docker, K8s manifests
- `scripts/` - Dev or CI automation
- `vendor/` - Vendored dependencies (if using `go mod vendor`)

## 🚀 Quick Start (Local)

```bash
make run
```

## 🐳 Docker

### Build and Run:

```bash
docker build -t go-init .
docker run -p 8080:8080 go-init
```

Or use Docker Compose:

```bash
docker-compose up --build
```

App will be available at [http://localhost:8080](http://localhost:8080)

### Prometheus

Prometheus is included and scrapes metrics from `/metrics`.

- Prometheus UI: [http://localhost:9090](http://localhost:9090)
- App metrics: [http://localhost:8080/metrics](http://localhost:8080/metrics)

## ⚙️ Config

Settings are loaded from a `.env` file.

### Example `.env`

```
PORT=8080
LOG_LEVEL=debug
```

## 🧪 Health

```http
GET /healthz
```

Returns `200 OK` with body `ok`.

## 📈 Metrics

```http
GET /metrics
```

Prometheus-compatible metrics including:
- `http_requests_total`
- `http_request_duration_seconds`
