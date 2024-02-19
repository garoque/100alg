package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConventMetersToCentimeters(t *testing.T) {
  assert.Equal(t, float32(1000), conventMetersToCentimeters(10), "returns 1000 with this meter")
}