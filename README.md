# Clipboard for Go

CLI tool in Go for cross-platform text copy/paste.

## NOTE

This is the [Original Repository](https://github.com/atotto/clipboard) fork (BSD-3-Clause license), oriented by atotto.

The differences from the original are as follows:

- Go version used has been rewritten to `1.23` or higher
- command has been merged and renamed `go-clip`
- `cobra` is introduced in the cli framework

The basic functionality and behavior follows the original.  
See the [forked link](https://github.com/atotto/clipboard) or [docs documentation](docs/README.original.md) for the original documentation.

## Usage

### copy text

```bash
echo foobazbar | go-clip copy
```

### paste text

```bash
go-clip paste
```
