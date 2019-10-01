package calculateSha

import (
	"bytes"
	"crypto/sha256"
	"fmt"

	"github.com/olegsu/cli-generator/pkg/generate"
	"github.com/olegsu/cli-generator/pkg/generate/language"
	"github.com/olegsu/cli-generator/pkg/logger"
	"github.com/spf13/viper"
)

type (
	// Handler - exposed struct that implementd Handler interface
	Handler struct{}

	processor struct {
		content *bytes.Buffer
		log     logger.Logger
	}
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

func (p *processor) Process(data []*language.RenderResult) error {
	for _, r := range data {
		p.append(r.Content)
	}
	fmt.Println(p.calc())
	return nil
}

func (s *processor) append(content *bytes.Buffer) {
	s.content.Write(content.Bytes())
}

func (s *processor) calc() *bytes.Buffer {
	sum := sha256.Sum256(s.content.Bytes())
	return bytes.NewBufferString(fmt.Sprintf("%x", sum))
}
