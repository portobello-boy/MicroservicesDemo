#!/bin/bash

# Echo Image/Container name
echo $1

docker build -t $1 .
docker rm $1
docker create -p 27017:27017 --name $1 $1
