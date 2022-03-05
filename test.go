package main

import (
	fu "fileutils"
	"fmt"
)

func main() {
	fmt.Println("main")
	files := fu.GetAllFiles()

	fmt.Println(files)
}
