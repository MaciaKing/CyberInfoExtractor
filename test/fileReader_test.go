package test

import (
	"CyberInfoExtractor/cmd/globals"
	"CyberInfoExtractor/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadFile(t *testing.T) {
	fr := models.FileReader{}

	wrongFilePath := "random"
	err := fr.ReadFile(wrongFilePath)
	assert.Error(t, err)

	filePath := "test.txt"
	err1 := fr.ReadFile(filePath)
	assert.NoError(t, err1)

	expected := []string{"line0", "line1", "line2"}
	for _, exp := range expected {
		got := <-globals.LinesReads
		assert.Equal(t, exp, got)
	}

	cleanChan(globals.LinesReads)
}

func TestReadFileFrom(t *testing.T) {
	fr := models.FileReader{}

	wrongFilePath := "random"
	linesNotReaded, err := fr.ReadFileFrom(wrongFilePath, 1, 2)
	assert.Error(t, err)
	assert.Equal(t, linesNotReaded, -1)

	filePath := "test.txt"
	linesNotReaded1, err1 := fr.ReadFileFrom(filePath, 6, 3)
	assert.NoError(t, err1)
	assert.Equal(t, linesNotReaded1, 0)

	expected := []string{"line5", "line6", "line7"}
	for _, exp := range expected {
		got := <-globals.LinesReads
		assert.Equal(t, exp, got)
	}

	linesNotReaded2, err2 := fr.ReadFileFrom(filePath, 9, 1)
	assert.NoError(t, err2)
	assert.Equal(t, linesNotReaded2, 0)

	expected2 := []string{"line8"}
	for _, exp := range expected2 {
		got := <-globals.LinesReads
		assert.Equal(t, exp, got)
	}

	linesNotReaded3, err3 := fr.ReadFileFrom(filePath, 9, 3)
	assert.NoError(t, err3)
	assert.Equal(t, linesNotReaded3, 1)

	expected3 := []string{"line8", "line9"}
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
