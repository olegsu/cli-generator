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
	"github.com/olegsu/cli-generator/configs/templates"
	"github.com/spf13/viper"
)

type (
	golang struct {
		logger           logger.Logger
		projectDirectory string
		generateHandlers bool
		runInitFlow      bool
		spec             *spec.CLISpec
	}
)

const (
	templateMain     = "go.main.tmpl"
	templateCmd      = "go.cmd.tmpl"
	templateHandler  = "go.handler.tmpl"
	templateMakefile = "go.makefile.tmpl"
)

func (g *golang) Render(data interface{}) ([]*RenderResult, error) {
	g.logger.Debug("Renderring")
	tmap := templates.TemplatesMap()
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
		result = append(result, renderFile(fmt.Sprintf("%s/main.go", g.projectDirectory), tmap[templateMain], data))

		rootData := map[string]interface{}{
			"cmd": rootJSON,
		}
		mergo.Merge(&rootData, data)

		result = append(result, renderFile(fmt.Sprintf("%s/cmd/root.go", g.projectDirectory), tmap[templateCmd], rootData))
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
			result = append(result, renderFile(fmt.Sprintf("%s/cmd/%s.go", g.projectDirectory, name), tmap[templateCmd], cmdData))
			if g.generateHandlers {
				result = append(result, renderFile(fmt.Sprintf("%s/pkg/%s/%s.go", g.projectDirectory, name, name), tmap[templateHandler], cmdData))
			}
		}
	}

	if g.runInitFlow {
		g.logger.Debug("Creating Makefile")
		result = append(result, renderFile(fmt.Sprintf("%s/Makefile", g.projectDirectory), tmap[templateMakefile], data))
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
