package main

import (
    "github.com/codegangsta/cli"
    "os"
    "os/exec"
    "syscall"
)

var List = cli.Command{
    Name:        "list",
    ShortName:   "l",
    Usage:       "Display currently running or existing vms",
    Description: "Shows detail about currently running or existing vms",
    Subcommands: []cli.Command{
        {
            Name:  "available",
            Usage: "Shows available vms",
            Action: func(c *cli.Context) {
                binary, err := exec.LookPath("VBoxManage")

                if err != nil {
                    panic(err)
                }

                args := []string{"VBoxManage", "list", "vms"}

                env := os.Environ()

                err = syscall.Exec(binary, args, env)

                if err != nil {
                    panic(err)
                }
            },
        },
        {
            Name:  "running",
            Usage: "Shows currently running vms",
            Action: func(c *cli.Context) {
                binary, err := exec.LookPath("VBoxManage")

                if err != nil {
                    panic(err)
                }

                args := []string{"VBoxManage", "list", "runningvms"}

                env := os.Environ()

                err = syscall.Exec(binary, args, env)

                if err != nil {
                    panic(err)
                }
            },
        },
    },
    Action: cli.ShowSubcommandHelp,
}
