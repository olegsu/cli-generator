package spec

import (
	"encoding/json"
	"fmt"

	"github.com/ghodss/yaml"
	"github.com/qri-io/jsonschema"
)

func ToJSON(in interface{}) (map[string]interface{}, error) {
	res := make(map[string]interface{})
	b, err := json.Marshal(in)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(b, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (cli *CLISpec) Validate(schema []byte) error {

	rs := &jsonschema.RootSchema{}
	if err := json.Unmarshal(schema, rs); err != nil {
		return err
	}
	b, err := cli.Marshal()
	if err != nil {
		return err
	}
	res, err := rs.ValidateBytes(b)
	if err != nil {
		return err
	}
	message := ""
	for _, e := range res {
		message = fmt.Sprintf("%s\n%s", message, e.Error())
	}
	if message != "" {
		return fmt.Errorf(message)
	}
	return nil
}

func GetCliSpec(path string, readFromFile func(path string) ([]byte, error)) (*CLISpec, error) {
	var err error
	var specData []byte
	var spec = CLISpec{}
	if specData, err = readFromFile(path); err != nil {
		return nil, err
	}
	if err = yaml.Unmarshal(specData, &spec); err != nil {
		return nil, err
	}
	return &spec, nil
}
