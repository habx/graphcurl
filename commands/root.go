package commands

import (
	"fmt"

	"github.com/habx/graphcurl/commands/post"
	"github.com/habx/graphcurl/flags"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// RootCommand: is the main command of the CLI.
var (
	RootCommand = &cobra.Command{
		Use:     "graphcurl",
		Version: flags.Version,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			v := viper.New()
			flags.BindFlags(cmd, v)
			return nil
		},
	}
)

func init() {
	// global flags
	RootCommand.PersistentFlags().StringVarP(&flags.LogLevel, "loglevel", "v", "info", "Set level log [debug,info]")
	PersistentFlagsErr(viper.BindPFlag("loglevel", RootCommand.PersistentFlags().Lookup("loglevel")))
	RootCommand.PersistentFlags().BoolVarP(&flags.Slient, "slient", "s", false, "Enable slient mode")
	PersistentFlagsErr(viper.BindPFlag("slient", RootCommand.PersistentFlags().Lookup("slient")))

	// Subcommands
	RootCommand.AddCommand(post.Command)
	RootCommand.AddCommand(&cobra.Command{
		Use:              "version",
		Short:            "Print the version number and quit",
		TraverseChildren: true,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Version: %s", flags.Version)
		},
	})
}
func PersistentFlagsErr(err error) {
	if err != nil {
		panic(err)
	}
}
