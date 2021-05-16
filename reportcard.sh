#!/bin/bash

go get -u github.com/client9/misspell/cmd/misspell
git clone https://github.com/gojp/goreportcard.git
cd goreportcard
make install
go install ./cmd/goreportcard-cli
cd .. && rm -rf goreportcard
cd /github/workspace
ls
echo "RUNNING REPORT CARD"
goreportcard-cli -v
reportCard=`goreportcard-cli`
echo $reportCard
cd ../..
