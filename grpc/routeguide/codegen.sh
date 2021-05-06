#!/bin/bash
# This script serves as an example to demonstrate how to generate the gRPC-Go
# interface and the related messages from .grpc file.
#
# It assumes the installation of i) Google grpc buffer compiler at
# https://github.com/google/protobuf (after v2.6.1) and ii) the Go codegen
# plugin at https://github.com/golang/protobuf (after 2015-02-20). If you have
# not, please install them first.
#
# We recommend running this script at $GOPATH/src.
#
# If this is not what you need, feel free to make your own scripts. Again, this
# script is for demonstration purpose.
#
WORKSPACE=$(cd $(dirname $0)/; pwd)
cd $WORKSPACE
files=`ls ./*.proto`
echo $files
protoc -I ./ -I $GOPATH/src/ --go_out=plugins=grpc:. $files