# # Use the official Golang image as a parent image
# FROM golang:1.21.1-alpine AS builder

# # Set the Current Working Directory inside the container to /build
# WORKDIR /build

# # copy go.mod and go.sum to the container
# COPY go.mod go.sum ./
# # Install any Go dependencies
# RUN go mod download -x

# # Copy the Go application source code into the container
# COPY . .

# # Set necessary environment variables needed for our image and build the API server.
# ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
# RUN go build -ldflags="-s -w" -o service . # run is the same as run a command in terminal


# # Use scratch image as a parent image
# FROM scratch

# # COPY sql ./sql

# # Copy binary and config files from /build to root folder of scratch container.
# # COPY --from=builder ["/build/service", "/build/.env", "/"]
# COPY --from=builder ["/build/service",  "/"]

# # Command to run when the container started.
# #using ENTRYPOINT
# # --> run main command/binary. after start will search additional argument
# ENTRYPOINT ["/service"]

# #using CMD, run any command when container start using full comment
# # --> the same with running 'go run main.go' in terminal
# # CMD ["go", "run", "main.go"]


FROM golang:1.21.1-alpine AS builder
RUN mkdir /app
ADD ./ /app/
WORKDIR /app
RUN  go install github.com/cosmtrek/air@latest
ENTRYPOINT ["air"]