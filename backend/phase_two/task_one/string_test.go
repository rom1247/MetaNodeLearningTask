package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {

	var str = "()[]{}"
	assert.Equal(t, true, isValid(str))

	str = "([)]"
	assert.Equal(t, false, isValid(str))

	str = "{[]}"
	assert.Equal(t, true, isValid(str))

}

func TestLongestCommonPrefix(t *testing.T) {

	var strs = []string{"flower", "flow", "flight"}
	/*	assert.Equal(t, "fl", longestCommonPrefix(strs))
		assert.Equal(t, "fl", longestCommonPrefix2(strs))
	*/
	strs = []string{"dog", "racecar", "car"}
	assert.Equal(t, "", longestCommonPrefix(strs))
	assert.Equal(t, "", longestCommonPrefix2(strs))

	strs = []string{}
	assert.Equal(t, "", longestCommonPrefix(strs))
	assert.Equal(t, "", longestCommonPrefix2(strs))

}
