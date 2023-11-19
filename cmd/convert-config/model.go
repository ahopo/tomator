package yjjy

import (
	"github.com/urfave/cli/v2"
)

type Model struct {
	YAML string
	JSON string
}

var CommonFlag = []cli.Flag{
	&cli.StringFlag{
		Name:  "file",
		Usage: "Will convert file",
	},
	&cli.StringFlag{
		Name:    "string",
		Aliases: []string{"str", "text"},
		Usage:   "Will convert string/text",
	},
}
