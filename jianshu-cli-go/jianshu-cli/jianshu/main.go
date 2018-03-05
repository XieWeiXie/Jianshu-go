package main

import (
	"os"

	"fmt"

	"jianshu-go/jianshu-cli-go/jianshu-cli/bintray"

	"github.com/urfave/cli"
)

const commandHelpTemplate = ""

const appHelpTemplate = ""

const subCommandHelpTemplate = ""

func main() {
	app := cli.NewApp()
	app.Name = "JianShu"
	app.CommandNotFound = func(ctx *cli.Context, command string) {
		fmt.Printf("Command not found: %v\n", command)
	}
	app.Commands = getCommands()
	cli.CommandHelpTemplate = commandHelpTemplate
	cli.AppHelpTemplate = appHelpTemplate
	cli.SubcommandHelpTemplate = subCommandHelpTemplate
	app.Run(os.Args)
}

func getCommands() []cli.Command {
	return []cli.Command{
		{
			Name:        "bin",
			Subcommands: bintray.GetCommands(),
		},
	}
}
