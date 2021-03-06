// Code generated by cli-generator; DO NOT EDIT.
package cmd



import (
	
	handler "github.com/olegsu/cli-generator/pkg/generate"
	
	"github.com/spf13/cobra"
)

var generateCmdOptions struct {
	language string
	goPackage string
	spec string
	projectDir string
	createHandlers bool
	runInitFlow bool
	runPostInitFlow bool
	
}

var generateCmd = &cobra.Command{
	Use:     "generate",

	RunE: func(cmd *cobra.Command, args []string) error {
		h := &handler.Handler{}
		return h.Handle(cnf)
	},
	Long: "Generate CLI entrypoints from spec file",
	PreRun: func(cmd *cobra.Command, args []string) {
		rootCmd.PreRun(cmd, args)
		
		cnf.Set("language", generateCmdOptions.language)
		
		cnf.Set("goPackage", generateCmdOptions.goPackage)
		
		cnf.Set("spec", generateCmdOptions.spec)
		
		cnf.Set("projectDir", generateCmdOptions.projectDir)
		
		cnf.Set("createHandlers", generateCmdOptions.createHandlers)
		
		cnf.Set("runInitFlow", generateCmdOptions.runInitFlow)
		
		cnf.Set("runPostInitFlow", generateCmdOptions.runPostInitFlow)
		
	},
}




func init() {

	
	generateCmd.PersistentFlags().StringVar(&generateCmdOptions.language, "language", cnf.GetString("language"), "The target language of generated code [options: go]")

	generateCmd.PersistentFlags().StringVar(&generateCmdOptions.goPackage, "go-package", cnf.GetString("goPackage"), "")

	generateCmd.PersistentFlags().StringVar(&generateCmdOptions.spec, "spec", cnf.GetString("spec"), "")

	generateCmd.PersistentFlags().StringVar(&generateCmdOptions.projectDir, "project-dir", cnf.GetString("projectDir"), "")

	generateCmd.PersistentFlags().BoolVar(&generateCmdOptions.createHandlers, "create-handlers", cnf.GetBool("createHandlers"), "")

	generateCmd.PersistentFlags().BoolVar(&generateCmdOptions.runInitFlow, "run-init-flow", cnf.GetBool("runInitFlow"), "")

	generateCmd.PersistentFlags().BoolVar(&generateCmdOptions.runPostInitFlow, "run-post-init-flow", cnf.GetBool("runPostInitFlow"), "")
	rootCmd.AddCommand(generateCmd)
}