package test

import (
	"CyberInfoExtractor/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDomainReport(t *testing.T) {
	vt := models.VirusTotal{}
	result := vt.DomainReport("ccc")
	assert.Contains(t, result, "error")
	assert.Contains(t, vt.InformationExtracted, "error")

	vt2 := models.VirusTotal{}
	result2 := vt2.DomainReport("8.8.8.8")
	assert.Contains(t, result2, "error")
	assert.Contains(t, vt2.InformationExtracted, "error")

	vt3 := models.VirusTotal{}
	result3 := vt3.DomainReport("google.com")
	assert.Contains(t, result3, "data")
	assert.Contains(t, vt3.InformationExtracted, "data")
}
