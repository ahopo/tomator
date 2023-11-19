package passwordgenerator

import "github.com/urfave/cli/v2"

func Command() (passCom *cli.Command) {
	passCom = &cli.Command{}
	passCom.Name = "generate-password"
	passCom.Usage = "To generate password."
	passCom.Aliases = []string{"gp"}
	passCom.Action = model().password
	passCom.Flags = CommonFlag
	return
}
