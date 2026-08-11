# User Service

Go + Gin service that owns the seeded customer data.

## API

- `GET /health`
- `GET /api/v1/users`
- `GET /api/v1/users/{id}`

The service listens on port `8081` by default. Override it with the `PORT` environment variable.

`GET /health` returns `status`, `service` and the deployed `version`.
Set `APP_VERSION` at runtime; it defaults to `1.0.0` for local runs.

## Run and test

```bash
go run ./cmd/server
go test ./...
```

## Container

```bash
docker build -t user-service .
```
