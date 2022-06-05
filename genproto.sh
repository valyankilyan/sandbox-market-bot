#!/bin/bash

protoc \
--go_out=pkg/api \
--go_opt=paths=source_relative \
--go-grpc_out=pkg/api \
--go-grpc_opt=paths=source_relative \
api/server.proto