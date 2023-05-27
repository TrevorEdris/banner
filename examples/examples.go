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

func blockBanner() {
	header := "HEADER"
	footer := ""
	block := `Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed hendrerit interdum consectetur. Duis condimentum nulla eu fringilla bibendum. Integer eget elementum nibh. Phasellus dignissim eleifend felis at eleifend. Aliquam congue tortor ligula, in tempus odio dignissim eu. Proin commodo nulla ut dui sollicitudin malesuada. Suspendisse nec mauris non lorem condimentum euismod nec et nisl. Donec sollicitudin ex sit amet orci auctor, a tincidunt turpis feugiat. Aliquam eget justo purus. Suspendisse tincidunt, nunc at lobortis auctor, velit mauris fermentum ligula, eu sagittis nisl nulla at mi. Curabitur id tellus libero. In finibus mauris a eros interdum placerat.`

	b := banner.NewBlock(header, block, footer, banner.Cyan())
	fmt.Println(b)
}

func blockBannerWithFooter() {
	header := "HEADER"
	footer := "FOOTER"
	block := `Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed hendrerit interdum consectetur. Duis condimentum nulla eu fringilla bibendum. Integer eget elementum nibh. Phasellus dignissim eleifend felis at eleifend. Aliquam congue tortor ligula, in tempus odio dignissim eu. Proin commodo nulla ut dui sollicitudin malesuada. Suspendisse nec mauris non lorem condimentum euismod nec et nisl. Donec sollicitudin ex sit amet orci auctor, a tincidunt turpis feugiat. Aliquam eget justo purus. Suspendisse tincidunt, nunc at lobortis auctor, velit mauris fermentum ligula, eu sagittis nisl nulla at mi. Curabitur id tellus libero. In finibus mauris a eros interdum placerat.`

	b := banner.NewBlock(header, block, footer, banner.Yellow())
	fmt.Println(b)
}

func blockBannerLongWords() {
	header := "HEADER"
	footer := ""
	block := "Thisisonereallylongwordthatshouldbesplitontomultiplelinesbutidkhowitllworkwithover2lengthsoflines"

	b := banner.NewBlock(header, block, footer, banner.Magenta())
	fmt.Println(b)
}

func main() {
	simpleBanner()
	bannerWithColor()
	bannerWithLength()
	bannerWithFrame()
	bannerWithChar()
	bannerWithDuplicateOpts()

	for _, color := range banner.AvailableColors() {
		f, _ := banner.GetColorFunc(color)
		fmt.Printf("%s - %s\n", banner.New(text, f()), color)
	}
	blockBanner()
	blockBannerWithFooter()
	blockBannerLongWords()
}
