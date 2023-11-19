package passwordgenerator

import (
	"fmt"
	"log"

	"github.com/sethvargo/go-password/password"
	"github.com/urfave/cli/v2"
)

func generate() Generate {
	return Generate{}
}
func (pg Generate) password(cCtx *cli.Context) error {
	n := cCtx.Int("length") / 2
	res, err := password.Generate(cCtx.Int("length"), n, 0, cCtx.Bool("allow-upper"), cCtx.Bool("allow-repeats"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
	return nil
}
