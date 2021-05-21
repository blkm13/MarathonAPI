#!/bin/sh

echo "Ready to work"

apt-get update
apt-get -y install postgresql-client
# shellcheck disable=SC2037
PGPASSWORD=12340 psql -U postgres -h db  -c "CREATE DATABASE marathon"
tern migrate

go run src/main.go

