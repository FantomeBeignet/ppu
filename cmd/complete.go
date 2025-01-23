package cmd

import (
	"errors"
	"fmt"
	"log"
	"slices"
	"strings"

	"github.com/atotto/clipboard"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/huh"
	wl "github.com/kklash/wordlist4096"
	"github.com/spf13/cobra"
)

func suggest(p string) []string {
	words := strings.Split(strings.ToLower(p), "-")
	lastWord := strings.ToLower(words[len(words)-1])
	res := []string{}
	for _, suff := range wl.Search(lastWord).Suffixes {
		res = append(res, p+suff)
	}
	return res
}

func validate(p string) error {
	words := strings.Split(strings.ToLower(p), "-")
	for _, w := range words {
		if !slices.Contains(wl.WordList, strings.ToLower(w)) {
			return errors.New("invalid passphrase")
		}
	}
	return nil
}

var (
	useClipboard bool
	accessible   bool
)

var completeCmd = &cobra.Command{
	Use:     "complete",
	Short:   "Input a passphrase, with autocomplete",
	Aliases: []string{"c", "comp"},
	RunE: func(cmd *cobra.Command, args []string) error {
		km := huh.NewDefaultKeyMap()
		km.Input.AcceptSuggestion = key.NewBinding(
			key.WithKeys("tab"),
			key.WithHelp("tab", "complete"),
		)

		var input string

		form := huh.NewForm(
			huh.NewGroup(
				huh.NewInput().
					Value(&input).
					Inline(true).
					Title("Passphrase").
					SuggestionsFunc(func() []string {
						return suggest(input)
					}, &input).
					Validate(validate),
			),
		).WithKeyMap(km).
			WithAccessible(accessible)
		err := form.Run()
		if err != nil {
			log.Fatal(err)
		}
		if useClipboard {
			if err = clipboard.WriteAll(input); err != nil {
				return err
			}
		} else {
			fmt.Println(input)
		}
		return nil
	},
}

func init() {
	completeCmd.Flags().
		BoolVarP(&useClipboard, "copy", "c", false, "Copy passphrase to clipboard instead of printing to stdout")
	completeCmd.Flags().
		BoolVarP(&accessible, "accessible", "a", false, "Use a more accessible rendering mode")
}
