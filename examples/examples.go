package main

import (
	"fmt"

	"github.com/TrevorEdris/banner"
)

const (
	text = "Surround this text"
)

func simpleBanner() {
	b := banner.New(text)
	fmt.Println(b)
}

func bannerWithColor() {
	b := banner.New(text, banner.Blue())
	fmt.Println(b)
}

func bannerWithLength() {
	b := banner.New(text, banner.WithLength(120))
	fmt.Println(b)
}

func bannerWithFrame() {
	b := banner.New(text, banner.WithFrame('(', ')'))
	fmt.Println(b)
}

func bannerWithChar() {
	b := banner.New(text, banner.WithChar('*'))
	fmt.Println(b)
}

func bannerWithDuplicateOpts() {
	// Later values of the same option will override the earlier values.
	// In this example, the following options will take effect:
	// -
	// 100
	// ~ ~
	// Cyan()
	b := banner.New(
		text,
		banner.WithChar('*'),
		banner.WithChar('-'),
		banner.WithLength(10),
		banner.WithLength(100),
		banner.WithFrame('{', '}'),
		banner.WithFrame('~', '~'),
		banner.Magenta(),
		banner.Cyan(),
	)
	fmt.Println(b)
}

func main() {
	simpleBanner()
	bannerWithColor()
	bannerWithLength()
	bannerWithFrame()
	bannerWithChar()
	bannerWithDuplicateOpts()
}
