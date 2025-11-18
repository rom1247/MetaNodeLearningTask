package main

import "slices"

/*
给定一个整数数组 nums 和一个整数目标值 target，
请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。
每个数组的数字只能匹配一次。
*/

func twoSum(nums []int, target int) []int {
	var res []int
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				if slices.Contains(res, i) || slices.Contains(res, j) { // 去重
					continue
				}
				res = append(res, i, j)
			}
		}
	}
	return res
}

func twoSum2(nums []int, target int) []int {
	var res []int
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if v, ok := m[target-nums[i]]; ok { //  map取值判断是否存在，ok为true则存在
			if slices.Contains(res, i) || slices.Contains(res, v) {
				continue
			}
			res = append(res, v, i) // 存在则返回索引
		}
		m[nums[i]] = i //保存数组元素和索引 ，用于m[target-nums[i]] 快速匹配
	}
	return res
}
