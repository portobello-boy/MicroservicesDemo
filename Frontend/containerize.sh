#!/bin/bash

# Echo Image/Container name
echo $1

docker build -t $1 .
docker rm $1
docker create -p 4200:4200 --name $1 $1
