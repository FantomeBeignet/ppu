package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "ppu",
	Short: "PassPhrase (Ultimate|Utils)",
	Long:  `Small tool generate and autocomplete passphrases`,
}

func init() {
	rootCmd.AddCommand(completeCmd)
	rootCmd.CompletionOptions.DisableDefaultCmd = true
}

func Execute() error {
	return rootCmd.Execute()
}
