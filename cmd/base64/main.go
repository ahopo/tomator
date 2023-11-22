package base64

import (
	"encoding/base64"
	"fmt"
	"tomator/cmd/common/helper"

	"github.com/urfave/cli/v2"
)

type model struct {
	d   string
	e   string
	raw string
}

func (m *model) encode() string {
	return base64.StdEncoding.EncodeToString([]byte(m.raw))
}
func (m *model) decode() string {
	decodedStr, err := base64.StdEncoding.DecodeString(m.raw)
	helper.CheckError(err, true)
	return string(decodedStr)
}

func initModel(raw string) *model {
	m := &model{
		d:   "decode",
		e:   "encode",
		raw: raw,
	}
	return m
}

func action(cCtx *cli.Context) error {
	m := initModel(cCtx.Args().Get(0))
	switch cCtx.Command.Name {
	case m.d:
		fmt.Println(m.decode())
	case m.e:
		fmt.Println(m.encode())
	}
	return nil
}

func b64() []helper.SBCM {
	m := initModel("")
	decode := helper.SubCom{Action: action, Name: m.d, Aliases: []string{"d"}, Usage: "To decode given base64 input"}
	encode := helper.SubCom{Action: action, Name: m.e, Aliases: []string{"e"}, Usage: "To encode base64 given raw input"}
	return []helper.SBCM{&decode, &encode}
}

func Subcommand() helper.MainCom {
	return helper.MainCom{SubCommands: b64, Name: "base64", Aliases: []string{"b64"}, Usage: "Convert fromt YAML to JSON or JSON to YAML"}
}
