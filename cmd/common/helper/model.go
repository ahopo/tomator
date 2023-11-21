package helper

import "github.com/urfave/cli/v2"

// subcommand
type SBCM interface {
	SubCommand() *cli.Command
}
type SubCom struct {
	Name    string
	Aliases []string
	Usage   string
	Action  func(cCtx *cli.Context) error
	Flag    []cli.Flag
}

// maincommand
type MNCM interface {
	MainCommand() *cli.Command
}

type MainCom struct {
	Name        string
	Aliases     []string
	Usage       string
	SubCommands func() []SBCM
}
