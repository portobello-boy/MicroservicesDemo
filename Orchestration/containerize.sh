#!/bin/bash

# Echo Image/Container name
echo $1

docker build -t $1 .
docker rm $1
docker create -p 3001:3001 --name $1 $1
