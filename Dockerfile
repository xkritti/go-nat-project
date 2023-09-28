
###########################
# STEP 1 build executable binary
###########################
FROM golang:1.21.1-alpine  as production

WORKDIR /app
COPY . .
# Fetch dependencies.
# Using go mod with go 1.11
RUN go mod download
RUN go mod verify
# Build the binary
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/app .

############################
# STEP 2 build a small image
############################
FROM alpine:latest

# COPY .env .env
# Import from builder.
COPY --from=production /app/app /go/bin/app-prod

# EXPOSE 3000

ENTRYPOINT ["/go/bin/app-prod"]
