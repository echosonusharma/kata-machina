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
	invalid_p *colors.ColorProfile = colors.New(
		colors.WithFgColor(colors.FgColorRegistry.Red),
		colors.WithBgColor(colors.BgColorRegistry.Black),
	)
	no_cmd_p *colors.ColorProfile = colors.New(
		colors.WithFgColor(colors.FgColorRegistry.Cyan),
		colors.WithBgColor(colors.BgColorRegistry.Black),
	)
)

var (
	invalid_cmd_msg string = fmt.Sprintf(`%s pls use %s for help`, invalid_p.Build("invalid command!"), help_p.Build("-h"))
	no_cmd_msg      string = fmt.Sprintf(`%s pls use %s for help`, no_cmd_p.Build("no command passed!"), help_p.Build("-h"))
)

func main() {
	args := os.Args

	if len(args) == 1 {
		fmt.Println(no_cmd_msg)
		return
	}

	if err := runCmd(args[1:]); err != nil {
		panic(err)
	}
}

func runCmd(args []string) error {
	cmd := args[0]
	cmd_args := args[1:]

	switch cmd {
	case "gen":
		return scripts.Generate(cmd_args)
	case "clean":
		return scripts.Clean(cmd_args)
	default:
		fmt.Println(invalid_cmd_msg)
		return nil
	}
}
