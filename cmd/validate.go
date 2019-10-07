// Code generated by cli-generator; DO NOT EDIT.
package cmd



import (
	
	handler "github.com/olegsu/cli-generator/pkg/validate"
	
	"github.com/spf13/cobra"
)

var validateCmdOptions struct {
	
}

var validateCmd = &cobra.Command{
	Use:     "validate",
	Args: func (cmd *cobra.Command, args []string) error {
		var validators []func(cmd *cobra.Command, args []string) error
		validators = append(validators, cobra.MinimumNArgs(1))
		for _, v := range validators {
			if err := v(cmd, args); err != nil {
				return err
			}
		}
		return nil
	},

	RunE: func(cmd *cobra.Command, args []string) error {
		h := &handler.Handler{}
		return h.Handle(cnf)
	},
	Long: "Generate CLI entrypoints from spec file",
	PreRun: func(cmd *cobra.Command, args []string) {
		cnf.Set("spec", args )
		rootCmd.PreRun(cmd, args)
		
	},
}




func init() {
	rootCmd.AddCommand(validateCmd)
}