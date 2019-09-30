package generate

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"

	"github.com/olegsu/cli-generator/pkg/generate/language"
	"github.com/olegsu/cli-generator/pkg/logger"
	"github.com/olegsu/cli-generator/pkg/spec"

	"github.com/ghodss/yaml"
	"github.com/spf13/viper"
)

type (
	Generate struct{}
)

func (g *Generate) Handle(cnf *viper.Viper) error {

	log := logger.New(nil)

	var err error
	var s *spec.CLISpec
	var specJSON map[string]interface{}

	if s, err = getCliSpec(cnf.GetString("spec"), ioutil.ReadFile); err != nil {
		return err
	}

	if specJSON, err = spec.ToJSON(s); err != nil {
		return err
	}

	lang := cnf.GetString("language")

	log.Debug("Running engine", "lang", lang)
	// root cmd
	var res []*language.RenderResult
	var renderError error
	engine := language.New(&language.Options{
		Type:             lang,
		Spec:             s,
		Logger:           log.New("type", lang),
		ProjectDirectory: cnf.GetString("projectDir"),
		GenerateHandlers: cnf.GetBool("createHandlers"),
	})
	key, store := engine.BuildData(cnf)

	if res, renderError = engine.Render(map[string]interface{}{
		"spec": specJSON,
		key:    store,
	}); renderError != nil {
		log.Error(renderError.Error())
	}

	for _, r := range res {
		dirPath, filePath := path.Split(r.File)
		log.Debug("Creating file", "dir", dirPath, "file", filePath)
		err := os.MkdirAll(dirPath, os.ModePerm)
		if err != nil {
			log.Error(renderError.Error())
		}
		file, err := os.Create(r.File)
		if err != nil {
			log.Error(renderError.Error())
		}
		if err := writeFile(r.Content, file); err != nil {
			log.Error(renderError.Error())
		}
		log.Debug("File created", "name", r.File)
	}

	return nil
}

func getCliSpec(path string, readFromFile func(path string) ([]byte, error)) (*spec.CLISpec, error) {
	var err error
	var specData []byte
	var spec = spec.CLISpec{}
	if specData, err = readFromFile(path); err != nil {
		return nil, err
	}
	if err = yaml.Unmarshal(specData, &spec); err != nil {
		return nil, err
	}
	return &spec, nil
}

func writeFile(content *bytes.Buffer, writer io.Writer) error {
	_, err := fmt.Fprintf(writer, "%s", content.String())
	if err != nil {
		return err
	}
	return nil
}
