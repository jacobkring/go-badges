# Container image that runs your code
FROM golang:1.16-alpine AS build

COPY . ./

RUN go build

FROM scratch
# Copies your code file from your action repository to the filesystem path `/` of the container
COPY --from=build go-badges /go-badges

# Code file to execute when the docker container starts up (`entrypoint.sh`)
ENTRYPOINT ["/go-badges"]