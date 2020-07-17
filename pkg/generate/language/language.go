package language

import (
	"bytes"

	"github.com/olegsu/cli-generator/pkg/engine"
	"github.com/olegsu/cli-generator/pkg/logger"
	"github.com/olegsu/cli-generator/pkg/spec"
	"github.com/spf13/viper"
)

type (
	TemplateEngine interface {
		Render(data interface{}) ([]*RenderResult, error)
		BuildData(cnf *viper.Viper) (string, map[string]interface{})
		PreInitFlow() []engine.Task
		PostInitFlow() []engine.Task
	}

	Options struct {
		Type             string
		Logger           logger.Logger
		ProjectDirectory string
		GenerateHandlers bool
		RunInitFlow      bool
		Spec             *spec.CLISpec
		GoPackage        string
	}

	RenderResult struct {
		Content *bytes.Buffer
		File    string
	}

	TaskRunner interface {
		Run(*engine.RunOptions) ([]byte, error)
	}
)

// New returns new template engine based on the target language
func New(opt *Options) TemplateEngine {
	if opt.Type == "go" {
		return &golang{
			logger:           opt.Logger,
			projectDirectory: opt.ProjectDirectory,
			spec:             opt.Spec,
			generateHandlers: opt.GenerateHandlers,
			runInitFlow:      opt.RunInitFlow,
			goPackage:        opt.GoPackage,
		}
	}
	return nil
}
