package main

import (
    "github.com/codegangsta/cli"
    "os"
    "os/exec"
    "syscall"
)

var Destroy = cli.Command{
    Name:        "destroy",
    ShortName:   "d",
    Usage:       "Kills a currently running vm",
    Description: "Expects the vm name as the argument.",
    Action: func(c *cli.Context) {

        vm_name := c.Args().First()

        if vm_name == "" {
            println("Please specify vm name")
            return
        }

        binary, err := exec.LookPath("VBoxManage")

        if err != nil {
            panic(err)
        }

        args := []string{"VBoxManage", "controlvm", vm_name, "poweroff"}

        env := os.Environ()

        err = syscall.Exec(binary, args, env)

        if err != nil {
            panic(err)
        }
    },
}
