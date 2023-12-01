package helper

import "github.com/urfave/cli/v2"

type CLI interface {
	Commands() *cli.Command
}
