# PostfixAdmin Fork Refacture in Golang

[![Release](https://github.com/jniltinho/postfixadmin/workflows/Release/badge.svg)](https://github.com/jniltinho/postfixadmin/actions?query=workflow%3ARelease)

## Disclaimer
This project is currently not working, it is just the skeleton of the beginning of the project.

## This project uses the tools below.

- Echo
- Templ
- Gorm
- Cobra
- Viper
- Zap Log
- Mysql

## Requirements to Devel/Build

```bash
## git, golang 1.22, templ, air, upx, goreleaser
## https://github.com/upx/upx/releases
## https://github.com/goreleaser/goreleaser/releases
go install github.com/a-h/templ/cmd/templ@latest
go install github.com/cosmtrek/air@latest
```


## Build and Run

```bash
git clone https://github.com/jniltinho/postfixadmin.git
cd postfixadmin
make
make run or ./postfixadmin serve
```

## Generate Config file

```bash
## Generate file prod_config.toml
./postfixadmin init
./postfixadmin serve --config=prod_config.toml
```
