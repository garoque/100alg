package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddTwoNumbers(t *testing.T) {
  assert.Equal(t, 4, addTwoNumbers(2, 2), "returns 4 with result")
  assert.Equal(t, 5, addTwoNumbers(15, -10), "returns 5 with result")

  assert.NotEqual(t, 7, addTwoNumbers(2, 2), "returns false for 2 + 2 equals 7")
}