package utils

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type Util interface {
	LoadValues(path string) []string
	ScannerLoadValues(path string) []string
}

type UtilImpl struct{}

func GetUtil() Util {
	return &UtilImpl{}
}

// Returns slice of strings (lines) from the file
func (u *UtilImpl) LoadValues(path string) []string {
	var data []string
	tmp, readErr := os.ReadFile(filepath.Join(path))
	if readErr != nil {
		log.Fatal(readErr)
	}

	data = strings.Split(string(tmp), "\n")

	return data
}

// Returns slice of strings (lines) from the file
func (u *UtilImpl) ScannerLoadValues(path string) []string {
	file, err := os.Open(filepath.Join(path))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var data []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return data
}
