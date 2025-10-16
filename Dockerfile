# Build stage
FROM golang:1.24.6 AS builder

WORKDIR /app

# Cache dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source
COPY . .

# Build with optimizations
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o main .

# Runtime stage
FROM gcr.io/distroless/base-debian12

WORKDIR /app

# Copy binary from builder
COPY --from=builder /app/main .

# Use non-root user for security
USER nonroot:nonroot

# Run the binary
CMD ["./main"]
