
// Code generated by go generate; DO NOT EDIT.
// using data from templates/templates
package templates

func TemplatesMap() map[string]string {
    templatesMap := make(map[string]string)

templatesMap["go.cmd.tmpl"] = `// Code generated by cli-generator; DO NOT EDIT.
package cmd

{{ $name := strings.CamelCase .cmd.name }}

import (
	{{ if .cmd.root }}
	"github.com/spf13/viper"
	"fmt"
	"os"
	{{ end }}

	{{- if not .cmd.loose }}
	handler "{{ .go.package }}/pkg/{{$name}}"
	{{end}}
	"github.com/spf13/cobra"
)

{{- if .cmd.root }}
var cnf *viper.Viper = viper.New()
{{- end }}

var {{$name}}CmdOptions struct {
	{{ range .cmd.flags }}
	{{- .name | strings.CamelCase }} {{ .type | toGolangType }}
	{{ end }}
}

var {{$name}}Cmd = &cobra.Command{
	{{- if .cmd.root }}
	Use:     "{{ .spec.metadata.name }}",
	Version: "{{ .spec.metadata.version }}",
	{{- else }}
	Use:     "{{ .cmd.name }}",
	{{- end }}
	{{- if not .cmd.loose }}

	{{- if (has .cmd "arg" )}}
	Args: func (cmd *cobra.Command, args []string) error {
		var validators []func(cmd *cobra.Command, args []string) error
		{{- range .cmd.arg.rules }}
		validators = append(validators, {{ . | golangRulesToArgsValidation }})
		{{- end}}
		for _, v := range validators {
			if err := v(cmd, args); err != nil {
				return err
			}
		}
		return nil
	},
	{{- end }}

	RunE: func(cmd *cobra.Command, args []string) error {
		h := &handler.Handler{}
		return h.Handle(cnf)
	},
	{{- end }}
	{{- if .spec.metadata.description }}
	Long: "{{ .spec.metadata.description }}",
	{{- end }}
	PreRun: func(cmd *cobra.Command, args []string) {

		{{- if (has .cmd "arg" )}}
		cnf.Set("{{.cmd.arg.name}}", args )
		{{- end }}

		{{- if not .cmd.root }}
		{{ .cmd.parent }}Cmd.PreRun(cmd, args)
		{{- end }}
		{{ range .cmd.flags }}
		cnf.Set("{{- .name | strings.CamelCase }}", {{$name}}CmdOptions.{{- .name | strings.CamelCase }})
		{{ end }}
	},
}


{{ if .cmd.root }}
// Execute - execute the root command
func Execute() {
	err := {{$name}}Cmd.Execute()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
{{ end }}

func init() {
{{- range .cmd.flags }}
	{{- if .envVar }}
	cnf.BindEnv("{{ .name }}", "{{ .envVar }}")
	{{- end }}
	{{- if ( has . "default" ) }}
	{{ if .type eq strgin }}
	cnf.SetDefault("{{ strings.CamelCase .name}}", {{ .default | qoute }})
	{{ else }}
	cnf.SetDefault("{{ strings.CamelCase .name}}", {{ .default }})
	{{ end }}
	{{- end }}

	{{ $description := "" }}
	{{- if (has . "description" ) }}
	{{ $description = .description }}
	{{- end }}

	{{- if (and ( has . "enum" ) .enum ) -}}
		{{- $description = (printf "%s [options:" $description ) -}}
		{{- range .enum -}}
		{{- $description = (printf "%s %s" $description .) -}}
		{{- end -}}
		{{- $description = (printf "%s]" $description ) -}}
	{{- end -}}

	{{- if (has . "envVar" ) -}}
		{{- $description = (printf "%s [$%s]" $description .envVar ) -}}
	{{- end -}}
	
	{{ $name }}Cmd.PersistentFlags().{{ .type | golangFlagFunc }}(&{{ $name }}CmdOptions.{{- .name | strings.CamelCase }}, "{{- .name }}", cnf.{{ golangFlagDefaultFunc .type }}("{{ strings.CamelCase .name}}"), "{{ $description }}")
	
{{- end }}

{{- if not .cmd.root }}
	{{ .cmd.parent }}Cmd.AddCommand({{$name}}Cmd)
{{- end }}
}` 

templatesMap["go.handler.tmpl"] = `package {{ strings.CamelCase .cmd.name }}

import (
	"fmt"
	"github.com/spf13/viper"
)

type (
	// Handler - exposed struct that implementd Handler interface
	Handler struct{}
)

// Handle - the function that will be called from the CLI with viper config
// to provide access to all flags
func (g *Handler) Handle(cnf *viper.Viper) error {
	fmt.Printf("Handler for command: {{ .cmd.name }}\n")
	{{ range .cmd.flags }}
	{{ if eq .type "arrayString" }}
	data := cnf.{{ .type | golangFlagDefaultFunc }}("{{.name}}")
	for index, d := range data {
		fmt.Printf("flag name: {{ .name }}_%d, value: %s\n", index, d)
	}
	{{ else }}
	fmt.Printf("flag name: {{ .name }}, value: %s\n", cnf.{{ .type | golangFlagDefaultFunc }}("{{.name}}"))
	{{ end }}
	{{ end }}
	return nil
}` 

templatesMap["go.main.tmpl"] = `// Code generated by cli-generator; DO NOT EDIT.
package main

import (
	cmd "{{ .go.package }}/cmd"
)

func main() {
	cmd.Execute()
}` 

templatesMap["go.makefile.tmpl"] = `outfile = {{.spec.metadata.name}}
build:
	@echo "Building go binary"
	@go build -o $(outfile) *.go
	@chmod +x $(outfile)
` 

templatesMap["spec.json"] = `{
    "$schema": "http://json-schema.org/draft-07/schema#",
    "type": "object",
    "definitions": {
        "flag": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "type": {
                    "type": "string",
                    "enum": [
                        "string",
                        "bool",
                        "number",
                        "arrayBool",
                        "arrayString",
                        "arrayNumber"
                    ]
                },
                "alias": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    },
                    "default": []
                },
                "default": {
                    "type": "string"
                },
                "required": {
                    "type": "boolean",
                    "default": false                    
                },
                "enum": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "description": {
                    "type": "string"
                },
                "envVar": {
                    "type": "string"
                }
            },
            "required": [
                "name",
                "required",
                "type"
            ]
        },
        "command": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string",
                    "pattern": ".+"
                },
                "parent": {
                    "type": "string",
                    "description": "Applicative property, user data will be ignored"
                },
                "root": {
                    "type": "boolean",
                    "description": "Applicative property, user data will be ignored",
                    "default": false
                },
                "flags": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/flag"
                    }
                },
                "loose": {
                    "type": "boolean",
                    "default": false
                },
                "arg": {
                    "$ref": "#/definitions/argument"
                },
                "commands": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/command"
                    }
                }
            },
            "required": [
                "name",
                "root"
            ]
        },
        "argument": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "rules": {
                    "type": "array",
                    "items": {
                        "type": "string",
                        "enum": [
                            "any",
                            "atLeastOne",
                            "atLeastTwo",
                            "exactlyOne"
                        ]
                    },
                    "default": [
                        "any"
                    ]
                }
            },
            "required": [
                "name"
            ]
        }
    },
    "properties": {
        "metadata": {
            "type": "object",
            "properties": {
                "name": {
                    "description": "CLI Name",
                    "type": "string",
                    "pattern": ".+"
                },
                "version": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                }
            },
            "required": [
                "name",
                "version"
            ]
        },
        "commands": {
            "type": "array",
            "items": {
                "$ref": "#/definitions/command"
            }
        },
        "flags": {
            "type": "array",
            "items": {
                "$ref": "#/definitions/flag"
            }
        },
        "loose": {
            "type": "boolean",
            "default": false
        },
        "arg": {
            "$ref": "#/definitions/argument"
        }
    },
    "required": [
        "metadata"
    ]
}` 

    return  templatesMap
}
