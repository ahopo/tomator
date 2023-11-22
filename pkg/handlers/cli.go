package handlers

import (
	"os"
	af "tomator/cmd/arrange-files"
	b64 "tomator/cmd/base64"
	"tomator/cmd/common/helper"
	cc "tomator/cmd/convert-config"
	pg "tomator/cmd/password-generator"

	"github.com/urfave/cli/v2"
)

func Init() {
	app := &cli.App{
		Name:     "tomator",
		Usage:    "automate your daily repetitive task!",
		Authors:  []*cli.Author{{Name: "aopo", Email: "alfieopo@gmail.com"}},
		Commands: commands(),
	}
	err := app.Run(os.Args)
	helper.CheckError(err, true)
}

func commands() (scom []*cli.Command) {
	comms := []helper.MNCM{
		cc.Subcommand(),
		b64.Subcommand(),
		af.Subcommand(),
	}
	scom = append(scom, pg.Command())
	for _, val := range comms {
		scom = append(scom, val.MainCommand())
	}

	return
}
