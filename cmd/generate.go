package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"

	"git.sr.ht/~fantomebeignet/ppu/internal/encoding"
)

var capitalize bool

var generateCmd = &cobra.Command{
	Use:     "generate <number of words>",
	Short:   "Generate a passphrase",
	Aliases: []string{"g", "gen"},
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := cobra.ExactArgs(1)(cmd, args); err != nil {
			return err
		}
		n, err := strconv.Atoi(args[0])
		if err != nil {
			return err
		}
		encoded, err := encoding.NewRandom(uint(n))
		if err != nil {
			return err
		}
		words := encoded.Words(capitalize)
		fmt.Println(strings.Join(words, "-"))
		return nil
	},
}

func init() {
	generateCmd.Flags().
		BoolVarP(&capitalize, "capitalize", "c", false, "capitalize each word of the passphase")
}
