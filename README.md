# Go Fiber

A simple API built with [`gofiber`](https://gofiber.io/)

# Requirements

- Golang 1.19

# Usage

To use hot reload, you'll need to install [`Air`](https://github.com/cosmtrek/air)

## Scripts

Just run `make` of `make help` to see what commands you can use.

## Manual

You should use the makefile

### Install

```sh
go mod download                             # Install projects dependencies
go install github.com/cosmtrek/air@latest   # Install air, for hot reload
cp .env.sample .env                         # Setup default `env` vars
```

### Start

```sh
air                         # Start the app in dev (watch) mode

go build -o ./bin/main .    # Production build
./bin/main                  # Strt production build
```
