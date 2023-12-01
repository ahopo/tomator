package handlers

import (
	"os"
	"tomator/cmd/base64"
	"tomator/cmd/common/helper"
	cc "tomator/cmd/convert-config"

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
	comms := []helper.CLI{
		cc.Model{
			Name:    "convert-config",
			Usage:   "Convert YAML to JSON or JSON to YAML",
			Aliases: []string{"cc"},
		},
		base64.Model{
			Name:    "base64",
			Usage:   "Encode and Decode",
			Aliases: []string{"b64"},
		}}
	for _, c := range comms {
		scom = append(scom, c.Commands())
	}
	return
}
