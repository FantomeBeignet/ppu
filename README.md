# ppu

Generate memorable passphrases, and get help typing them in case they aren't so
memorable to you.

<img alt="ppu demo" width="800" src="demo/ppu.gif">

## Usage

```console
user@host:~$ ppu help
Small tool generate and autocomplete passphrases

Usage:
  ppu [command]

Available Commands:
  complete    Input a passphrase, with autocomplete
  decode      Decode a passphrase as a byte string
  encode      Encode a byte string as a passphrase
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

### Encoding data as a passphrase

```console
user@host:~$ ppu help encode
Encode a byte string as a passphrase

Usage:
  ppu encode [data] [flags]

Aliases:
  encode, e, enc

Flags:
  -a, --accessible   use a more accessible rendering mode
  -C, --capitalize   copy passphrase to clipboard instead of printing to stdout
  -c, --copy         copy passphrase to clipboard instead of printing to stdout
  -h, --help         help for encode
```

`ppu encode` encodes data (represented in hex, octal, binary or decimal) as a
passphrase. The data can either be passed directly through the command line,
or through an password input prompt if no argument is passed.

### Decoding a passphrase as data

```console
user@host:~$ ppu help decode
Decode a passphrase as a byte string

Usage:
  ppu decode [passphrase] [flags]

Aliases:
  decode, d, dec

Flags:
  -a, --accessible   use a more accessible rendering mode
      --bin          output result as binary
  -c, --copy         copy passphrase to clipboard instead of printing to stdout
      --dec          output result as decimal
  -h, --help         help for decode
      --hex          output result as hexadecimal (default true)
      --oct          output result as octal
```

`ppu decode` decodes a passphrase, passed either through the command line or an
autocompleting input prompt, as a byte string, either in hex, octal, binary or
decimal format.

## Notes

The wordlist used to generate the passphrases is [wordlist4096](https://github.com/kklash/wordlist4096).
