#!/bin/bash

cd /github/workspace
echo "RUNNING REPORT CARD"
goreportcard-cli -v
reportCard=`goreportcard-cli`
echo $reportCard
cd ../..
