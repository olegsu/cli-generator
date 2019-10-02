package generate

import (
	"github.com/olegsu/cli-generator/pkg/logger"
	"github.com/spf13/viper"
)

type (
	// Handler implement Handler function
	Handler struct{}

	// Options additional options to Handle func
	Options struct {
		ResultRenderProcessor resultRenderProcessor
		Logger                logger.Logger
	}
)

// Handle - entry point to generate logic
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
