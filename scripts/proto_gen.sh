#!/bin/sh
# Ref: https://studygolang.com/articles/25743
# https://colobu.com/2019/10/03/protobuf-ultimate-tutorial-in-go/#%E4%B8%80%E4%B8%AA%E7%AE%80%E5%8D%95%E7%9A%84%E4%BE%8B%E5%AD%90
protoc -I=. -I=$GOPATH/src --go_out=. *.proto