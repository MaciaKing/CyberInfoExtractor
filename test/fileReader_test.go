package test

import (
	"CyberInfoExtractor/cmd/globals"
	"CyberInfoExtractor/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadFile(t *testing.T) {
	wrongFilePath := "random"
	err := models.ReadFile(wrongFilePath)
	assert.Error(t, err)

	filePath := "test.txt"
	err1 := models.ReadFile(filePath)
	assert.NoError(t, err1)

	expected := []string{"line0", "line1", "line2"}
	for _, exp := range expected {
		got := <-globals.LinesReads
		assert.Equal(t, exp, got)
	}

	cleanChan(globals.LinesReads)
}

func TestReadFileFrom(t *testing.T) {
	wrongFilePath := "random"
	err := models.ReadFileFrom(wrongFilePath, 1, 2)
	assert.Error(t, err)

	filePath := "test.txt"
	err1 := models.ReadFileFrom(filePath, 6, 3)
	assert.NoError(t, err1)

	expected := []string{"line6", "line7", "line8"}
	for _, exp := range expected {
		got := <-globals.LinesReads
		assert.Equal(t, exp, got)
	}

	err2 := models.ReadFileFrom(filePath, 9, 1)
	assert.NoError(t, err2)

	expected2 := []string{"line9"}
	for _, exp := range expected2 {
		got := <-globals.LinesReads
		assert.Equal(t, exp, got)
	}

	err3 := models.ReadFileFrom(filePath, 9, 2)
	assert.NoError(t, err3)

	expected3 := []string{"line9"}
	for _, exp := range expected3 {
		got := <-globals.LinesReads
		assert.Equal(t, exp, got)
	}

	close(globals.LinesReads)
	cleanChan(globals.LinesReads)
}

func cleanChan(chanToClean chan string) {
	for len(chanToClean) > 0 {
		<-globals.LinesReads
	}
}
