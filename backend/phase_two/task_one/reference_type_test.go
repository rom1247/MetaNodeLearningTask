package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReferenceType(t *testing.T) {

	//非严格递增数组
	arr := []int{1, 1, 1, 2, 2, 3, 3, 4, 5, 6, 7, 8, 8, 9, 10}

	//严格递增数组
	arr2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	l := removeDuplicates(arr)
	assert.Equal(t, len(arr2), l)
	for i, v := range arr2 {
		assert.Equal(t, v, arr[i])
	}
}

func TestMerge(t *testing.T) {
	intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	res := merge(intervals)
	assert.Equal(t, [][]int{{1, 6}, {8, 10}, {15, 18}}, res)
}
