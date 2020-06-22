#!/bin/sh
# Ref: https://studygolang.com/articles/25743
# https://colobu.com/2019/10/03/protobuf-ultimate-tutorial-in-go/#%E4%B8%80%E4%B8%AA%E7%AE%80%E5%8D%95%E7%9A%84%E4%BE%8B%E5%AD%90
# export GOPATH=/path/to/go (e.g. /home/sunflower/go)
# in out-of-your-project/
# protoc -I=. -I=Reactive-Welfare-Housing-System/src/messages -I=$GOPATH/src --go_out=plugins=grpc:. Reactive-Welfare-Housing-System/src/messages/*.proto
cd $RWHS_HOME/../
protoc -I=. -I=Reactive-Welfare-Housing-System/src/messages -I=$GOPATH/src --gogoslick_out=plugins=grpc:. Reactive-Welfare-Housing-System/src/messages/government.proto
protoc -I=. -I=Reactive-Welfare-Housing-System/src/messages -I=$GOPATH/src --gogoslick_out=plugins=grpc:. Reactive-Welfare-Housing-System/src/messages/property.proto
protoc -I=. -I=Reactive-Welfare-Housing-System/src/messages -I=$GOPATH/src --gogoslick_out=plugins=grpc:. Reactive-Welfare-Housing-System/src/messages/tenant.proto
protoc -I=. -I=Reactive-Welfare-Housing-System/src/messages -I=$GOPATH/src --gogoslick_out=plugins=grpc:. Reactive-Welfare-Housing-System/src/messages/shared.proto
protoc -I=. -I=Reactive-Welfare-Housing-System/src/messages -I=$GOPATH/src --gogoslick_out=plugins=grpc:. Reactive-Welfare-Housing-System/src/messages/verifier.proto
protoc -I=. -I=Reactive-Welfare-Housing-System/src/messages -I=$GOPATH/src --gogoslick_out=plugins=grpc:. Reactive-Welfare-Housing-System/src/messages/distributor.proto
protoc -I=. -I=Reactive-Welfare-Housing-System/src/messages -I=$GOPATH/src --gogoslick_out=plugins=grpc:. Reactive-Welfare-Housing-System/src/messages/manager.proto
