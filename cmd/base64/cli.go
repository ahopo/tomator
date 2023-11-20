package base64

import "github.com/urfave/cli/v2"

func (m *Model) SubCommand() (_cli *cli.Command) {

	_cli = &cli.Command{}
	_cli.Name = m.Name
	_cli.Aliases = m.Aliases
	_cli.Usage = m.Usage
	_cli.Action = m.Action
	return
}
func Command() *cli.Command {
	com := &cli.Command{
		Name:    "base64",
		Aliases: []string{"b64"},
		Usage:   "To decode and encode.",
	}
	decode := Model{Action: decode, Name: "decode", Aliases: []string{"d"}, Usage: "To decode given base64 input"}
	encode := Model{Action: encode, Name: "encode", Aliases: []string{"e"}, Usage: "To encode base64 given raw input"}
	b64 := []B64{&decode, &encode}
	for _, val := range b64 {
		com.Subcommands = append(com.Subcommands, val.SubCommand())
	}
	return com
}
