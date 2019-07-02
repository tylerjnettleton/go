#!/usr/bin/env bash
protoc -I ../auth \
    -I${GOPATH}/src \
    --go_out=plugins=grpc:../auth \
    ../auth/proto/auth.proto