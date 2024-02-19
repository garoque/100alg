package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateVolumeSphere(t *testing.T) {
  assert.Equal(t, float64(65.45), calculateVolumeSphere(2.5), "calculate volume of 2.5")
}