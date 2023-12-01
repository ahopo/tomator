package base64

import (
	"encoding/base64"
	"fmt"
	"tomator/cmd/common/helper"

	"github.com/urfave/cli/v2"
)

func (m *Model) encode() string {
	return base64.StdEncoding.EncodeToString([]byte(m.raw))
}
func (m *Model) decode() string {
	decodedStr, err := base64.StdEncoding.DecodeString(m.raw)
	helper.CheckError(err, true)
	return string(decodedStr)
}

func initModel(raw string) *Model {
	m := &Model{
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

type Model struct {
	Name    string
	Usage   string
	Aliases []string
	d       string
	e       string
	raw     string
}

func (m *Model) subCommands() []*cli.Command {
	m = initModel("")
	return []*cli.Command{{
		Action: action, Name: m.d, Usage: "To decode given base64 input",
	}, {Action: action, Name: m.e, Usage: "To encode base64 given raw input"}}
}
func (m Model) Commands() *cli.Command {
	return &cli.Command{
		Subcommands: m.subCommands(),
		Name:        m.Name,
		Usage:       m.Usage,
		Aliases:     m.Aliases,
	}
}
