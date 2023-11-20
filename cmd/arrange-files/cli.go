package arrangefiles

import "github.com/urfave/cli/v2"

func targetCom() (targetCom *cli.Command) {
	targetCom = &cli.Command{}
	targetCom.Name = "target-file"
	targetCom.Aliases = []string{"tf"}
	targetCom.Usage = "To arrange specific file extensions."
	targetCom.Action = model().target
	targetCom.Flags = CommonFlag
	return
}

func allfilesCom() (allfilescom *cli.Command) {
	allfilescom = &cli.Command{}
	allfilescom.Name = "all"
	allfilescom.Usage = "To arrange all files."
	allfilescom.Action = model().all
	allfilescom.Flags = CommonFlag
	return
}
func Command() *cli.Command {

	com := &cli.Command{
		Name:  "arrange",
		Usage: "To arrage your files.",
	}
	com.Subcommands = append(com.Subcommands, targetCom())
	com.Subcommands = append(com.Subcommands, allfilesCom())

	return com
}
