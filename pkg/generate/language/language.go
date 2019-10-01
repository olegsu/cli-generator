package language

import (
	"bytes"

	"github.com/olegsu/cli-generator/pkg/logger"
	"github.com/olegsu/cli-generator/pkg/spec"
	"github.com/spf13/viper"
)

type (
	Engine interface {
		Render(data interface{}) ([]*RenderResult, error)
		BuildData(cnf *viper.Viper) (string, map[string]interface{})
	}

	Options struct {
		Type             string
		Logger           logger.Logger
		ProjectDirectory string
		GenerateHandlers bool
		RunInitFlow      bool
		Spec             *spec.CLISpec
	}

	RenderResult struct {
		Content *bytes.Buffer
		File    string
	}
)

func New(opt *Options) Engine {
	if opt.Type == "go" {
		return &golang{
			logger:           opt.Logger,
			projectDirectory: opt.ProjectDirectory,
			spec:             opt.Spec,
			generateHandlers: opt.GenerateHandlers,
			runInitFlow:      opt.RunInitFlow,
		}
	}
	return nil
}
