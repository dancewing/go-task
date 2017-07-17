package main

import (
	"os"

	"github.com/urfave/cli"

	"github.com/dancewing/go-task/cmd"
	"github.com/dancewing/go-task/pkg/setting"
)

const APP_VER = "0.11.24.0627"

func init() {
	setting.AppVer = APP_VER
}

func main() {
	app := cli.NewApp()
	app.Name = "Gogs"
	app.Usage = "A painless self-hosted Git service"
	app.Version = APP_VER
	app.Commands = []cli.Command{
		cmd.Web,
	}
	app.Flags = append(app.Flags, []cli.Flag{}...)
	app.Run(os.Args)
}
