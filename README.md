# Go Fiber

A simple API built with [`gofiber`](https://gofiber.io/)

# Requirements

- Golang 1.19

# Usage

## Install

To use hot reload, [`Air`](https://github.com/cosmtrek/air) is preconfigured, but you can use anything you want

Just run `./install.sh` to install dependencies and generate `.env`, or use this commands to do so manually :

```sh
go mod download                             # Install projects dependencies
go install github.com/cosmtrek/air@latest   # For hot reload
cp .env.sample .env                         # Setup default `env` vars
```

## Start

```sh
go run main.go    # Simply start the app
air               # Start the app in dev (watch) mode
```
