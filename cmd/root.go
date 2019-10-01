
// Code generated by cli-generator; DO NOT EDIT.
package cmd



import (
	
	"github.com/spf13/viper"
	
	"github.com/spf13/cobra"
	
)
var cnf *viper.Viper = viper.New()

var rootCmdOptions struct {
	verbose bool
	
}

var rootCmd = &cobra.Command{
	Use:     "cli-generator",
	Version: "0.5.0",
	
	Long: "Generate CLI entrypoints from spec file",
	PreRun: func(cmd *cobra.Command, args []string) {
		
		cnf.Set("verbose", rootCmdOptions.verbose)
		
	},
}



// Execute - execute the root command
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		panic(err)
	}
}


func init() {
	cnf.BindEnv("verbose", "VERBOSE")
	cnf.SetDefault("verbose", false)

	rootCmd.Flags().BoolVar(&rootCmdOptions.verbose, "verbose", cnf.GetBool("verbose"), "Set to see more logs")
}