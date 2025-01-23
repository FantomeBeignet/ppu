# ppu

Generate memorable passphrases, and get help typing them in case they aren't so
memorable to you.

## Usage

```console
user@host:~$ ppu help
Small tool generate and autocomplete passphrases

Usage:
  ppu [command]

Available Commands:
  complete    Input a passphrase, with autocomplete
  generate    Generate a passphrase
  help        Help about any command

Flags:
  -h, --help   help for ppu

Use "ppu [command] --help" for more information about a command.
```

### Generating a passphrase

```console
user@host:~$ ppu help generate
Generate a passphrase

Usage:
  ppu generate <number of words> [flags]

Aliases:
  generate, g, gen

Flags:
  -c, --capitalize   capitalize each word of the passphase
  -h, --help         help for generate
```

Given an argument `x`, `ppu generate x` prints an x-word long passphrase to
`stdout`. The words in the passphrase are separated by dashes.

### Autocompleting a passphrase

```console
user@host:~$ ppu help complete
Input a passphrase, with autocomplete

Usage:
  ppu complete [flags]

Aliases:
  complete, c, comp

Flags:
  -a, --accessible   use a more accessible rendering mode
  -c, --copy         copy passphrase to clipboard instead of printing to stdout
  -h, --help         help for complete
```

`ppu complete` gives you a prompt to help type your passphrase in, providing
autocompletion on each word of the passphrase. The resulting passphrase is then
printed to `stdout`.

## Notes

The wordlist used to generate the passphrases is [wordlist4096](https://github.com/kklash/wordlist4096).
