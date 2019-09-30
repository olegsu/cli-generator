package generate

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

type (
	shaCalculator struct {
		path    string
		content *bytes.Buffer
	}
)

func (s *shaCalculator) append(content *bytes.Buffer) {
	s.content.Write(content.Bytes())
}

func (s *shaCalculator) calc() *bytes.Buffer {
	sum := sha256.Sum256(s.content.Bytes())
	return bytes.NewBufferString(fmt.Sprintf("%x", sum))
}
