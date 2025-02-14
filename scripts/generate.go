package scripts

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	colors "github.com/echosonusharma/kata-machina/colors"
)

const (
	baseFolderName = "dsa"
	dsaNamePrefix  = "day"
	daySeparator   = "_"
)

var (
	dir_p *colors.ColorProfile = colors.New(
		colors.WithFgColor(colors.FgColorRegistry.Yellow),
		colors.WithBgColor(colors.BgColorRegistry.Black),
		colors.WithBoldText,
	)
)

func Generate(args []string) error {
	if err := checkBaseFolderExits(); err != nil {
		return err
	}

	nextDirNumber, err := getLastDirNumber()
	if err != nil {
		return err
	}

	nextDirNumber += 1

	nextDirPath := filepath.Join(baseFolderName, fmt.Sprintf("%s%s%d", dsaNamePrefix, daySeparator, nextDirNumber))
	if err := os.Mkdir(nextDirPath, 0750); err != nil {
		return err
	}

	fmt.Printf("%s dir created.\n", dir_p.Build(filepath.Base(nextDirPath)))
	return nil
}

// creates dsa dir if not exists
func checkBaseFolderExits() error {
	fi, err := os.Stat(baseFolderName)

	if err != nil {
		if os.IsNotExist(err) {
			return os.Mkdir(baseFolderName, 0750)
		}

		return err
	}

	if !fi.IsDir() {
		if err := os.Remove(baseFolderName); err != nil {
			return err
		}

		return os.Mkdir(baseFolderName, 0750)
	}

	return nil
}

// retrieves the last day's number
func getLastDirNumber() (int, error) {
	files, err := os.ReadDir(baseFolderName)
	if err != nil {
		return 0, err
	}

	var lastDay int

	for _, file := range files {
		if !file.IsDir() {
			continue
		}

		fName := strings.SplitN(file.Name(), daySeparator, 2)
		if len(fName) < 2 || fName[0] != dsaNamePrefix {
			continue
		}

		day, err := strconv.Atoi(fName[1])
		if err == nil && day > lastDay {
			lastDay = day
		}
	}

	return lastDay, nil
}
