package main

import (
	"fmt"
	"os"

	"github.com/echosonusharma/kata-machina/scripts"
)

func main() {
	args := os.Args
	allArgs := args[1:]

	// fmt.Println("cli flags passed", allArgs)

	if err := cli(allArgs[0]); err != nil {
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
		fmt.Printf("%s\n", "argument not found ^_^")
		return nil
	}
}
