package arrangefiles

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

func model() Model {
	return Model{}
}

func (a Model) target(cCtx *cli.Context) error {

	for _, d := range cCtx.StringSlice("src") {
		fmt.Print(d)
	}
	return nil
}
func (a Model) all(cCtx *cli.Context) error {
	fmt.Println("the sources from", cCtx.String("src"))
	fmt.Println("the destination to", cCtx.String("des"))
	print("TXT Arrange!")
	return nil
}
