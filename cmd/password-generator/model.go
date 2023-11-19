package passwordgenerator

import "github.com/urfave/cli/v2"

type Generate struct {
}

var CommonFlag = []cli.Flag{
	&cli.IntFlag{
		Name:    "length",
		Value:   8,
		Aliases: []string{"L"},
		Usage:   "Password length",
	},
	&cli.BoolFlag{
		Name:  "allow-upper",
		Usage: "Allow Uppercase.",
	},
	&cli.BoolFlag{
		Name:  "allow-repeats",
		Usage: "Allow repeats of character.",
	},
}
