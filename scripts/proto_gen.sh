#!/bin/sh
# Ref: https://studygolang.com/articles/25743

# protoc -I=../src -I=$GOPATH/src --go_out=../src/messages ../src/tenant/messages/*.proto
protoc -I=../src -I=$GOPATH/src --go_out=../ ../src/tenant/messages/*.proto
protoc -I=../src -I=$GOPATH/src --go_out=../ ../src/verifier/messages/*.proto