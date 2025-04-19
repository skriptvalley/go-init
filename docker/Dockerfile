# ---- Build Stage ----
FROM golang:1.21-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /go-init ./cmd/server

# ---- Runtime Stage ----
FROM alpine:latest

RUN adduser -D -g '' appuser
USER appuser

COPY --from=builder /go-init /go-init

EXPOSE 8080

ENTRYPOINT ["/go-init"]
