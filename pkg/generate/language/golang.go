package language

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/hairyhenderson/gomplate"
	"github.com/iancoleman/strcase"
	"github.com/imdario/mergo"
	"github.com/olegsu/cli-generator/pkg/logger"
	"github.com/olegsu/cli-generator/pkg/spec"
	"github.com/spf13/viper"
)

type (
	golang struct {
		logger           logger.Logger
		projectDirectory string
		generateHandlers bool
		spec             *spec.CLISpec
	}
)

const (
	templateMain    = "pkg/generate/language/templates/main.tmpl"
	templateCmd     = "pkg/generate/language/templates/cmd.tmpl"
	templateHandler = "pkg/generate/language/templates/handler.tmpl"
)

func (g *golang) Render(data interface{}) ([]*RenderResult, error) {
	g.logger.Debug("Renderring")

	rootFlag := spec.Command{
		Flags: g.spec.Flags,
		Name:  "root",
		Root:  true,
	}

	rootJSON, err := spec.ToJSON(rootFlag)
	if err != nil {
		return nil, err
	}
	var result []*RenderResult
	{
		result = append(result, renderFile(fmt.Sprintf("%s/main.go", g.projectDirectory), goMainTemplate, data))

		rootData := map[string]interface{}{
			"cmd": rootJSON,
		}
		mergo.Merge(&rootData, data)

		result = append(result, renderFile(fmt.Sprintf("%s/cmd/root.go", g.projectDirectory), goCmdTemplate, rootData))
		for _, cmd := range g.spec.Commands {
			name := strcase.ToLowerCamel(cmd.Name)
			parent := "root"
			cmd.Parent = &parent
			cmdJSON, err := spec.ToJSON(cmd)
			if err != nil {
				return nil, err
			}
			cmdData := map[string]interface{}{
				"cmd": cmdJSON,
			}
			mergo.Merge(&cmdData, data)
			result = append(result, renderFile(fmt.Sprintf("%s/cmd/%s.go", g.projectDirectory, name), goCmdTemplate, cmdData))
			if g.generateHandlers {
				result = append(result, renderFile(fmt.Sprintf("%s/pkg/%s/%s.go", g.projectDirectory, name, name), goHandlerTemplate, cmdData))
			}
		}

	}

	return result, nil
}

func (g *golang) BuildData(cnf *viper.Viper) (string, map[string]interface{}) {
	res := map[string]interface{}{
		"package": cnf.GetString("goPackage"),
	}
	return "go", res
}

func getCommonTemplateFuncs() template.FuncMap {
	funcs := gomplate.Funcs(nil)
	funcs["toGolangType"] = toGolangType
	funcs["golangFlagFunc"] = golangFlagFunc
	funcs["golangFlagDefaultFunc"] = golangFlagDefaultFunc
	return funcs
}

func renderFile(name string, tmpl string, data interface{}) *RenderResult {
	res := &RenderResult{
		File: name,
	}
	out := new(bytes.Buffer)
	template.Must(template.New("").Funcs(getCommonTemplateFuncs()).Parse(string(tmpl))).Execute(out, data)
	res.Content = out
	return res
}
