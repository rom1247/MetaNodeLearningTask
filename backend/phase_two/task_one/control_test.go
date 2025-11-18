package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSingleNumber(t *testing.T) {

	nums := []int{2, 2, 1}
	res := singleNumber(nums)
	assert.Equal(t, 1, res)

	nums = []int{4, 1, 2, 1, 2}
	res = singleNumber2(nums)
	assert.Equal(t, 4, res)

	nums = []int{4, 1, 2, 1, 2, 5}
	number3 := singleNumber3(nums)
	assert.Equal(t, []int{4, 5}, number3)

}

func TestIsPalindrome(t *testing.T) {

	x := 121
	res := isPalindrome(x)
	assert.Equal(t, true, res)

	x = 123
	res = isPalindrome2(x)
	assert.Equal(t, false, res)

	x = 12321
	res = isPalindrome3(x)
	assert.Equal(t, true, res)

}
