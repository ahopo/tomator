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
		Commands: *commands(),
	}
	err := app.Run(os.Args)
	helper.CheckError(err, true)
}
func commands() *cli.Commands {
	coms := cli.Commands{}

	coms = append(coms, af.Command())
	coms = append(coms, pg.Command())

	comms := []helper.MNCM{cc.Command(), b64.Command()}
	for _, v := range comms {
		coms = append(coms, v.MainCommand())
	}

	return &coms
}
