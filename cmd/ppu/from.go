package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/atotto/clipboard"
	"github.com/spf13/cobra"

	"git.sr.ht/~fantomebeignet/ppu/encoding"
)

var salt string

var fromCmd = &cobra.Command{
	Use:   "from [data]",
	Short: "Generate a secure passphrase from data",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := cobra.ExactArgs(2)(cmd, args); err != nil {
			return err
		}
		n, err := strconv.Atoi(args[1])
		if err != nil {
			return err
		}
		encoded := encoding.NewFromSeed([]byte(args[0]), []byte(salt), uint(n))
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
	fromCmd.Flags().
		BoolVarP(&useClipboard, "copy", "c", false, "copy passphrase to clipboard instead of printing to stdout")
	fromCmd.Flags().
		BoolVarP(&capitalize, "capitalize", "C", false, "copy passphrase to clipboard instead of printing to stdout")
	fromCmd.Flags().
		BoolVarP(&accessible, "accessible", "a", false, "use a more accessible rendering mode")
	fromCmd.Flags().StringVarP(&salt, "salt", "s", "", "salt to pass to the derivation function")
}
