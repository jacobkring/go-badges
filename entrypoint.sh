#!/bin/sh -l

goreportcard-cli -v
reportCard=`goreportcard-cli`
go run badges.go
