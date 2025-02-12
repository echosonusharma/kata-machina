package main

import (
	"fmt"
	"os"

	colors "github.com/echosonusharma/kata-machina/colors"
	"github.com/echosonusharma/kata-machina/scripts"
)

var (
	help_p *colors.ColorProfile = colors.New(
		colors.WithFgColor(colors.FgColorRegistry.Yellow),
		colors.WithBoldText,
	)
	fail_p *colors.ColorProfile = colors.New(
		colors.WithFgColor(colors.FgColorRegistry.BrightRed),
	)
)

var help_msg string = fmt.Sprintf(`%s pls use %s for help`, fail_p.Build("invalid command!"), help_p.Build("-h"))

func main() {
	args := os.Args

	if err := cli(args[1]); err != nil {
		panic(err)
	}
}

func cli(arg string) error {
	switch arg {
	case "gen":
		return scripts.Generate()
	case "clean":
		return scripts.Clean()
	default:
		fmt.Println(help_msg)
		return nil
	}
}
