#!/bin/bash

if ! command -v protoc &> /dev/null; then
    echo "Error: protoc is not installed. Please install it and try again."
    exit 1
fi

SCRIPT_DIR=$(dirname "$(realpath "$0")")

PROTO_DIR="$SCRIPT_DIR/../protos"
AI_PROXY_DIR="$SCRIPT_DIR/../ai-proxy"
PROTO_CODEGEN_DIR="$SCRIPT_DIR/../common/proto"

mkdir -p "$AI_PROXY_DIR" "$PROTO_CODEGEN_DIR"

protoc -I="$PROTO_DIR" --python_out="$AI_PROXY_DIR" "$PROTO_DIR/llmproxy.proto"
protoc -I="$PROTO_DIR" --python_out="$AI_PROXY_DIR" "$PROTO_DIR/sso.proto"

export GOPATH=$PROTO_CODEGEN_DIR
export PATH=$PATH:$GOPATH/bin
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

protoc -I="$PROTO_DIR" --go_out="$PROTO_CODEGEN_DIR" --go-grpc_out="$PROTO_CODEGEN_DIR" "$PROTO_DIR/llmproxy.proto"
protoc -I="$PROTO_DIR" --go_out="$PROTO_CODEGEN_DIR" --go-grpc_out="$PROTO_CODEGEN_DIR" "$PROTO_DIR/sso.proto"

echo "Protobuf files have been successfully compiled."
echo "Protobuf from ${PROTO_DIR} was created in ${AI_PROXY_DIR} and $PROTO_CODEGEN_DIR"
