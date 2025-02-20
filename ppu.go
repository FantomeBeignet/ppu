package ppu

import (
	"errors"
	"slices"
	"strings"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/huh"
	wl "github.com/kklash/wordlist4096"
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

func NewPassphraseInput(inputVar *string, title string, accessible bool) *huh.Input {
	return huh.NewInput().
		Value(inputVar).
		Inline(true).
		Title(title).
		SuggestionsFunc(func() []string {
			return suggest(*inputVar)
		}, inputVar).
		Validate(validate)
}

func NewPassphraseInputForm(inputVar *string, title string, accessible bool) *huh.Form {
	km := huh.NewDefaultKeyMap()
	km.Input.AcceptSuggestion = key.NewBinding(
		key.WithKeys("tab"),
		key.WithHelp("tab", "complete"),
	)
	km.Quit = key.NewBinding(key.WithKeys("esc"))
	form := huh.NewForm(
		huh.NewGroup(
			NewPassphraseInput(inputVar, title, accessible),
		),
	).WithKeyMap(km).WithAccessible(accessible).WithTheme(huh.ThemeCatppuccin())
	return form
}
