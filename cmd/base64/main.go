package base64

import (
	"encoding/base64"
	"fmt"
	"tomator/cmd/common/helper"

	"github.com/urfave/cli/v2"
)

const (
	Decode = "decode"
	Encode = "encode"
)

func theaction(cCtx *cli.Context) error {
	str := cCtx.Args().Get(0)
	switch cCtx.Command.Name {
	case Decode:
		decodedStr, err := base64.StdEncoding.DecodeString(str)
		helper.CheckError(err, true)
		fmt.Println(string(decodedStr))
	case Encode:
		encodedStr := base64.StdEncoding.EncodeToString([]byte(str))
		fmt.Println(encodedStr)
	}
	return nil
}

func b64() []helper.SBCM {
	decode := helper.SubCom{Action: theaction, Name: Decode, Aliases: []string{"d"}, Usage: "To decode given base64 input"}
	encode := helper.SubCom{Action: theaction, Name: Encode, Aliases: []string{"e"}, Usage: "To encode base64 given raw input"}
	return []helper.SBCM{&decode, &encode}
}

func Command() helper.MainCom {
	return helper.MainCom{SubCommands: b64, Name: "base64", Aliases: []string{"b64"}, Usage: "Convert fromt YAML to JSON or JSON to YAML"}
}
