package utils

import (
	"fmt"
	"os"
	"strings"
)

func GetFileContents(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("Error getting the contents of file %s: %w", path, err)
	}
	return string(data), nil
}

func GetLines(s string) []string {
	s = strings.ReplaceAll(s, "\r\n", "\n") // normalize Windows line endings
	s = strings.ReplaceAll(s, "\r", "\n")   // normalize old Mac line endings
	return strings.Split(s, "\n")
}

func GetFileLines(path string) ([]string, error) {
	s, err := GetFileContents(path)
	if err != nil {
		return nil, fmt.Errorf("Error getting the lines for file %s: %w", path, err)
	}
	return GetLines(s), nil
}
