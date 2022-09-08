# syntax=docker/dockerfile:1

##
## Build
##
##FROM golang:1.18-buster AS build
FROM golang:latest AS build

WORKDIR /app

COPY . ./
RUN go mod tidy

RUN go build -v

##
## Deploy
##
FROM gcr.io/distroless/base-debian10:latest
#FROM golang:latest

WORKDIR /

COPY --from=build /app/address-book-api /address-book-api

EXPOSE 8080

USER 1001

ENTRYPOINT ["/address-book-api"]
