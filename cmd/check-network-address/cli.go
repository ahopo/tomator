package checknetworkaddress

import "github.com/urfave/cli/v2"

func (m *Model) SubCommand() (_cli *cli.Command) {

	return
}
func Command() *cli.Command {
	com := &cli.Command{
		Name:    "base64",
		Aliases: []string{"b64"},
		Usage:   "To decode and encode.",
	}

	return com
}
