#!/bin/sh
set -e

go build -o todos -ldflags "cmd.binary_name=todos" ./main.go
mv ./todos /usr/local/bin/todos
