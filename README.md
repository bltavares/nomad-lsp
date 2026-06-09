# Nomad LSP

![nomad version](https://img.shields.io/badge/nomad-1.11.3-blue.svg)

This is LSP(Language Server Protocol) for Nomad

[![asciicast](https://asciinema.org/a/246266.svg)](https://asciinema.org/a/246266)


## Installation

```sh
go install https://github.com/bltavares/nomad-lsp@latest
```

## Features

- Error checking
- Autocompletion
- Formatting

## To-do

- Add the full schema from nomad jobspec

## Credits
- LSP structure using [Sourcegraph's go-lsp](https://github.com/sourcegraph/go-lsp)
- JSON-RPC 2.0 using [jrpc2](https://github.com/creachadair/jrpc2)

## License
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fbltavares%2Fnomad-lsp.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2Fbltavares%2Fnomad-lsp?ref=badge_large)
