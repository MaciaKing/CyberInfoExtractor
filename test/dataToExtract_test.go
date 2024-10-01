package test

import (
	"CyberInfoExtractor/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDetectIp(t *testing.T) {
	assert.Equal(t, models.DetectIp("google.com"), 0)
	assert.Equal(t, models.DetectIp("http://8.8.8.8"), 0)
	assert.Equal(t, models.DetectIp("8.8.8.8"), 1)
}
