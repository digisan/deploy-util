#!/bin/bash

set -e

R=`tput setaf 1`
G=`tput setaf 2`
Y=`tput setaf 3`
W=`tput sgr0`

./clean.sh

LDFLAGS="-s -w"
TM=`date +%F@%T@%Z`
OUT=deploy\($TM\)

GOARCH=amd64
CGO_ENABLED=0 GOOS="linux" GOARCH="$GOARCH" go build -ldflags="$LDFLAGS" -o $OUT
echo "${G}${OUT}(linux64)${W} built"