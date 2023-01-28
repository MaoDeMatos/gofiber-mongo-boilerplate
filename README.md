# Go Fiber

A simple API built with [`Gofiber`](https://gofiber.io/)

# Requirements

- Golang 1.19

# Usage

## Install

To use hot reload, [`Air`](https://github.com/cosmtrek/air) is preconfigured, but you can use anything you want

```sh
go mod download                             # Install projects dependencies
go install github.com/cosmtrek/air@latest   # For hot reload
```

## Start

```sh
go run main.go    # Simply start the app
air               # Start the app in dev (watch) mode
```
