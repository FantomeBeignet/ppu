package main

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "ppu",
	Short: "PassPhrase (Utils|Ultimate)",
	Long:  `Small tool generate and autocomplete passphrases`,
}

func init() {
	rootCmd.AddCommand(completeCmd)
	rootCmd.AddCommand(generateCmd)
	rootCmd.AddCommand(decodeCmd)
	rootCmd.AddCommand(encodeCmd)
	rootCmd.AddCommand(fromCmd)
	rootCmd.CompletionOptions.DisableDefaultCmd = true
}
