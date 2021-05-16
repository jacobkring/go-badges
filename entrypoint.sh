#!/bin/sh -l

ls
cd /github/workspace
ls
goreportcard-cli -v
cd ../..
reportCard=`goreportcard-cli`
./go-badges
