package passwordgenerator

import (
	"fmt"
	"log"
	"tomator/cmd/common/util"

	"github.com/sethvargo/go-password/password"
	"github.com/urfave/cli/v2"
)

func model() Model {
	return Model{}
}
func (pg Model) password(cCtx *cli.Context) error {
	n := cCtx.Int("length") / 2
	allowUpper := cCtx.Bool("allow-upper")
	allowRepeats := cCtx.Bool("allow-repeats")
	pass, err := password.Generate(cCtx.Int("length"), n, 0, !allowUpper, allowRepeats)
	if err != nil {
		log.Fatal(err)
	}
	pg.Allow_repeats = allowRepeats
	pg.Allow_upper = allowUpper
	pg.Length = cCtx.Int("length")
	pg.Pass = pass

	fmt.Println(util.SturctToYML(pg))
	return nil
}
