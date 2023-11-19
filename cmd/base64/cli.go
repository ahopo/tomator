package base64

import "github.com/urfave/cli/v2"

func decodeCom() (decode *cli.Command) {
	decode = &cli.Command{}
	decode.Name = "decode"
	decode.Aliases = []string{"D", "d"}
	decode.Usage = "Decode base64."
	decode.Action = model().decode
	return
}
func encodeCom() (encode *cli.Command) {
	encode = &cli.Command{}
	encode.Name = "encode"
	encode.Aliases = []string{"E", "d"}
	encode.Usage = "Encode base64."
	encode.Action = model().encode
	return
}
func Command() *cli.Command {

	com := &cli.Command{
		Name:    "base64",
		Aliases: []string{"b64"},
		Usage:   "To decode and encode.",
	}
	com.Subcommands = append(com.Subcommands, encodeCom())
	com.Subcommands = append(com.Subcommands, decodeCom())
	return com
}
