# syntax=docker/dockerfile:1

##
## Build
##
FROM golang:1.16-buster AS build

WORKDIR /app

COPY ./go.mod ./
COPY ./go.sum ./
RUN go mod download

COPY ./cmd/proxy/*.go ./
COPY ./internal internal/
RUN go build -o /proxy

##
## Deploy
##
FROM gcr.io/distroless/base-debian10

WORKDIR /

#COPY ./configs configs/

COPY --from=build /proxy /proxy

EXPOSE 9000

USER nonroot:nonroot

ENTRYPOINT ["/proxy"]