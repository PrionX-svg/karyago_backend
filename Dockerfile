# Build stage
FROM golang:1.24.6 AS builder

WORKDIR /app

# Install libwebp (required for chai2010/webp)
RUN apt-get update && apt-get install -y libwebp-dev && rm -rf /var/lib/apt/lists/*

# Copy dependencies and download
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build binary with CGO enabled
RUN CGO_ENABLED=1 GOOS=linux go build -ldflags="-s -w" -o main .

# Runtime stage
FROM gcr.io/distroless/base-debian12

WORKDIR /app

# Copy binary and necessary shared libs
COPY --from=builder /app/main .
COPY --from=builder /usr/lib/x86_64-linux-gnu/libwebp.so.7 /usr/lib/x86_64-linux-gnu/

# Use non-root user for security
USER nonroot:nonroot

# Run the binary
CMD ["./main"]