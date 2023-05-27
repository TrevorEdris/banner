package banner

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

const (
	defaultLength     = 80
	defaultChar       = '='
	defaultColor      = ""
	defaultFrameLeft  = '['
	defaultFrameRight = ']'
)

func Banner(text string, opts ...Option) string {
	length := defaultLength
	char := defaultChar
	color := defaultColor
	frameLeft := defaultFrameLeft
	frameRight := defaultFrameRight
	for _, o := range opts {
		if o.length != 0 {
			length = o.length
		}
		if o.char != 0 {
			char = o.char
		}
		if o.color != "" {
			color = string(o.color)
		}
		if o.frameLeft != 0 {
			frameLeft = o.frameLeft
		}
		if o.frameRight != 0 {
			frameRight = o.frameRight
		}
	}

	textLength := utf8.RuneCountInString(text)
	if length <= textLength {
		return text
	}

	surroundingChars := (length - textLength - 4) / 2

	bannerText := color
	bannerText += strings.Repeat(string(char), surroundingChars)
	bannerText += fmt.Sprintf("%c%c %s %c%c", char, frameLeft, text, frameRight, char)
	bannerText += strings.Repeat(string(char), surroundingChars)
	bannerText += string(reset)

	return bannerText
}
