package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateAverage(t *testing.T) {
  assert.Equal(t, float32(4), calculateAverage([]float32{2, 4, 6}), "returns 4 with result")
  assert.NotEqual(t, float32(7), calculateAverage([]float32{2, 4, 6}), "returns false for 2 + 4 + 6 / 3 equals 7")
}