# Minion

Minion runs commands on multiple directories

## Install

Option 1) Using gobinaries

```shell
$ curl -sf https://gobinaries.com/pablote/minion | sh
```

Option 2) Using `go install`. Make sure $GOPATH/bin is on your path.

```shell
$ go install github.com/pablote/minion@latest
```

Add a config file to your home:

```shell
$ vim ~/minion.yaml
```

```
someProject:
  - /path/to/repository
  - /path/to/another/repository

anotherProject:
  - /path/to/microservice/a
  - /path/to/microservice/b
  - /path/to/microservice/c
```

## Usage

Run `minion` to get a list of available commands and help.
