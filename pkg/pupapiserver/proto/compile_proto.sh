#!/usr/bin/env bash

protoc -I ../../pupapi \
    -I${GOPATH}/src \
    --go_out=plugins=grpc:../../pupapi \
    ../../pupapi/proto/*.proto
