#! /bin/sh -x
go mod download
go install github.com/cosmtrek/air@latest
cp .env.sample .env