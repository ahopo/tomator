package yjjy

import (
	"fmt"
	"log"
	"tomator/cmd/common/util"

	"github.com/ghodss/yaml"
	"github.com/urfave/cli/v2"
)

func model() Model {
	return Model{
		YAML: ".yml",
		JSON: ".json",
	}
}
func (m *Model) convertYJ(cCtx *cli.Context) error {
	//STRING INPUT
	str := cCtx.String("string")
	//check if str flag has value then process.
	if len(str) > 0 {
		y, _ := yaml.YAMLToJSON([]byte(str))
		fmt.Println(string(y))
		return nil
	}
	//FILE INPUT
	file := util.InitFile(cCtx.String("file"))
	//validate file exist and in yaml extension.
	if file.IsValid && file.Ext != m.YAML {
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
func (m *Model) convertJY(cCtx *cli.Context) error {
	//STRING INPUT
	str := cCtx.String("string")
	//check if str flag has value then process.
	if len(str) > 0 {
		y, _ := yaml.JSONToYAML([]byte(str))
		fmt.Println(string(y))
		return nil
	}
	//FILE INTPU
	file := util.InitFile(cCtx.String("file"))
	//validate file exist and in yaml extension.
	if file.IsValid && file.Ext != m.JSON {
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
