/* greet.go */
package main

import (
    "github.com/codegangsta/cli"
    "os"
    "os/exec"
    "syscall"
)

func main() {
    app := cli.NewApp()

    start := cli.Command{
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

    list := cli.Command{
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

    destroy := cli.Command{
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

    app.Commands = []cli.Command{
        start,
        list,
        destroy,
    }

    app.Version = "0.0.1"
    app.Name = "managevm"
    app.Usage = " A simple script to manage virtualbox vms in headless mode"
    app.Action = cli.ShowAppHelp

    app.Run(os.Args)
}
