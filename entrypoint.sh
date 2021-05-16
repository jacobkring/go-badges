#!/bin/sh -l

git clone https://github.com/gojp/goreportcard.git && \
  cd goreportcard && \
  make install && \
  go install ./cmd/goreportcard-cli && \
  goreportcard-cli && \
  cd ..
pwd
ls

cd /github/workspace
pwd
ls
goreportcard-cli -v
cd ../..
reportCard=`goreportcard-cli`
./go-badges
