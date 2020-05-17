#!/bin/sh
# Ref: https://studygolang.com/articles/25743

# protoc -I=../src -I=$GOPATH/src --go_out=../src/messages ../src/tenant/messages/*.proto
protoc -I=../src -I=$GOPATH/src --go_out=../ ../src/tenant/messages/*.proto
protoc -I=../src -I=$GOPATH/src --go_out=../ ../src/verifier/messages/*.proto
# protoc -I=../src -I=$GOPATH/src --gogoslick_out=plugins=grpc:../ ../src/tenant/messages/*.proto
# protoc -I=../src -I=$GOPATH/src --gogoslick_out=plugins=grpc:../ ../src/verifier/messages/*.proto