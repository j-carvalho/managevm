package main

import (
    "github.com/codegangsta/cli"
    "os"
)

func main() {
    app := cli.NewApp()

    app.Commands = []cli.Command{
        List,
        Start,
        Destroy,
    }

    app.Version = "0.0.1"
    app.Name = "managevm"
    app.Usage = " A simple script to manage virtualbox vms in headless mode"
    app.Action = cli.ShowAppHelp

    app.Run(os.Args)
}
