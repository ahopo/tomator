package base64

import (
	"encoding/base64"
	"fmt"
	"tomator/cmd/common/helper"

	"github.com/urfave/cli/v2"
)

func decode(cCtx *cli.Context) error {

	str := cCtx.Args().Get(0)
	decodedStr, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		panic("malformed input")
	}
	fmt.Println(string(decodedStr))
	return nil
}
func encode(cCtx *cli.Context) error {
	str := cCtx.Args().Get(0)
	encodedStr := base64.StdEncoding.EncodeToString([]byte(str))
	fmt.Println(encodedStr)
	return nil
}

func Command() *cli.Command {
	com := &cli.Command{
		Name:    "base64",
		Aliases: []string{"b64"},
		Usage:   "To decode and encode.",
	}
	decode := helper.SubCom{Action: decode, Name: "decode", Aliases: []string{"d"}, Usage: "To decode given base64 input"}
	encode := helper.SubCom{Action: encode, Name: "encode", Aliases: []string{"e"}, Usage: "To encode base64 given raw input"}
	b64 := []helper.SBCM{&decode, &encode}
	for _, val := range b64 {
		com.Subcommands = append(com.Subcommands, val.SubCommand())
	}
	return com
}
