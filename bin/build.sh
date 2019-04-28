#! /bin/bash
set -e

if [[ -n "${DEBUG}" ]]; then
    set -x
fi

mkdir -p /`dirname $1`

echo "$(date): Building cmd$1"
cd /app
go build -o "$1" "./cmd$1"

touch /app/app.env

echo "$(date): Executing $*"
exec "$@"
