package main

import (
	"fmt"
	"os"

	"github.com/TrevorEdris/banner"
	"github.com/spf13/cobra"
)

const (
	flagChar       = "char"
	flagLength     = "length"
	flagFrameLeft  = "frame-left"
	flagFrameRight = "frame-right"
	flagColor      = "color"
)

func main() {
	cmd := &cobra.Command{
		Use:   "banner 'Surround this text'",
		Short: "Generate a banner of text",
		Args: func(cmd *cobra.Command, args []string) error {
			if err := cobra.MinimumNArgs(1)(cmd, args); err != nil {
				return err
			}
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			text := args[0]
			length, err := cmd.Flags().GetInt(flagLength)
			printAndExit(err)
			fchar, err := cmd.Flags().GetString(flagChar)
			printAndExit(err)
			fframeLeft, err := cmd.Flags().GetString(flagFrameLeft)
			printAndExit(err)
			fframeRight, err := cmd.Flags().GetString(flagFrameRight)
			printAndExit(err)
			fcolor, err := cmd.Flags().GetString(flagColor)
			printAndExit(err)

			char, err := parseChar(fchar)
			printAndExit(err)
			frameLeft, err := parseChar(fframeLeft)
			printAndExit(err)
			frameRight, err := parseChar(fframeRight)
			printAndExit(err)
			colorFunc, err := parseColor(fcolor)
			printAndExit(err)

			fmt.Println(banner.New(
				text,
				banner.WithChar(char),
				banner.WithLength(length),
				banner.WithFrame(frameLeft, frameRight),
				colorFunc(),
			))
		},
	}
	cmd.Flags().Int(flagLength, 80, "The total length of the banner")
	cmd.Flags().String(flagChar, "=", "The characer to surround the text with")
	cmd.Flags().String(flagFrameLeft, "[", "The left framing character")
	cmd.Flags().String(flagFrameRight, "]", "The right framing character")
	cmd.Flags().String(flagColor, "", fmt.Sprintf("The color of the banner (none) oneOf %v", banner.AvailableColors()))

	err := cmd.Execute()
	printAndExit(err)
}

func printAndExit(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func parseChar(s string) (rune, error) {
	if len(s) > 1 {
		return 0, fmt.Errorf("unable to use %s as surrounding character; must be exactly 1 character in length", s)
	}
	if s == "" {
		return 0, nil
	}
	return []rune(s)[0], nil
}

func parseColor(c string) (banner.ColorFunc, error) {
	f, err := banner.GetColorFunc(c)
	if err != nil {
		return f, fmt.Errorf("%w: available colors %v", err, banner.AvailableColors())
	}
	return f, nil
}
