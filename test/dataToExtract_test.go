package test

import (
	"CyberInfoExtractor/database"
	"CyberInfoExtractor/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDetectIp(t *testing.T) {
	assert.Equal(t, models.DetectIp("google.com"), 0)
	assert.Equal(t, models.DetectIp("http://8.8.8.8"), 0)
	assert.Equal(t, models.DetectIp("8.8.8.8"), 1)
}

func TestExtractData(t *testing.T) {
	database.Connect()
	database.Migrate()

	test1 := models.DataToExtract{}
	test1.ExtractData(database.DB, "google.com", false)

	// Find test1 on database
	search_test1 := models.DataToExtract{}
	if err := database.DB.Where("id = ?", test1.Id).First(&search_test1).Error; err != nil {
	}
	assert.Equal(t, search_test1, test1)
}
