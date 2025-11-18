package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBasicValueType(t *testing.T) {

	digits := []int{2, 7, 1, 1}
	assert.Equal(t, []int{2, 7, 1, 2}, plusOne(digits))

	digits = []int{2, 7, 1, 1}
	assert.Equal(t, []int{2, 7, 1, 2}, plusOne2(digits))

	digits = []int{9, 9, 9}
	assert.Equal(t, []int{1, 0, 0, 0}, plusOne(digits))

	digits = []int{9, 9, 9}
	assert.Equal(t, []int{1, 0, 0, 0}, plusOne2(digits))

	digits = []int{1, 9}
	assert.Equal(t, []int{2, 0}, plusOne(digits))

	digits = []int{1, 9}
	assert.Equal(t, []int{2, 0}, plusOne2(digits))
}
