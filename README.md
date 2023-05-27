# banner
Golang app / package to generate a banner of text

## Building

### Requirements

- **Go 1.19**
- (Optional) **Make**

A binary can be built either via the Makefile target `make build` or via `go build` directly.
The Makefile target will place the binary into the `bin` directory.

```zsh
❯ make build
mkdir -p bin
go build -ldflags "-s -w " -o bin/banner cmd/banner/main.go

❯ ls bin
banner
```

## CLI Usage

```zsh
❯ make install
go install ./cmd/banner
❯ banner
Error: requires at least 1 arg(s), only received 0
Usage:
  banner 'Surround this text' [flags]

Flags:
      --char string          The characer to surround the text with (=) (default "=")
      --color string         The color of the banner (none) oneOf [BLUE WHITE BRIGHT_BLACK BRIGHT_RED BRIGHT_GREEN BRIGHT_YELLOW GREEN MAGENTA CYAN BRIGHT_CYAN BLACK BRIGHT_BLUE BRIGHT_MAGENTA YELLOW BRIGHT_WHITE RED]
      --frame-left string    The left framing character ([) (default "[")
      --frame-right string   The right framing character (]) (default "[")
  -h, --help                 help for banner
      --length int           The total length of the banner (80) (default 80)

requires at least 1 arg(s), only received 0
```

**Example (zsh)**

```zsh
❯ banner "Surround me" --color green --char '*' --frame-left '{' --frame-right '}' --length 40
*************{ Surround me }*************
```

## Package Usage

See [examples.go](./examples/examples.go) for usage within a Go project.
