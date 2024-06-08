package scripts

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
)

const (
	BASE_FOLDER_NAME = "dsa"
	DSA_NAME_PREFIX  = "day"
)

// make sure main folder exits
// find the next day count
// create the file

func Generate() error {
	if err := checkBaseFolderExits(); err != nil {
		return err
	}

	nextDirNumber, err := getNextDirNumber()
	if err != nil {
		return nil
	}

	nextDirPath := filepath.Join(BASE_FOLDER_NAME, fmt.Sprintf("%s%d", DSA_NAME_PREFIX, nextDirNumber))
	if err := os.Mkdir(nextDirPath, 0750); err != nil {
		return err
	}

	fmt.Printf("%s dir created.\n", filepath.Base(nextDirPath))

	return nil
}

func checkBaseFolderExits() error {
	fi, err := os.Stat(BASE_FOLDER_NAME)

	if err != nil {
		if os.IsNotExist(err) {
			return os.Mkdir(BASE_FOLDER_NAME, 0750)
		}

		return err
	}

	if !fi.IsDir() {
		if err := os.Remove(BASE_FOLDER_NAME); err != nil {
			return err
		}

		return os.Mkdir(BASE_FOLDER_NAME, 0750)
	}

	return nil
}

func getNextDirNumber() (int, error) {
	files, err := os.ReadDir(BASE_FOLDER_NAME)
	if err != nil {
		return 0, err
	}

	var maxDay int
	re := regexp.MustCompile(`^` + DSA_NAME_PREFIX + `(\d+)$`)

	for _, file := range files {
		if file.IsDir() {
			matches := re.FindStringSubmatch(file.Name())
			if len(matches) == 2 {
				day, err := strconv.Atoi(matches[1])
				if err == nil && day > maxDay {
					maxDay = day
				}
			}
		}
	}

	return maxDay + 1, nil
}
