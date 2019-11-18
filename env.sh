#!/bin/bash
# 设置临时 GOPATH

export GOPATH=${PWD}

thirdparty=${PWD%/*}"/thirdparty"
export GOPATH=$GOPATH:${thirdparty}

environment=${PWD%/*}"/environment"
export GOPATH=$GOPATH:${environment}

sgo=${PWD%/*}"/sgo"
export GOPATH=$GOPATH:${sgo}

ccgwf=${PWD%/*}"/ccgwf"
export GOPATH=$GOPATH:${ccgwf}

echo "GOPATH: "$GOPATH

