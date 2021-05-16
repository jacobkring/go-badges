# Container image that runs your code
FROM golang:1.16-alpine AS build

WORKDIR /go/go-badges

RUN apk update
RUN apk add git

COPY . ./

RUN go build

FROM golang:1.16-alpine

WORKDIR /

RUN git clone https://github.com/gojp/goreportcard.git \
    cd goreportcard \
    make install \
    go install ./cmd/goreportcard-cli \
    goreportcard-cli \
    cd ..

# Copies your code file from your action repository to the filesystem path `/` of the container
COPY --from=build /go/go-badges/ .

# Code file to execute when the docker container starts up (`entrypoint.sh`)
ENTRYPOINT ["/go-badges"]
