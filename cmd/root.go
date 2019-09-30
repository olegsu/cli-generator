package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cnf *viper.Viper

var rootCmdOptions struct {
	Verbose          bool
	ProjectDirectory string
}

var rootCmd = &cobra.Command{
	Use:     "cli-generator",
	Long:    "",
	Version: "0.0.1",
}

// Execute - execute the root command
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		panic(err)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVar(&rootCmdOptions.Verbose, "verbose", false, "Set to get more detailed output")
	rootCmd.PersistentFlags().StringVar(&rootCmdOptions.ProjectDirectory, "project-dir", "", "Set path to project dir")
}
