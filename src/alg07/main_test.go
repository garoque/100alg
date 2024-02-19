package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateAreaSquare(t *testing.T) {
  assert.Equal(t, float64(4), calculateAreaSquare(2), "calculate area of 2")
  assert.Equal(t, float64(4), calculateAreaSquare(-2), "calculate area of -2")
}