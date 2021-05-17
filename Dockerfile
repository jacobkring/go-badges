# Container image that runs your code
FROM golang:1.16-alpine AS build

WORKDIR /go/go-badges

COPY . ./

RUN go build

FROM alpine:3.13

WORKDIR /

RUN apk update
RUN apk add git

# Copies your code file from your action repository to the filesystem path `/` of the container
COPY --from=build /go/go-badges/go-badges /go-badges

# Code file to execute when the docker container starts up (`commit.sh`)
ENTRYPOINT ["/go-badges"]
