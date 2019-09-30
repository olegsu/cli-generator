package language

import (
	"github.com/olegsu/cli-generator/pkg/spec"
)

func toGolangType(t string) string {
	switch spec.Type(t) {
	case spec.String:
		return "string"
	case spec.Bool:
		return "bool"
	case spec.Number:
		return "float64"
	}
	return ""
}

func golangFlagDefaultFunc(t string) string {
	switch spec.Type(t) {
	case spec.String:
		return "GetString"
	case spec.Bool:
		return "GetBool"
	case spec.Number:
		return "GetFloat64"
	}
	return ""
}

func golangFlagFunc(t string) string {
	switch spec.Type(t) {
	case spec.String:
		return "StringVar"
	case spec.Bool:
		return "BoolVar"
	case spec.Number:
		return "Float64Var"
	}
	return ""
}
