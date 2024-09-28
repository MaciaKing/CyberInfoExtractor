package models

import (
	"CyberInfoExtractor/cmd/globals"
	"bufio"
	"fmt"
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

func ReadFileFrom(filepath string, startLineToRead int,
	totalsLineToRead int) error {
	file, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	actualLineRead := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// if all lines are read
		if totalsLineToRead >= startLineToRead+actualLineRead {
			break
		}

		if actualLineRead >= startLineToRead {
			globals.LinesReads <- scanner.Text()
		}
		actualLineRead += 1
	}

	fmt.Println(totalsLineToRead - startLineToRead + actualLineRead)
	return nil
}
