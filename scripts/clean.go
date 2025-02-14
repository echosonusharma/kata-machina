package scripts

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	colors "github.com/echosonusharma/kata-machina/colors"
)

var (
	ignored_dir_p *colors.ColorProfile = colors.New(
		colors.WithFgColor(colors.FgColorRegistry.BrightCyan),
	)
	del_dir_p *colors.ColorProfile = colors.New(
		colors.WithFgColor(colors.FgColorRegistry.BrightRed),
	)
	no_action_p *colors.ColorProfile = colors.New(
		colors.WithFgColor(colors.FgColorRegistry.BrightCyan),
	)
	warn_p *colors.ColorProfile = colors.New(
		colors.WithFgColor(colors.FgColorRegistry.Yellow),
	)
)

func Clean(args []string) error {
	deleteAll := false
	deleteLast := false

	if len(args) > 0 {
		arg := args[0]

		if arg == "-all" {
			warn_p.Print("deleting all the dsa directory's...\n")
			deleteAll = true
		} else if arg == "-last" {
			warn_p.Print("deleting the last dsa directory...\n")
			deleteLast = true
		}
	}

	files, err := os.ReadDir(baseFolderName)
	if err != nil {
		return err
	}

	if deleteAll {
		for _, file := range files {
			if file.IsDir() {
				nameCheck := strings.HasPrefix(file.Name(), fmt.Sprintf("%s%s", dsaNamePrefix, daySeparator))

				if nameCheck {
					fmt.Printf("%s dir is deleted\n", del_dir_p.Build(file.Name()))
					if err := os.RemoveAll(filepath.Join(baseFolderName, file.Name())); err != nil {
						return err
					}
				} else {
					fmt.Printf("%s dir is ignored\n", ignored_dir_p.Build(file.Name()))
				}
			}
		}

		return nil
	}

	if deleteLast {
		lastDirNumber, err := getLastDirNumber()
		if err != nil {
			return nil
		}

		if lastDirNumber == 0 {
			return nil
		}

		dirName := fmt.Sprintf("%s%s%d", dsaNamePrefix, daySeparator, lastDirNumber)
		fmt.Printf("%s dir is deleted\n", del_dir_p.Build(dirName))
		if err := os.RemoveAll(filepath.Join(baseFolderName, dirName)); err != nil {
			return err
		}

		return nil
	}

	no_action_p.Print("no action taken! pls check your cmd\n")
	return nil
}
