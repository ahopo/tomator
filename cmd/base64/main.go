package base64

import (
	"encoding/base64"
	"fmt"

	"github.com/urfave/cli/v2"
)

func model() Model {
	return Model{}
}
func (m Model) decode(cCtx *cli.Context) error {

	str := cCtx.Args().Get(0)
	decodedStr, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		panic("malformed input")
	}
	fmt.Println(string(decodedStr))
	return nil
}
func (m Model) encode(cCtx *cli.Context) error {
	str := cCtx.Args().Get(0)
	encodedStr := base64.StdEncoding.EncodeToString([]byte(str))
	fmt.Println(encodedStr)
	return nil
}
