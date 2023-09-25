# Stage 1: Production environment (สำหรับ production)
FROM golang:1.21.1-alpine as production

WORKDIR /app

# Copy go.mod and go.sum to download dependencies
COPY go.mod go.sum ./

# Download and install dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go application with production settings
RUN CGO_ENABLED=0 GOOS=linux go build -o app-prod

# Stage 2: Final lightweight image
FROM alpine:latest

WORKDIR /app

# Copy the built binary from the appropriate stage based on the build argument
COPY --from=production /app/app-prod ./app-prod

# Expose the port your Go Fiber application listens on (default is 3000)
EXPOSE 3000

# Run the Go Fiber application
CMD ["./app"]