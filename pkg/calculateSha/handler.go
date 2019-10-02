package calculateSha

import (
	"bytes"

	"github.com/olegsu/cli-generator/pkg/generate"
	"github.com/olegsu/cli-generator/pkg/logger"
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
	h := generate.Handler{}
	h.Handle(cnf, generate.Options{
		ResultRenderProcessor: &processor{
			content: bytes.NewBuffer(nil),
		},
		Logger: log,
	})
	return nil
}
