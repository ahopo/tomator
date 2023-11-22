package arrangefiles

import (
	"fmt"
	"log"
	"tomator/cmd/common/helper"

	"github.com/urfave/cli/v2"
)

func model() Model {
	return Model{}
}

func target(cCtx *cli.Context) error {

	log.Println(cCtx.String("src"))
	return nil
}
func all(cCtx *cli.Context) error {
	fmt.Println("the sources from", cCtx.String("src"))
	fmt.Println("the destination to", cCtx.String("des"))
	print("TXT Arrange!")
	return nil
}

func af() []helper.SBCM {
	decode := helper.SubCom{Action: target, Name: "target", Usage: "Target specific extension"}
	encode := helper.SubCom{Action: all, Name: "all", Usage: "All files"}
	return []helper.SBCM{&decode, &encode}
}

func Subcommand() helper.MainCom {
	return helper.MainCom{SubCommands: af, Name: "arrange-file", Aliases: []string{"af"}, Usage: "Arrange files"}
}
