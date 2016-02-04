#!/bin/bash

echo "Compiling for linux..."
GOOS=linux GOARCH=amd64 go build .

echo "Constructing Docker image"
docker build -t chambbj/diego-sample-go .
docker push chambbj/diego-sample-go

echo "Cleaning up..."
rm diego-sample-go
