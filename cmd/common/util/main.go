package util

import (
	"tomator/cmd/common/helper"

	"gopkg.in/yaml.v2"
)

func SturctToYML(_struct interface{}) string {
	data, err := yaml.Marshal(_struct)
	helper.CheckError(err, false)
	return string(data)
}
