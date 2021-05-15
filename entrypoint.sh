#!/bin/sh -l

ls

echo "Hello $1, $2, $3, $4, $5"

time=$(date)
echo "::set-output name=time::$time"