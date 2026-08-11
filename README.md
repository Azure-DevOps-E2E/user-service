# User Service

Go + Gin service that owns the seeded customer data.

## API

- `GET /health`
- `GET /api/v1/users`
- `GET /api/v1/users/{id}`

The service listens on port `8081` by default. Override it with the `PORT` environment variable.

## Run and test

```bash
go run ./cmd/server
go test ./...
```

## Container

```bash
docker build -t user-service .
```
