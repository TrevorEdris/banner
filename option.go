package banner

import (
	"errors"
	"strings"
)

type (
	Option struct {
		color      Color
		length     int
		char       rune
		frameLeft  rune
		frameRight rune
	}

	Color     string
	ColorFunc func() Option
)

var (
	black         Color = "\u001b[30m"
	red           Color = "\u001b[31m"
	green         Color = "\u001b[32m"
	yellow        Color = "\u001b[33m"
	blue          Color = "\u001b[34m"
	magenta       Color = "\u001b[35m"
	cyan          Color = "\u001b[36m"
	white         Color = "\u001b[37m"
	brightBlack   Color = "\u001b[90m"
	brightRed     Color = "\u001b[91m"
	brightGreen   Color = "\u001b[92m"
	brightYellow  Color = "\u001b[93m"
	brightBlue    Color = "\u001b[94m"
	brightMagenta Color = "\u001b[95m"
	brightCyan    Color = "\u001b[96m"
	brightWhite   Color = "\u001b[97m"
	reset         Color = "\u001b[0m"

	colorFuncs = map[string]ColorFunc{
		"BLACK":          Black,
		"RED":            Red,
		"GREEN":          Green,
		"YELLOW":         Yellow,
		"BLUE":           Blue,
		"MAGENTA":        Magenta,
		"CYAN":           Cyan,
		"WHITE":          White,
		"BRIGHT_BLACK":   BrightBlack,
		"BRIGHT_RED":     BrightRed,
		"BRIGHT_GREEN":   BrightGreen,
		"BRIGHT_YELLOW":  BrightYellow,
		"BRIGHT_BLUE":    BrightBlue,
		"BRIGHT_MAGENTA": BrightMagenta,
		"BRIGHT_CYAN":    BrightCyan,
		"BRIGHT_WHITE":   BrightWhite,
	}
	allColors = []string{}

	ErrUnsupportedColor = errors.New("unsupported color")
)

func WithLength(l int) Option {
	return Option{length: l}
}

func WithChar(c rune) Option {
	return Option{char: c}
}

func WithFrame(left, right rune) Option {
	return Option{
		frameLeft:  left,
		frameRight: right,
	}
}

func GetColorFunc(color string) (ColorFunc, error) {
	color = strings.ToUpper(color)
	color = strings.TrimSpace(color)
	if color == "" {
		return resetColor, nil
	}

	if f, exists := colorFuncs[color]; exists {
		return f, nil
	}

	return func() Option { return Option{} }, ErrUnsupportedColor
}

func AvailableColors() []string {
	// Simple memoization in case this is called multiple times
	if len(allColors) > 0 {
		return allColors
	}

	cs := make([]string, 0)
	for color := range colorFuncs {
		cs = append(cs, color)
	}
	allColors = cs

	return cs
}

func Black() Option {
	return Option{color: black}
}

func Red() Option {
	return Option{color: red}
}

func Green() Option {
	return Option{color: green}
}

func Yellow() Option {
	return Option{color: yellow}
}

func Blue() Option {
	return Option{color: blue}
}

func Magenta() Option {
	return Option{color: magenta}
}

func Cyan() Option {
	return Option{color: cyan}
}

func White() Option {
	return Option{color: white}
}

func BrightBlack() Option {
	return Option{color: brightBlack}
}

func BrightRed() Option {
	return Option{color: brightRed}
}

func BrightGreen() Option {
	return Option{color: brightGreen}
}

func BrightYellow() Option {
	return Option{color: brightYellow}
}

func BrightBlue() Option {
	return Option{color: brightBlue}
}

func BrightMagenta() Option {
	return Option{color: brightMagenta}
}

func BrightCyan() Option {
	return Option{color: brightCyan}
}

func BrightWhite() Option {
	return Option{color: brightWhite}
}

func resetColor() Option {
	return Option{}
}
