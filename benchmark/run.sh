#!/bin/bash

echo "Benchmarks"

echo "Running bash script (iterating over each cluster and using grep)"
time _=$(./cmd.sh "$1")

echo -e "\n-----\nRunning kube-s\n"
time _=$(go run ../main.go "$1")
