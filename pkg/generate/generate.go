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
	Handler struct{}
)

func (g *Handler) Handle(cnf *viper.Viper) error {

	log := logger.New(nil)

	calculateSha := cnf.GetBool("calculateSha")
	projectDir := cnf.GetString("projectDir")
	lang := cnf.GetString("language")

	sha := shaCalculator{
		path:    fmt.Sprintf("%s/.cli-generator.sha", projectDir),
		content: bytes.NewBuffer(nil),
	}

	var err error
	var s *spec.CLISpec
	var specJSON map[string]interface{}

	if s, err = getCliSpec(cnf.GetString("spec"), ioutil.ReadFile); err != nil {
		return err
	}

	if specJSON, err = spec.ToJSON(s); err != nil {
		return err
	}

	log.Debug("Running template engine", "lang", lang)
	var res []*language.RenderResult
	var renderError error
	engine := language.New(&language.Options{
		Type:             lang,
		Spec:             s,
		Logger:           log.New("type", lang),
		ProjectDirectory: projectDir,
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
		if err := write(r.Content, file); err != nil {
			log.Error(renderError.Error())
		}
		if calculateSha {
			sha.append(r.Content)
		}
		log.Debug("File created", "name", r.File)
	}

	if calculateSha {
		data := sha.calc()
		file, err := os.Create(sha.path)
		if err != nil {
			log.Error("Failed to create .cli-generator.sha file")
			return err
		}
		if err = write(data, file); err != nil {
			log.Error("Failed to write to .cli-generator.sha file")
		}
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

func write(content *bytes.Buffer, writer io.Writer) error {
	_, err := fmt.Fprintf(writer, "%s", content.String())
	if err != nil {
		return err
	}
	return nil
}
