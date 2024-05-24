package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"go-blog-ddd/pkg/shutdown"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "blog",
	Short: "blog is root commandhandler",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

func init() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%s", err)
		shutdown.Exit(1)
	}
	shutdown.WaitExit()
}
