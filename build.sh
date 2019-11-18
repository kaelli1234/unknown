#!/bin/bash
# 编译
# go build -X 变量值中不能有空格

source ./env.sh
appName=${PWD##*/}

BuildTime=$(date)
BuildTime=${BuildTime// /_}
echo "BuildTime: "$BuildTime

BuildUser=$(whoami)
BuildUser=${BuildUser// /_}
echo "BuildUser: "$BuildUser

# BuildVersion="aaa "`git rev-list HEAD -n 1 | cut -c 1-7`

git rev-list HEAD | sort > config.git-hash
LOCALVER=`wc -l config.git-hash | awk '{print $1}'`
if [ $LOCALVER \> 1 ] ; then
    VER=`git rev-list origin/master | sort | join config.git-hash - | wc -l | awk '{print $1}'`
    if [ $VER != $LOCALVER ] ; then
        VER="$VER+$(($LOCALVER-$VER))"
    fi
    if git status | grep -q "modified:" ; then
        VER="${VER}M"
    fi
    VER="$VER $(git rev-list HEAD -n 1 | cut -c 1-7)"
    GIT_VERSION=r$VER
fi
rm -f config.git-hash
BuildVersion=$GIT_VERSION
BuildVersion=${BuildVersion// /_}
echo "BuildVersion: "$BuildVersion

BuildMachine=$(/sbin/ifconfig | grep "inet" | grep -v "127.0.0.1" | grep -v "inet6" | awk '{print $2}'| tr "\n" " ")
BuildMachine=${BuildMachine// /_}
echo "BuildMachine: "$BuildMachine

BuildGcc=$(gcc --version | head -n 1 | sed 's/[()]//g')
BuildGcc=${BuildGcc// /_}
echo "BuildGcc: "$BuildGcc

go build -ldflags " -X main.BuildTime=${BuildTime} -X main.BuildUser=${BuildUser} -X main.BuildVersion=${BuildVersion} -X main.BuildMachine=${BuildMachine} -X main.BuildGcc=${BuildGcc}" -o ./bin/${appName} ./src/${appName}/main.go

echo ${appName}" build done"
