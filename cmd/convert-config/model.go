package yjjy

import (
	"github.com/urfave/cli/v2"
)

const (
	YAML = ".yml"
	JSON = ".json"
)

var commonFlag = []cli.Flag{
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
