package fileutils

import (
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
)

func GetAllFiles() []string {
	var files []string
	var dir string

	godotenv.Load("../.env")
	// Directory for testing the program, try to change it to your own directory
	dir = os.Getenv("DIR")

	if runtime.GOOS == "windows" && dir == "" {
		dir = "C:\\Users\\"
	} else if dir == "" {
		dir = "/home/"
	}

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
