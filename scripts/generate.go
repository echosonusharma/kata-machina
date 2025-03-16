package scripts

import (
	"embed"
	"encoding/json"
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
	file_p *colors.ColorProfile = colors.New(
		colors.WithFgColor(colors.FgColorRegistry.Magenta),
		colors.WithBgColor(colors.BgColorRegistry.Black),
		colors.WithBoldText,
	)
)

//go:embed static/*
var configFiles embed.FS

type dsaCode struct {
	Name   string   `json:"name"`
	Import string   `json:"import"`
	Func   []string `json:"func"`
}

type dsaConfig struct {
	Dsa  []string  `json:"dsa"`
	Code []dsaCode `json:"code"`
}

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

	data, err := configFiles.ReadFile("static/dsa.config.json")
	if err != nil {
		return err
	}

	var dsaConfig dsaConfig

	err = json.Unmarshal(data, &dsaConfig)
	if err != nil {
		return err
	}

	for _, v := range dsaConfig.Code {
		fileName := fmt.Sprintf("%s.go", v.Name)
		var fileContent string
		fileContent += fmt.Sprintf("package %s", fmt.Sprintf("%s%d", dsaNamePrefix, nextDirNumber))
		fileContent += "\n\n"
		if len(v.Import) > 0 {
			fileContent += v.Import
			fileContent += "\n\n"
		}
		fileContent += strings.Join(v.Func, "\n\n")

		err := os.WriteFile(filepath.Join(nextDirPath, fileName), []byte(fileContent), 0644)
		if err != nil {
			return err
		}

		fmt.Printf("%s file created.\n", file_p.Build(fileName))
	}

	fmt.Printf("all done %s", dir_p.Build("^_____^"))
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
