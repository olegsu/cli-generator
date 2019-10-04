package language

import (
	"github.com/olegsu/cli-generator/pkg/spec"
)

func toGolangType(t string) string {
	switch spec.Type(t) {
	case spec.String:
		return "string"
	case spec.ArrayString:
		return "[]string"
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
	case spec.ArrayString:
		return "GetStringSlice"
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
	case spec.ArrayString:
		return "StringArrayVar"
	case spec.Bool:
		return "BoolVar"
	case spec.Number:
		return "Float64Var"
	}
	return ""
}

func golangRulesToArgsValidation(rule string) string {
	res := ""

	switch spec.Rule(rule) {
	case spec.Any:
		return "cobra.ArbitraryArgs"
	case spec.AtLeastOne:
		return "cobra.MinimumNArgs(1)"
	case spec.AtLeastTwo:
		return "cobra.MinimumNArgs(2)"
	}

	return res
}
