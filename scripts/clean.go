package scripts

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

func Clean() error {
	files, err := os.ReadDir(BASE_FOLDER_NAME)
	if err != nil {
		return err
	}

	re := regexp.MustCompile(`^` + DSA_NAME_PREFIX + `(\d+)$`)

	for _, file := range files {
		if file.IsDir() {
			matches := re.FindStringSubmatch(file.Name())
			if len(matches) == 2 {
				fmt.Println("deleting dir - ", file.Name())
				if err := os.RemoveAll(filepath.Join(BASE_FOLDER_NAME, file.Name())); err != nil {
					return err
				}
			} else {
				fmt.Println("ignoring dir - ", file.Name())
			}
		}
	}

	return nil
}
