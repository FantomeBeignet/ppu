package main

import (
	"fmt"
	"strings"

	"github.com/atotto/clipboard"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"

	"git.sr.ht/~fantomebeignet/ppu/encoding"
)

var encodeCmd = &cobra.Command{
	Use:     "encode [data]",
	Short:   "Encode a byte string as a passphrase",
	Aliases: []string{"e", "enc"},
	RunE: func(cmd *cobra.Command, args []string) error {
		var inputVar string
		if len(args) > 0 {
			inputVar = args[0]
		} else {
			km := huh.NewDefaultKeyMap()
			km.Input.AcceptSuggestion = key.NewBinding(
				key.WithKeys("tab"),
				key.WithHelp("tab", "complete"),
			)

			form := huh.NewForm(
				huh.NewGroup(
					huh.NewInput().
						Value(&inputVar).
						Inline(true).
						Title("Data").EchoMode(huh.EchoModePassword),
				),
			).WithKeyMap(km).WithAccessible(accessible).WithTheme(huh.ThemeCatppuccin())
			err := form.Run()
			if err != nil {
				return err
			}
		}
		encoded, err := encoding.FromString(inputVar)
		if err != nil {
			return err
		}
		words := encoded.Words(capitalize)
		val := strings.Join(words, "-")
		if useClipboard {
			if err = clipboard.WriteAll(val); err != nil {
				return err
			}
		} else {
			fmt.Println(val)
		}
		return nil
	},
}

func init() {
	encodeCmd.Flags().
		BoolVarP(&useClipboard, "copy", "c", false, "copy passphrase to clipboard instead of printing to stdout")
	encodeCmd.Flags().
		BoolVarP(&capitalize, "capitalize", "C", false, "copy passphrase to clipboard instead of printing to stdout")
	encodeCmd.Flags().
		BoolVarP(&accessible, "accessible", "a", false, "use a more accessible rendering mode")
}
