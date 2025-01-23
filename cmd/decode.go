package cmd

import (
	"fmt"
	"strings"

	"github.com/atotto/clipboard"
	"github.com/spf13/cobra"

	"git.sr.ht/~fantomebeignet/ppu/internal/encoding"
	"git.sr.ht/~fantomebeignet/ppu/internal/input"
)

var (
	hex = true
	oct = false
	bin = false
	dec = false
)

func base() int {
	if oct {
		return 8
	}
	if bin {
		return 2
	}
	if dec {
		return 10
	}
	return 16
}

func prefix() string {
	if oct {
		return "0o"
	}
	if bin {
		return "0b"
	}
	if dec {
		return ""
	}
	return "0x"
}

var decodeCmd = &cobra.Command{
	Use:     "decode [passphrase]",
	Short:   "Decode a passphrase as a byte string",
	Aliases: []string{"d", "dec"},
	RunE: func(cmd *cobra.Command, args []string) error {
		var inputVar string
		if len(args) > 0 {
			inputVar = args[0]
		} else {
			form := input.NewPassphraseInput(&inputVar, accessible)
			err := form.Run()
			if err != nil {
				return err
			}
		}
		encoded, err := encoding.FromWords(strings.Split(inputVar, "-"))
		if err != nil {
			return err
		}
		val := prefix() + encoded.ToString(base())
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
	decodeCmd.Flags().
		BoolVarP(&useClipboard, "copy", "c", false, "copy passphrase to clipboard instead of printing to stdout")
	decodeCmd.Flags().
		BoolVarP(&accessible, "accessible", "a", false, "use a more accessible rendering mode")

	decodeCmd.Flags().BoolVar(&hex, "hex", true, "output result as hexadecimal")
	decodeCmd.Flags().BoolVar(&bin, "oct", false, "output result as octal")
	decodeCmd.Flags().BoolVar(&bin, "bin", false, "output result as binary")
	decodeCmd.Flags().BoolVar(&bin, "dec", false, "output result as decimal")
	decodeCmd.MarkFlagsMutuallyExclusive("hex", "oct", "bin", "dec")
}
