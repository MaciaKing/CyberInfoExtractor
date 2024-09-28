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

	expected := []string{"line1", "line2", "line3"}
	for _, exp := range expected {
		got := <-globals.LinesReads
		assert.Equal(t, exp, got)
	}
	close(globals.LinesReads)
}
