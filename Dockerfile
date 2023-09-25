FROM golang:1.21.1-alpine AS builder
RUN mkdir /app
ADD ./ /app/
WORKDIR /app
RUN  go install github.com/cosmtrek/air@latest
ENTRYPOINT ["air"]