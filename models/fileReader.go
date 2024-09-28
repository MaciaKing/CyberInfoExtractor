package models

import (
	"CyberInfoExtractor/cmd/globals"
	"bufio"
	"os"
)

type FileReader struct {
}

func ReadFile(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		globals.LinesReads <- scanner.Text()
	}
	file.Close()
	return nil
}
