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

	var start = cli.Command{
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

	app.Commands = []cli.Command{
		start,
	}

	app.Version = "0.0.1"
	app.Name = "managevm"
	app.Usage = " A simple script to manage virtualbox vms in headless mode"
	app.Action = cli.ShowAppHelp

	app.Run(os.Args)
}
