package banner

import (
	"bufio"
	"fmt"
	"math"
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

func New(text string, opts ...Option) string {

	o := consolidateOpts(opts...)

	textLength := utf8.RuneCountInString(text)
	if o.length <= textLength {
		return text
	}

	surroundingChars := (o.length - textLength - 4) / 2

	bannerText := string(o.color)
	bannerText += strings.Repeat(string(o.char), surroundingChars)
	if o.noFrame {
		// Add 2 additional surrounding chars to account for the missing frame
		bannerText += strings.Repeat(string(o.char), 2)
		bannerText += fmt.Sprintf("%c%s%c", o.char, text, o.char)
		bannerText += strings.Repeat(string(o.char), 2)
	} else {
		bannerText += fmt.Sprintf("%c%c %s %c%c", o.char, o.frameLeft, text, o.frameRight, o.char)
	}
	bannerText += strings.Repeat(string(o.char), surroundingChars)

	if !o.noReset {
		bannerText += string(reset)
	}

	return bannerText
}

func NewBlock(header, block, footer string, opts ...Option) string {

	o := consolidateOpts(opts...)

	headerText := New(header, o, NoReset())

	maxLineLength := utf8.RuneCountInString(header)
	if o.length > maxLineLength {
		maxLineLength = o.length
	}

	// Parse the block one word at a time
	scanner := bufio.NewScanner(strings.NewReader(block))
	scanner.Split(bufio.ScanWords)

	var words []string
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	// Split the block into multiple lines with rune count up to maxLineLength
	var lines []string
	var wordsInLine []string
	lineLength := 0
	wordCount := 0
	// Note: This algorithm is not perfect. Blocks with multiple short words have too much
	//       white space at the end of each line, potentially due to the re-calculation of
	//       splitLineCount. This will be addressed in the near future.
	for _, w := range words {
		wordLen := utf8.RuneCountInString(w)
		wordCount += 1

		// If this is not the first word added to the line,
		// add 1 more to account for the space between
		// the current word and the previous word.
		if lineLength > 0 {
			lineLength += 1
		}
		lineLength += wordLen

		// If the addition of this word would exceed the maximum
		// line length, add the current wordsInLine to the
		// list of lines and reset the counters.
		if lineLength > maxLineLength {
			// If the maximum line length has been exceeded with
			// just 1 word, split this word 5 chars from the max line length
			// and add a - char.
			if wordCount == 1 {
				splitLineCount := int(math.Ceil(float64(wordLen) / float64(maxLineLength)))
				// Check if the splitLineCount would increase with the additional - chars
				splitLineCount = int(math.Ceil(float64(wordLen+splitLineCount) / float64(maxLineLength)))
				pivot := int(math.Min(float64(wordLen-1), float64(maxLineLength-1)))
				w1 := ""
				leftover := w
				for i := 0; i < splitLineCount-1; i++ {
					w1 = leftover[:pivot]
					w1 += "-"
					lines = append(lines, w1)
					leftover = leftover[pivot:]
				}
				wordCount = 1
				wordsInLine = []string{leftover}
				lineLength = utf8.RuneCountInString(leftover)
			} else {
				l := strings.Join(wordsInLine, " ")
				lines = append(lines, l)
				wordsInLine = []string{}
				lineLength = wordLen
			}
		} else {
			wordsInLine = append(wordsInLine, w)
		}
	}
	// If a full line was not completed, generate the final line.
	if len(wordsInLine) > 0 {
		lines = append(lines, strings.Join(wordsInLine, " "))
	}

	blockText := strings.Join(lines, "\n")

	footerFrame := Option{}
	if footer == "" {
		footerFrame = NoFrame()
	}
	footerText := New(footer, o, footerFrame)

	return fmt.Sprintf("%s\n%s\n%s", headerText, blockText, footerText)
}

func consolidateOpts(opts ...Option) Option {
	consolidated := Option{
		length:     defaultLength,
		char:       defaultChar,
		color:      defaultColor,
		frameLeft:  defaultFrameLeft,
		frameRight: defaultFrameRight,
	}
	for _, o := range opts {
		if o.length != 0 {
			consolidated.length = o.length
		}
		if o.char != 0 {
			consolidated.char = o.char
		}
		if o.color != "" {
			consolidated.color = o.color
		}
		if o.frameLeft != 0 {
			consolidated.frameLeft = o.frameLeft
		}
		if o.frameRight != 0 {
			consolidated.frameRight = o.frameRight
		}
		consolidated.noFrame = o.noFrame
		consolidated.noReset = o.noReset
	}
	return consolidated
}
