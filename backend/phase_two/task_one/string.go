package main

/*
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
每个右括号都有一个对应的相同类型的左括号。
*/

func isValid(s string) bool {
	var stack []rune
	for _, v := range s {
		if v == '(' || v == '[' || v == '{' {
			stack = append(stack, v)
		} else {
			if len(stack) == 0 {
				return false
			}
			if v == ')' && stack[len(stack)-1] != '(' {
				return false
			}
			if v == ']' && stack[len(stack)-1] != '[' {
				return false
			}
			if v == '}' && stack[len(stack)-1] != '{' {
				return false
			}
			stack = stack[:len(stack)-1] // 弹出栈顶元素，相当于java的stack.pop()
		}
	}
	return len(stack) == 0
}

// 字符串数组中的最长公共前缀
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	var res string
	for i := 0; i < len(strs[0]); i++ {
		for j := 1; j < len(strs); j++ {
			if i >= len(strs[j]) || strs[j][i] != strs[0][i] {
				return res
			}
		}
		res += string(strs[0][i]) //累加相同的字符
	}
	return res
}

// 字符串数组中的最长公共前缀
func longestCommonPrefix2(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	var res = strs[0]
	for i := 1; i < len(strs); i++ {
		for j := 0; j < len(strs[i]); j++ {
			if len(res) < j {
				break
			}
			if res[j] != strs[i][j] {
				if j == 0 {
					return ""
				} else {
					res = res[:j-1] // 截取相同的字符串
				}

			}
		}

	}

	return res
}
