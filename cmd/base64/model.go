package base64

import "github.com/urfave/cli/v2"

type Model struct {
	Name    string
	Aliases []string
	Usage   string
	Action  func(cCtx *cli.Context) error
}
type B64 interface {
	SubCommand() *cli.Command
}
