#!/bin/sh -l

goreportcard-cli -v
reportCard=`goreportcard-cli`
./go-badges
