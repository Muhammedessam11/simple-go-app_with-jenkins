# Build stage
FROM golang:1.17-alpine AS builder

# Set working directory
WORKDIR /app

# Copy source files and build the app
COPY . .
RUN go mod init simple-go-app && go build -o simple-go-app

# Deployment stage
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/simple-go-app .

# Expose the app port and run
EXPOSE 8080
CMD ["./simple-go-app"]

