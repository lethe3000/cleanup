# Clean images which out of date

### usage

```bash
$ ./cleanup -h
NAME:
   docker image clean tool - A simple command line tool for docker image clean.

USAGE:
   cleanup [global options] command [command options] [arguments...]

VERSION:
   0.0.0

COMMANDS:
     dummy    dummy command
     ls       list images
     info     list all tags for given image repo
     rm       remove images
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --docker-entry value, -d value  docker api entry
   --help, -h                      show help
   --version, -v                   print the version

```

### install

```bash
go install github.com/lethe3000/cleanup
```
