package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateTrianguleArea(t *testing.T) {
  assert.Equal(t, float32(10), calculateTrianguleArea(10, 2), "returns 10 with this perimeter")
  assert.Equal(t, float32(6), calculateTrianguleArea(-2, 6), "returns 6 with negative value")
	assert.Equal(t, float32(4.375), calculateTrianguleArea(2.5, 3.5), "returns 4.375 with float number")
}