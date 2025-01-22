package main

import (
	"errors"
	"fmt"
	"log"
	"slices"
	"strings"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/huh"
	wl "github.com/kklash/wordlist4096"
)

func suggest(p string) []string {
	words := strings.Split(strings.ToLower(p), "-")
	lastWord := words[len(words)-1]
	res := []string{}
	for _, suff := range wl.Search(lastWord).Suffixes {
		res = append(res, p+suff)
	}
	return res
}

func validate(p string) error {
	words := strings.Split(strings.ToLower(p), "-")
	for _, w := range words {
		if !slices.Contains(wl.WordList, w) {
			return errors.New("invalid passphrase")
		}
	}
	return nil
}

func main() {
	km := huh.NewDefaultKeyMap()
	km.Input.AcceptSuggestion = key.NewBinding(key.WithKeys("tab"), key.WithHelp("tab", "complete"))

	var input string

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Value(&input).
				Inline(true).
				Title("Passphrase").
				Placeholder("word").
				SuggestionsFunc(func() []string {
					return suggest(input)
				}, &input).
				Validate(validate),
		),
	).WithKeyMap(km)
	err := form.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(input)
}
