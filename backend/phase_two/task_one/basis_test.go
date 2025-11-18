package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBasic(t *testing.T) {

	//声明整数数组 nums 和一个整数目标值 target
	nums := []int{2, 7, 11, 15}
	target := 9
	assert.Equal(t, twoSum(nums, target), []int{0, 1})
	assert.Equal(t, twoSum2(nums, target), []int{0, 1})

	nums = []int{3, 2, 4, 3, 4}
	target = 6
	assert.Subset(t, twoSum(nums, target), []int{1, 2, 0, 3})
	assert.Subset(t, twoSum2(nums, target), []int{0, 3, 1, 2})
}
