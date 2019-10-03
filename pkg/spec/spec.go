// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    cLISpec, err := UnmarshalCLISpec(bytes)
//    bytes, err = cLISpec.Marshal()

package spec

import "encoding/json"

func UnmarshalCLISpec(data []byte) (CLISpec, error) {
	var r CLISpec
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *CLISpec) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type CLISpec struct {
	Commands []Command `json:"commands"`
	Flags    []Flag    `json:"flags,omitempty"`
	Loose    *bool     `json:"loose,omitempty"`
	Metadata Metadata  `json:"metadata"`
}

type Command struct {
	Flags  []Flag  `json:"flags,omitempty"`
	Loose  *bool   `json:"loose,omitempty"`
	Name   string  `json:"name"`
	Parent *string `json:"parent,omitempty"`
	Root   bool    `json:"root"`
}

type Flag struct {
	Alias       []string `json:"alias,omitempty"`
	Default     *string  `json:"default,omitempty"`
	Description *string  `json:"description,omitempty"`
	Enum        []string `json:"enum,omitempty"`
	EnvVar      *string  `json:"envVar,omitempty"`
	Name        string   `json:"name"`
	Required    bool     `json:"required"`
	Type        Type     `json:"type"`
}

type Metadata struct {
	Description *string `json:"description,omitempty"`
	Name        string  `json:"name"` // CLI Name
	Version     string  `json:"version"`
}

type Type string

const (
	ArrayBool   Type = "arrayBool"
	ArrayNumber Type = "arrayNumber"
	ArrayString Type = "arrayString"
	Bool        Type = "bool"
	Number      Type = "number"
	String      Type = "string"
)
