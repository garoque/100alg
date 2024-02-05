package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateRectanglePerimeter(t *testing.T) {
  assert.Equal(t, float32(36), calculateRectanglePerimeter(11, 7), "returns 36 with this perimeter")
  assert.Equal(t, float32(36), calculateRectanglePerimeter(-11, 7), "returns 36 with negative value")
}