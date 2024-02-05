package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertCelsiusToFahrenheit(t *testing.T) {
  assert.Equal(t, float32(96.8), convertCelsiusToFahrenheit(36), "convert 36 graus C return 96.8 graus F")
  assert.Equal(t, float32(59.9), convertCelsiusToFahrenheit(15.5), "convert 15.5 graus C return 59.9 graus F")
	assert.Equal(t, float32(14), convertCelsiusToFahrenheit(-10), "convert -10 graus C return 14 graus F")
}