package cmd

import (
	"github.com/olegsu/cli-generator/pkg/generate"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var generateCmdOptions struct {
	language       string
	goPackage      string
	spec           string
	createHandlers bool
}

var generateCmd = &cobra.Command{
	Use: "generate",
	RunE: func(cmd *cobra.Command, args []string) error {
		handler := &generate.Generate{}
		cnf := viper.New()
		cnf.Set("language", generateCmdOptions.language)
		cnf.Set("goPackage", generateCmdOptions.goPackage)
		cnf.Set("spec", generateCmdOptions.spec)
		cnf.Set("projectDir", rootCmdOptions.ProjectDirectory)
		cnf.Set("createHandlers", generateCmdOptions.createHandlers)
		return handler.Handle(cnf)
	},
}

func init() {
	generateCmd.Flags().StringVar(&generateCmdOptions.language, "language", "", "Target language")
	generateCmd.Flags().StringVar(&generateCmdOptions.goPackage, "go-package", "", "Package name of the golang CLI")
	generateCmd.Flags().StringVar(&generateCmdOptions.spec, "spec", "", "Path to CLI.yaml file")
	generateCmd.Flags().BoolVar(&generateCmdOptions.createHandlers, "create-handlers", false, "Create handlers if not exists")
	rootCmd.AddCommand(generateCmd)
}
