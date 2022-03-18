# Minion

Minion runs commands on multiple directories

## Install

```shell
$ curl -sf https://gobinaries.com/pablote/minion | sh
```

If `gobinaries` doesn't work on your system, try with `goblin`:

```shell
$ curl -sf https://goblin.reaper.im/github.com/pablote/minion | sh
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
