#!/bin/bash

# Set module path for Go code generation
# GO_MODULE_PATH="github.com/Akash-Manikandan/blogger-service"

# # Generate Go code
# cd blogger-service
# protoc --proto_path=../protobuf \
#     --go_out=. --go_opt=module=${GO_MODULE_PATH} \
#     --go-grpc_out=. --go-grpc_opt=module=${GO_MODULE_PATH} \
#     --validate_out="lang=go:." --validate_opt=paths=source_relative \
#     ../protobuf/*.proto

# Generate TypeScript code for frontend
# cd ../blogger-app
# PROTOC_GEN_TS_PATH="./node_modules/.bin/protoc-gen-ts"
# PROTOC_GEN_GRPC_PATH="./node_modules/.bin/grpc_tools_node_protoc_plugin"

# mkdir -p protobuf/generated

# # Generate TypeScript definitions using @grpc/proto-loader
# npx proto-loader-gen-types \
#     --longs=String \
#     --enums=String \
#     --defaults \
#     --oneofs \
#     --grpcLib=@grpc/grpc-js \
#     --outDir=protobuf/generated \
#     ../protobuf/blogger.proto
