package validate

import (
	"io/ioutil"

	"github.com/olegsu/cli-generator/configs/templates"
	"github.com/olegsu/cli-generator/pkg/logger"
	"github.com/olegsu/cli-generator/pkg/spec"
	"github.com/spf13/viper"
)

type (
	// Handler - exposed struct that implementd Handler interface
	Handler struct{}
)

// Handle - the function that will be called from the CLI with viper config
// to provide access to all flags
func (g *Handler) Handle(cnf *viper.Viper) error {
	log := logger.New(&logger.Options{
		Verbose: cnf.GetBool("verbose"),
	})
	var s *spec.CLISpec
	var err error

	log.Debug("Validating spec", "path", cnf.GetStringSlice("spec")[0])
	if s, err = spec.GetCliSpec(cnf.GetStringSlice("spec")[0], ioutil.ReadFile); err != nil {
		return err
	}

	if err = s.Validate([]byte(templates.TemplatesMap()["spec.json"])); err != nil {
		return err
	}
	return nil
}
