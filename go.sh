#!/bin/bash

# 读取当前环境, 读取不同配置
source ./env.sh
appname=${PWD##*/}
cfgpath=${PWD}"/etc/unknown.yml"
echo $cfgpath

echo "go run ./src/"$appname"/main.go -f "$cfgpath
apidoc -i ./src/${appname}/controllers -o doc/
go run ./src/${appname}/main.go -f $cfgpath
