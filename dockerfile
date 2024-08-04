# Start from a Go base image
FROM golang:1.16-alpine AS builder

# Set the working directory
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o deeplx-load-balancer ./cmd/server

# Start a new stage from scratch
FROM alpine:latest  

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the pre-built binary file from the previous stage
COPY --from=builder /app/deeplx-load-balancer .
COPY --from=builder /app/config/config.yaml ./config/

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./deeplx-load-balancer"]