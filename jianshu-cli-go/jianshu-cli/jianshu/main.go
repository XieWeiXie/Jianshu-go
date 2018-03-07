package main

import (
	"os"

	"fmt"
	"jianshu-go/jianshu-cli-go/jianshu-cli/bintray"

	"jianshu-go/jianshu-cli-go/jianshu-cli/utils"

	"github.com/urfave/cli"
)

const commandHelpTemplate = `
{{.HelpName}}{{if .UsageText}}
Arguments:
{{.UsageText}}
{{end}}{{if .Flags}}
Options:
	{{range .Flags}}{{.}}
	{{end}}{{end}}{{if .ArgsUsage}}
Environment Variables:
{{.ArgsUsage}}{{end}}`

const appHelpTemplate = `
NAME:
   {{.Name}} - {{.Usage}}

USAGE:
   {{if .UsageText}}{{.UsageText}}{{else}}{{.HelpName}} {{if .Flags}}[global options]{{end}}{{if .Commands}} command [command options]{{end}} [arguments...]{{end}}
   {{if .Version}}
VERSION:
   {{.Version}}
   {{end}}{{if len .Authors}}
AUTHOR(S):
   {{range .Authors}}{{ . }}{{end}}
   {{end}}{{if .Commands}}
COMMANDS:
   {{range .Commands}}{{join .Names ", "}}{{ "\t" }}{{.Usage}}
   {{end}}{{end}}{{if .Flags}}
GLOBAL OPTIONS:
   {{range .Flags}}{{.}}
   {{end}}`

const subCommandHelpTemplate = `
NAME:
   {{.HelpName}} - {{.Usage}}

USAGE:
   {{.HelpName}} command{{if .Flags}} [command options]{{end}}[arguments...]

COMMANDS:
   {{range .Commands}}{{join .Names ", "}}{{ "\t" }}{{.Usage}}
   {{end}}{{if .Flags}}
OPTIONS:
   {{range .Flags}}{{.}}
   {{end}}`

var author cli.Author

func init() {
	author = cli.Author{
		Name:  "xieWei",
		Email: "wuxiaoxiaoshen@shu.edu.cn",
	}
}

func main() {
	app := cli.NewApp()
	app.Name = utils.AppName
	app.Usage = utils.AppUsage
	app.Version = utils.AppVersion
	app.Authors = []cli.Author{author}
	app.CommandNotFound = getCommandNotFound
	app.Commands = getCommands()
	//cli.CommandHelpTemplate = commandHelpTemplate
	//cli.AppHelpTemplate = appHelpTemplate
	//cli.SubcommandHelpTemplate = subCommandHelpTemplate
	app.Action = func(c *cli.Context) {
		fmt.Println(app.Version)
	}
	app.Run(os.Args)

}

func getCommands() []cli.Command {
	return bintray.GetCommands()
}

func getCommandNotFound(cli *cli.Context, command string) {
	fmt.Printf("[WARNING] Command [ %s ] Not Found in %s .\n", command, utils.AppName)
	fmt.Printf("[MESSAGE] Please Type : %s --help .\n", utils.AppName)
}
