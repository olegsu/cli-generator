package calculateSha

import (
	"bytes"
	"crypto/sha256"
	"fmt"

	"github.com/olegsu/cli-generator/pkg/generate/language"
	"github.com/olegsu/cli-generator/pkg/logger"
)

type (
	processor struct {
		content *bytes.Buffer
		log     logger.Logger
	}
)

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
