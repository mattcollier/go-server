# Stage 1: Builder
FROM golang:1.25-alpine AS builder

WORKDIR /app

# Copy go.mod and go.sum first to leverage Docker's build cache
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/main .

# Stage 2: Final Image
FROM alpine:latest

# Install necessary certificates for HTTPS if your application makes external calls
RUN apk add --no-cache ca-certificates

WORKDIR /root/

# Copy the built binary from the builder stage
COPY --from=builder /app/main .

# Expose the port your application listens on
EXPOSE 8080

# Command to run the application
CMD ["./main"]
