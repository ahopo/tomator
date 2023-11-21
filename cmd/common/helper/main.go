package helper

import (
	"log"

	"github.com/urfave/cli/v2"
)

func CheckError(e error, isfatal bool) {
	if isfatal && e != nil {
		log.Fatal(e)
	} else if !isfatal && e != nil {
		log.Println(e)
	}
}
func (c *SubCom) SubCommand() (_cli *cli.Command) {
	_cli = &cli.Command{}
	_cli.Name = c.Name
	_cli.Aliases = c.Aliases
	_cli.Usage = c.Usage
	_cli.Action = c.Action
	_cli.Flags = c.Flag
	return
}
func (m MainCom) MainCommand() *cli.Command {

	com := &cli.Command{
		Name:    m.Name,
		Aliases: m.Aliases,
		Usage:   m.Usage,
	}

	for _, val := range m.SubCommands {
		com.Subcommands = append(com.Subcommands, val.SubCommand())
	}
	return com
}
