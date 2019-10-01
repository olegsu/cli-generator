package generate

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"

	"github.com/olegsu/cli-generator/pkg/generate/language"
	"github.com/olegsu/cli-generator/pkg/logger"
	"github.com/olegsu/cli-generator/pkg/spec"

	"github.com/ghodss/yaml"
	"github.com/spf13/viper"
)

type (
	Handler struct{}

	Options struct {
		ResultRenderProcessor resultRenderProcessor
		Logger                logger.Logger
	}

	resultRenderProcessor interface {
		Process([]*language.RenderResult) error
	}
)

func (g *Handler) Handle(cnf *viper.Viper, opt ...Options) error {
	var log logger.Logger = logger.New(&logger.Options{
		Verbose: cnf.GetBool("verbose"),
	})
	var rrp resultRenderProcessor = &processor{log}
	if len(opt) == 1 {
		if opt[0].ResultRenderProcessor != nil {
			rrp = opt[0].ResultRenderProcessor
		}
		if opt[0].Logger != nil {
			log = opt[0].Logger
		}
	}
	return handle(cnf, log, rrp)
}

func handle(cnf *viper.Viper, log logger.Logger, processor resultRenderProcessor) error {

	projectDir := cnf.GetString("projectDir")
	lang := cnf.GetString("language")

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
	var renderError error
	engine := language.New(&language.Options{
		Type:             lang,
		Spec:             s,
		Logger:           log.New("type", lang),
		ProjectDirectory: projectDir,
		GenerateHandlers: cnf.GetBool("createHandlers"),
	})
	key, store := engine.BuildData(cnf)

	var res []*language.RenderResult
	if res, renderError = engine.Render(map[string]interface{}{
		"spec": specJSON,
		key:    store,
	}); renderError != nil {
		log.Error(renderError.Error())
	}

	return processor.Process(res)

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
