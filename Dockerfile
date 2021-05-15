# Container image that runs your code
FROM scratch AS build
COPY --from=golang:1.16-alpine /usr/local/go/ /usr/local/go/

COPY . ./

RUN go build

FROM scratch
# Copies your code file from your action repository to the filesystem path `/` of the container
COPY --from=build go-badges /go-badges

# Code file to execute when the docker container starts up (`entrypoint.sh`)
ENTRYPOINT ["go-badges"]