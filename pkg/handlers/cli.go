package handlers

import (
	"log"
	"os"
	af "tomator/cmd/arrange-files"
	b64 "tomator/cmd/base64"
	cc "tomator/cmd/convert-config"
	pg "tomator/cmd/password-generator"

	"github.com/urfave/cli/v2"
)

func Init() {

	coms := cli.Commands{}
	coms = append(coms, af.Command())
	coms = append(coms, pg.Command())
	coms = append(coms, b64.Command())
	coms = append(coms, cc.Command())

	app := &cli.App{
		Name:     "tomator",
		Usage:    "automate your daily repetitive task!",
		Authors:  []*cli.Author{{Name: "aopo", Email: "alfieopo@gmail.com"}},
		Commands: coms,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
