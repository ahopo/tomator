package yjjy

import "github.com/urfave/cli/v2"

func (m *Model) yamltojsonCom() (yamltojson *cli.Command) {
	yamltojson = &cli.Command{}
	yamltojson.Name = "yamltojson"
	yamltojson.Aliases = []string{"yj"}
	yamltojson.Usage = "Convert fom YAML to JSON"
	yamltojson.Action = m.convertYJ
	yamltojson.Flags = CommonFlag
	return
}
func (m *Model) jsontoyamlCom() (yamltojson *cli.Command) {
	yamltojson = &cli.Command{}
	yamltojson.Name = "jsontoyaml"
	yamltojson.Aliases = []string{"jy"}
	yamltojson.Usage = "Convert fom JSON to YAML"
	yamltojson.Action = m.convertJY
	yamltojson.Flags = CommonFlag
	return
}
func Command() *cli.Command {

	com := &cli.Command{
		Name:    "convert-config",
		Aliases: []string{"cc"},
		Usage:   "Convert fom YAML to JSON o JSON to YAML.",
	}
	yjjy := model()
	com.Subcommands = append(com.Subcommands, yjjy.yamltojsonCom())
	com.Subcommands = append(com.Subcommands, yjjy.jsontoyamlCom())
	return com
}
