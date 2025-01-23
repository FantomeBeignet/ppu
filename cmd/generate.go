package cmd

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"strconv"
	"strings"

	wl "github.com/kklash/wordlist4096"
	"github.com/spf13/cobra"
)

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
		buf := make([]byte, n*2)
		if _, err := rand.Read(buf); err != nil {
			return err
		}
		words := make([]string, n)
		for i := range n {
			index := binary.BigEndian.Uint16(buf[2*i : 2*(i+1)])
			words[i] = wl.WordList[index%4096]
		}
		fmt.Println(strings.Join(words, "-"))
		return nil
	},
}
