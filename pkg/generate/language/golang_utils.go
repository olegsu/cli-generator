package language

import (
	"strings"
	"text/template"

	"github.com/hairyhenderson/gomplate"
	"github.com/iancoleman/strcase"
	"github.com/olegsu/cli-generator/pkg/spec"
)

func toGolangType(t spec.Type) string {
	switch t {
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

func golangFlagDefaultFunc(t spec.Type) string {
	switch t {
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

func golangFlagFunc(t spec.Type) string {
	switch t {
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

func golangRulesToArgsValidation(rule spec.Rule) string {
	res := ""

	switch rule {
	case spec.Any:
		return "cobra.ArbitraryArgs"
	case spec.AtLeastOne:
		return "cobra.MinimumNArgs(1)"
	case spec.AtLeastTwo:
		return "cobra.MinimumNArgs(2)"
	case spec.ExactlyOne:
		return "cobra.ExactArgs(1)"
	}
	return res
}

func buildFullCmdName(cmd spec.Command) string {
	if cmd.Root {
		return cmd.Name
	}
	return strcase.ToLowerCamel(strings.Join(append(cmd.Parents, cmd.Name), "_"))
}

func aggregateCmdFullName(names ...string) string {
	if len(names) == 1 {
		return names[0]
	}
	return strcase.ToLowerCamel(strings.Join(names, "_"))
}

func getCommonTemplateFuncs() template.FuncMap {
	funcs := gomplate.Funcs(nil)
	funcs["toGolangType"] = toGolangType
	funcs["golangFlagFunc"] = golangFlagFunc
	funcs["golangFlagDefaultFunc"] = golangFlagDefaultFunc
	funcs["golangRulesToArgsValidation"] = golangRulesToArgsValidation
	funcs["buildFullCmdName"] = buildFullCmdName
	return funcs
}
