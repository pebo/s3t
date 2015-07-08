#!/bin/bash -e

# build linx amd64 image: s3t
docker run --rm -v $PWD:/src -w /src golang:1.4 make
