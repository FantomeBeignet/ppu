package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "ppu",
	Short: "PassPhrase (Utils|Ultimate)",
	Long:  `Small tool generate and autocomplete passphrases`,
}

func init() {
	rootCmd.AddCommand(completeCmd)
	rootCmd.AddCommand(generateCmd)
	rootCmd.CompletionOptions.DisableDefaultCmd = true
}

func Execute() error {
	return rootCmd.Execute()
}
