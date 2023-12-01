package yjjy

import (
	"fmt"
	"log"
	file "tomator/cmd/common/util/file"

	"github.com/ghodss/yaml"
	"github.com/urfave/cli/v2"
)

func convertYJ(cCtx *cli.Context) error {
	//STRING INPUT
	str := cCtx.String("string")
	//check if str flag has value then process.
	if len(str) > 0 {
		y, _ := yaml.YAMLToJSON([]byte(str))
		fmt.Println(string(y))
		return nil
	}
	//FILE INPUT
	file := file.Init(cCtx.String("file"))
	//validate file exist and in yaml extension.
	if file.IsValid && file.Ext != YAML {
		log.Fatal("Not a yaml file.")
	}

	//covert file content and create new file.
	y, _ := yaml.YAMLToJSON([]byte(file.Content))
	newFile := file.CreateUsingExistingName(file.ExtType.JSON)
	if file.Write(string(y)) {
		log.Println("Done: ", newFile)
	}
	return nil
}
func convertJY(cCtx *cli.Context) error {
	//STRING INPUT
	str := cCtx.String("string")
	//check if str flag has value then process.
	if len(str) > 0 {
		y, _ := yaml.JSONToYAML([]byte(str))
		fmt.Println(string(y))
		return nil
	}
	//FILE INTPU
	file := file.Init(cCtx.String("file"))
	//validate file exist and in yaml extension.
	if file.IsValid && file.Ext != JSON {
		log.Fatal("Not a yaml file.")
	}

	//covert file content and create new file.
	y, _ := yaml.JSONToYAML([]byte(file.Content))
	newFile := file.CreateUsingExistingName(file.ExtType.JSON)
	if file.Write(string(y)) {
		log.Println("Done: ", newFile)
	}

	return nil
}
func (m Model) subCommands() []*cli.Command {
	return []*cli.Command{
		{Action: convertJY, Name: "jsontoyaml", Aliases: []string{"jy"}, Usage: "Convert fom JSON to YAML", Flags: commonFlags},
		{Action: convertYJ, Name: "yamltojson", Aliases: []string{"yj"}, Usage: "Convert fom YAML to JSON", Flags: commonFlags},
	}
}
func (m Model) Commands() *cli.Command {
	return &cli.Command{
		Subcommands: m.subCommands(),
		Name:        m.Name,
		Usage:       m.Usage,
		Aliases:     m.Aliases,
	}
}
