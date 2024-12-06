package dir_processing

import (
	"io"
	"os"
	"path/filepath"
)

func ReadFile(path string) []byte {
	if _, err := os.Stat(path); err != nil {
		return nil
	}

	file, err := os.Open(path)
	if err != nil {
		return nil
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil
	}
	return data
}

func ReadDir(path string) []string {
	files, _ := os.ReadDir(path)
	absPath, _ := filepath.Abs(path)
	fullPaths := []string{}
	for _, file := range files {
		fullPath := filepath.Join(absPath, file.Name())
		fullPaths = append(fullPaths, fullPath)
	}
	return fullPaths
}

func CheckIsDir(path string) (bool, error) {
	info, err := os.Stat(path)
	return info.IsDir(), err
}
