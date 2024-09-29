package test

import (
	"CyberInfoExtractor/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDomainReport(t *testing.T) {
	result := models.DomainReport("ccc")
	assert.Contains(t, result, "error")

	result2 := models.DomainReport("8.8.8.8")
	assert.Contains(t, result2, "error")

	result3 := models.DomainReport("google.com")
	assert.Contains(t, result3, "data")
}
