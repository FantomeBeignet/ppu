package main

import (
	"fmt"

	"github.com/atotto/clipboard"
	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"

	"git.sr.ht/~fantomebeignet/ppu"
)

var (
	useClipboard bool
	accessible   bool
)

var completeCmd = &cobra.Command{
	Use:     "complete",
	Short:   "Input a passphrase, with autocomplete",
	Aliases: []string{"c", "comp"},
	RunE: func(cmd *cobra.Command, args []string) error {
		var inputVar string
		form := ppu.NewPassphraseInputForm(&inputVar, "Passphrase", accessible)
		err := form.Run()
		if err == huh.ErrUserAborted {
			return nil
		} else if err != nil {
			panic(err)
		}
		if useClipboard {
			if err = clipboard.WriteAll(inputVar); err != nil {
				return err
			}
		} else {
			fmt.Println(inputVar)
		}
		return nil
	},
}

func init() {
	completeCmd.Flags().
		BoolVarP(&useClipboard, "copy", "c", false, "copy passphrase to clipboard instead of printing to stdout")
	completeCmd.Flags().
		BoolVarP(&accessible, "accessible", "a", false, "use a more accessible rendering mode")
}
