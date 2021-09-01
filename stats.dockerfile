# syntax=docker/dockerfile:1

##
## Build
##
FROM golang:1.16-buster AS build

WORKDIR /app

COPY ./go.mod ./
COPY ./go.sum ./
RUN go mod download

COPY ./cmd/client/*.go ./
COPY ./internal internal/
RUN go build -o /stats

##
## Deploy
##
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /stats /stats

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/stats"]