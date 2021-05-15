#!/bin/sh -l

echo "Hello $1, $2, $3, $4, $5"

go run badges.go $1 $2 $3 $4 $5

time=$(date)
echo "::set-output name=time::$time"