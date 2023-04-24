#!/bin/sh
cd "${0%/*}"
go run knt/cmd/main.go -tags "libsqlite3"