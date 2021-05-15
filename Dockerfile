# Container image that runs your code
FROM golang:1.16-alpine AS build

WORKDIR /go/go-badges

COPY . ./

RUN go build

FROM golang:1.16-alpine
# Copies your code file from your action repository to the filesystem path `/` of the container
COPY --from=build /go/go-badges /go-badges
RUN ls
# Code file to execute when the docker container starts up (`entrypoint.sh`)
ENTRYPOINT ["./go-badges"]