package generate

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path"

	"github.com/olegsu/cli-generator/pkg/generate/language"
	"github.com/olegsu/cli-generator/pkg/logger"
)

type (
	processor struct {
		log logger.Logger
	}
)

func (p *processor) Process(data []*language.RenderResult) error {
	var err error
	for _, r := range data {
		var file *os.File
		dirPath, filePath := path.Split(r.File)
		p.log.Debug("Creating file", "dir", dirPath, "file", filePath)
		if err = os.MkdirAll(dirPath, os.ModePerm); err != nil {
			return err
		}

		if file, err = os.Create(r.File); err != nil {
			return err
		}

		if err := write(r.Content, file); err != nil {
			return err
		}
		p.log.Debug("File created", "name", r.File)
	}
	return nil
}

func write(content *bytes.Buffer, writer io.Writer) error {
	_, err := fmt.Fprintf(writer, "%s", content.String())
	if err != nil {
		return err
	}
	return nil
}
