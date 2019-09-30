package spec

import "encoding/json"

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
