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
   --docker-entry value, -d value  docker api entry, default unix://var/run/docker.sock
   --help, -h                      show help
   --version, -v                   print the version

```

Remove images with repository name `mybox` except latest(by created time) 5

```
$ ./cleanup rm --name mybox --count 5 
```

Remove images with repository name `mybox` which are created 5 days before, if all images are created 5 days ago, at least one image will be reserved

```
$ ./cleanup rm --name mybox --days 5
```

### install

```bash
go install github.com/lethe3000/cleanup
```
