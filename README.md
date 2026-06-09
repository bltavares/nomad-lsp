# Nomad LSP

![nomad version](https://img.shields.io/badge/nomad-1.11.3-green.svg)

This is LSP(Language Server Protocol) for Nomad

[![asciicast](https://asciinema.org/a/246266.svg)](https://asciinema.org/a/246266)


## Installation

```sh
# 1. Build from source
go install github.com/bltavares/nomad-lsp@latest

# 2. Donwload from the release pages

# 3. Install with mise
# mise use github:bltavares/nomad-lsp@latest # Pre-Built binaries
# mise use go:github.com/bltavares/nomad-lsp@latest # Build from source
```

## Features

- Error checking
- Autocompletion
- Formatting

## To-do

- Add the full schema from nomad jobspec for autocompletion

## Credits
- LSP structure using [Sourcegraph's go-lsp](https://github.com/sourcegraph/go-lsp)
- JSON-RPC 2.0 using [jrpc2](https://github.com/creachadair/jrpc2)
- @juliosueiras for their [initial implementation](https://github.com/juliosueiras/nomad-lsp)
