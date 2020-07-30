#!/bin/bash

#
# Run the CI flow and build the binary
# Prerequisite -
# 1. Go runtime
#

# absolute directory
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

ALL_PKGS=""

cd $DIR/../src

echo run go build
mkdir -p ${DIR}/../bin
rm -f ${DIR}/../bin/trigger-only
go build -o ${DIR}/../bin/trigger-only .
