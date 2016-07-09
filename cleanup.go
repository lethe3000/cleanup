package main

import (
  "fmt"
  "os"
  "github.com/urfave/cli"
  "testbird.com/docker/command"
)

func main() {
  start()
}

func start() {
    app := cli.NewApp()
    app.Name = "docker image clean tool"
    app.Version = "0.0.0"
    cli.VersionPrinter = func(c *cli.Context) {
        fmt.Fprintf(c.App.Writer, "cleanup version: %v\n", c.App.Version)
    }
    app.Usage = "A simple command line tool for docker image clean."
    app.Flags = []cli.Flag{
        cli.StringFlag{Name: "docker-entry, d", Usage: "docker api entry"},
    }
    app.Commands = []cli.Command{
        command.NewDummyCommand(),
        command.NewImageListCommand(),
        command.NewImageInfoCommand(),
        command.NewRemoveCommand(),
    }

    err := app.Run(os.Args)
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
    }
}