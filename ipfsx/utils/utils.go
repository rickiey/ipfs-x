package utils

import (
	"bufio"
	"io"
	"os"
)

// ReadFile reads the content of a file
func ReadFile(filePath string) ([]byte, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	return io.ReadAll(reader)
}

// FileExists checks if a file exists
func FileExists(filePath string) bool {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return false
	}
	return true
}
