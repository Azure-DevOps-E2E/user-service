FROM golang:1.26.6-alpine3.24 AS build

WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go test ./... \
    && CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags="-s -w" -o /out/user-service ./cmd/server

FROM alpine:3.24

RUN apk add --no-cache ca-certificates \
    && addgroup -S app \
    && adduser -S -G app app

COPY --from=build /out/user-service /usr/local/bin/user-service

USER app
EXPOSE 8081
ENTRYPOINT ["/usr/local/bin/user-service"]
