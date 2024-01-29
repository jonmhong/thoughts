#!/bin/sh
set -e

go build
mv ./thoughts /usr/local/bin
