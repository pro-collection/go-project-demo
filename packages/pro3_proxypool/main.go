package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

// execPath returns the executable path.
func execPath() (string, error) {
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		return "", err
	}
	return filepath.Abs(file)
}

func main() {
	filePath, err := execPath()

	if err != nil {
		return
	}

	fmt.Println("path： ", filePath)
}
