package main

/*
给定一个表示 大整数 的整数数组 digits，其中 digits[i] 是整数的第 i 位数字。
这些数字按从左到右，从最高位到最低位排列。这个大整数不包含任何前导 0。

将大整数加 1，并返回结果的数字数组
*/

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
		}
		digits[i] = 0 // 当前位等于9，则置为0 进位
	}
	digits = append([]int{1}, digits...) // 数组头部添加一个1
	return digits
}

func plusOne2(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i]++
		if digits[i] < 10 {
			return digits
		}
		digits[i] = 0 // 当前位等于9，则置为0 进位
	}
	digits = append([]int{1}, digits...) // 数组头部添加一个1
	return digits

}
