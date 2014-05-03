package main

import (
    "github.com/codegangsta/cli"
    "os"
    "os/exec"
    "syscall"
)

var Start = cli.Command{
    Name:        "start",
    ShortName:   "s",
    Usage:       "Initializes the desired vm in headless mode",
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

        args := []string{"VBoxManage", "startvm", vm_name, "-type", "headless"}

        env := os.Environ()

        err = syscall.Exec(binary, args, env)

        if err != nil {
            panic(err)
        }
    },
}
