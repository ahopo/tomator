package arrangefiles

import "github.com/urfave/cli/v2"

type Model struct {
}

var CommonFlag = []cli.Flag{
	&cli.StringFlag{
		Name:     "src",
		Usage:    "Source folder",
		Required: true,
	},
	&cli.StringFlag{
		Name:  "des",
		Usage: "Destination folder",
	},
	&cli.StringFlag{
		Name:    "directory",
		Aliases: []string{"folder-name"},
	},
	&cli.StringSliceFlag{
		Name:    "extensions",
		Aliases: []string{"ex"},
		Usage:   "List of extension to arrange.",
	},
}
