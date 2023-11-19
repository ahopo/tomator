package passwordgenerator

import "github.com/urfave/cli/v2"

type Model struct {
	Allow_repeats bool   `json:"Allow Repeats"`
	Allow_upper   bool   `json:"Allow Upper"`
	Length        int    `json:"Length"`
	Pass          string `json:"Pass"`
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
		Value: false,
	},
	&cli.BoolFlag{
		Name:  "allow-repeats",
		Usage: "Allow repeats of character.",
		Value: false,
	},
}
