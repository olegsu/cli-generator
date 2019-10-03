package generate

import (
	"io/ioutil"

	"github.com/olegsu/cli-generator/configs/templates"
	"github.com/olegsu/cli-generator/pkg/generate/language"
	"github.com/olegsu/cli-generator/pkg/logger"
	"github.com/olegsu/cli-generator/pkg/spec"
	"github.com/spf13/viper"
)

type (
	resultRenderProcessor interface {
		Process([]*language.RenderResult) error
	}
)

func handle(cnf *viper.Viper, log logger.Logger, processor resultRenderProcessor) error {
	projectDir := cnf.GetString("projectDir")
	lang := cnf.GetString("language")

	var err error
	var s *spec.CLISpec
	var specJSON map[string]interface{}

	if s, err = spec.GetCliSpec(cnf.GetString("spec"), ioutil.ReadFile); err != nil {
		return err
	}

	if specJSON, err = spec.ToJSON(s); err != nil {
		return err
	}

	if err := s.Validate([]byte(templates.TemplatesMap()["spec.json"])); err != nil {
		return err
	}

	log.Debug("Running template engine", "lang", lang)
	var renderError error
	engine := language.New(&language.Options{
		Type:             lang,
		Spec:             s,
		Logger:           log.New("type", lang),
		ProjectDirectory: projectDir,
		RunInitFlow:      cnf.GetBool("runInitFlow"),
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
