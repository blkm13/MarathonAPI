#!/bin/sh

echo "Ready to work"
PGPASSWORD=12340 psql -U postgres -h db  -c "CREATE DATABASE marathon;"
go run src/main.go
