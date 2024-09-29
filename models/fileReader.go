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

func ReadFileFrom(filepath string, startLineToRead int, totalsLineToRead int) (int, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return -1, err
	}
	defer file.Close()

	currentLine := 0
	linesRead := 0

	scanner := bufio.NewScanner(file)
	// Read file
	for scanner.Scan() {
		currentLine++
		//  Start reading from the startLineToRead line.
		if currentLine < startLineToRead {
			continue
		}
		globals.LinesReads <- scanner.Text()
		linesRead++
		// If we have already read the total number of lines requested, we are done.
		if linesRead >= totalsLineToRead {
			break
		}
	}

	// If we have not read enough lines but have reached the end of the file,
	// we simply report that we have reached the end and there are no
	// more lines to read. We return the number of lines left to read.
	if linesRead < totalsLineToRead {
		return totalsLineToRead - linesRead, nil
	}

	return 0, nil
}
