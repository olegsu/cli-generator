package language

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/iancoleman/strcase"
	"github.com/imdario/mergo"
	"github.com/olegsu/cli-generator/configs/templates"
	"github.com/olegsu/cli-generator/pkg/logger"
	"github.com/olegsu/cli-generator/pkg/spec"
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
	var rootLoose *bool
	if g.spec.Loose != nil {
		rootLoose = g.spec.Loose
	} else {
		l := true
		rootLoose = &l
	}
	r := "root"
	rootCmd := spec.Command{
		Flags:    g.spec.Flags,
		Name:     r,
		Parent:   &r,
		Root:     true,
		Loose:    rootLoose,
		Commands: []spec.Command{},
	}
	for _, cmd := range g.spec.Commands {
		rootCmd.Commands = append(rootCmd.Commands, cmd)
	}

	rootJSON, err := spec.ToJSON(rootCmd)
	if err != nil {
		return nil, err
	}
	var result []*RenderResult
	{
		result = append(result, g.renderFile(fmt.Sprintf("%s/main.go", g.projectDirectory), tmap[templateMain], data))

		rootData := map[string]interface{}{
			"cmd": rootJSON,
		}
		mergo.Merge(&rootData, data)

		result = append(result, g.renderFile(fmt.Sprintf("%s/cmd/root.go", g.projectDirectory), tmap[templateCmd], rootData))
		res, err := g.renderCommands(rootCmd, tmap, data)
		if err != nil {
			return nil, err
		}
		for _, r := range res {
			result = append(result, r)
		}

	}

	if g.runInitFlow {
		g.logger.Debug("Creating Makefile")
		result = append(result, g.renderFile(fmt.Sprintf("%s/Makefile", g.projectDirectory), tmap[templateMakefile], data))
	}

	return result, nil
}

func (g *golang) BuildData(cnf *viper.Viper) (string, map[string]interface{}) {
	res := map[string]interface{}{
		"package": cnf.GetString("goPackage"),
	}
	return "go", res
}

func (g *golang) renderFile(name string, tmpl string, data interface{}) *RenderResult {
	res := &RenderResult{
		File: name,
	}
	out := new(bytes.Buffer)
	template.Must(template.New("").Funcs(getCommonTemplateFuncs()).Parse(string(tmpl))).Execute(out, data)
	res.Content = out
	return res
}

func (g *golang) renderCommands(root spec.Command, templateMap map[string]string, data interface{}) ([]*RenderResult, error) {
	result := []*RenderResult{}
	for _, cmd := range root.Commands {
		name := strcase.ToLowerCamel(cmd.Name)
		cmd.Parent = &root.Name
		cmdJSON, err := spec.ToJSON(cmd)
		if err != nil {
			return nil, err
		}
		cmdData := map[string]interface{}{
			"cmd": cmdJSON,
		}
		mergo.Merge(&cmdData, data)
		result = append(result, g.renderFile(fmt.Sprintf("%s/cmd/%s.go", g.projectDirectory, name), templateMap[templateCmd], cmdData))
		if g.generateHandlers {
			result = append(result, g.renderFile(fmt.Sprintf("%s/pkg/%s/handler.go", g.projectDirectory, name), templateMap[templateHandler], cmdData))
		}
		res, err := g.renderCommands(cmd, templateMap, data)
		if err != nil {
			return nil, err
		}
		for _, r := range res {
			result = append(result, r)
		}

	}
	return result, nil
}
