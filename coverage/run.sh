#!/bin/bash

cd ../
# go test -cover ./...
go test -coverprofile ./coverage/cover.out ./
go test -coverprofile ./coverage/cover.handlers.out ./handlers
go test -coverprofile ./coverage/cover.formatters.out ./formatters
cd -
go tool cover -html=cover.out -o cover.html
go tool cover -html=cover.handlers.out -o cover.handlers.html
go tool cover -html=cover.formatters.out -o cover.formatters.html
