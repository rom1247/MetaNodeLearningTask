package main

import (
	"strconv"
	"strings"
)

// 声明一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次
func singleNumber(nums []int) int { //声明一个函数，参数为整数切到，即动态数组，返回一个整数
	var res int                //定义一个整数变量
	for _, num := range nums { //用range遍历数组nums，_是忽略索引，只获取值num
		res ^= num //进行异或运算  相同数字会变为0，不同数字会异或为结果
	}
	return res
}

// 使用map计数实现求值，除了某个元素只出现一次以外，其余每个元素均出现两次
func singleNumber2(nums []int) int {

	var res int            //定义一个整数变量
	m := make(map[int]int) //定义一个map变量
	for _, num := range nums {
		m[num]++ //统计数组中每个数字出现的次数
	}
	for k, v := range m {
		if v == 1 { //判断次数为1的数字
			res = k
		}
	}
	return res
}

// 优化，求解一个数组中只出现一次的数字，返回一个新数组
func singleNumber3(nums []int) []int {
	var res []int //定义一个切片变量
	m := make(map[int]int)
	for _, num := range nums {
		m[num]++
	}
	for k, v := range m {
		if v == 1 {
			res = append(res, k) //添加到切片中
		}
	}
	return res
}

// 声明一个函数，判断一个数字是否是回文数
func isPalindrome(x int) bool {
	if x < 0 { //判断是否小于0
		return false
	}
	var res int   //定义一个整数变量
	for x > res { //核心逻辑是取一半数字判断
		res = res*10 + x%10          // res是保存从个位数翻转的数字 ，x%10取模，获取不能整除的数字
		x /= 10                      // res每翻转一次，x去除一位
		if x == res || x/10 == res { //判断偶数翻转后的数字是否相等，或者奇数翻转后的数字是否相等，x/10意思是去除中间的那个数字
			return true
		}
	}
	return false
}

// 优化，判断一个数字是否是回文数，通过字符串判断
func isPalindrome2(x int) bool {
	if x < 0 { //判断是否小于0
		return false
	}
	s := strconv.Itoa(x)
	//声明一个反转字符串变量
	var res string
	for i := len(s) - 1; i >= 0; i-- {
		res += string(s[i])
	}
	return res == s
}

// 优化，判断一个数字是否是回文数，直接判断字符串首位两端的字符是否相同
func isPalindrome3(x int) bool {
	if x < 0 { //判断是否小于0
		return false
	}
	s := strconv.Itoa(x)
	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func main() {
	nums := []int{2, 2, 1}
	res := singleNumber(nums)
	println(res)

	nums = []int{4, 1, 2, 1, 2}
	res = singleNumber2(nums)
	println(res)

	nums = []int{1, 2, 1, 3, 2, 5}
	res3 := singleNumber3(nums)

	//strings只操作字符串类型 所以先转一下
	res4 := make([]string, len(res3))
	for i, v := range res3 {
		res4[i] = strconv.Itoa(v)
	}
	res5 := strings.Join(res4, ",")
	println(res5)

	x := 1234321
	b := isPalindrome(x)
	println(b)

	x = 1234321
	b = isPalindrome2(x)
	println(b)

	x = 1234321
	b = isPalindrome3(x)
	println(b)

}
