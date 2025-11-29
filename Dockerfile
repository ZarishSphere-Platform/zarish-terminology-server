# Build Stage
FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o terminology cmd/terminology/main.go

# Runtime Stage
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/terminology .

# Copy data directory for loading artifacts
# Assuming data is mounted or copied. For now, let's copy the data module if needed, 
# but in K8s we likely mount it. 
# However, the loader expects data relative to the binary or via env.
# Let's assume we copy the necessary data into the image for self-containment or use a volume.
# For simplicity in this Dockerfile, we'll rely on a volume mount at runtime for /data 
# and set DATA_DIR env var.

ENV DATA_DIR=/app/data

EXPOSE 8081

CMD ["./terminology"]
