#! /bin/bash
set -e

readonly workdir="/app"
readonly appname="$1"
readonly appdir="/$(dirname "$1")"
mkdir -p  "$appdir"
cd "$workdir"
go build -o "$1" "./cmd$appname"
exec "$@"
