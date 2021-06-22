package main

import (
	"strings"
)

//给定一个以字符串表示的非负整数 num，移除这个数中的 k 位数字，使得剩下的数字最小。
//
//注意:
//
//num 的长度小于 10002 且 ≥ k。
//num 不会包含任何前导零。
//示例 1 :
//
//输入: num = "1432219", k = 3
//输出: "1219"
//解释: 移除掉三个数字 4, 3, 和 2 形成一个新的最小的数字 1219。
//力扣 402
func RemoveKdigits(num string, k int) string {
	var stack []uint8 //作为栈使用
	var ret string
	for i := 0; i < len(num); i++ {
		number := num[i] - '0' //转化为数字
		for len(stack) != 0 && stack[len(stack)-1] > number && k > 0 {
			stack = stack[:len(stack)-1] //从栈顶出栈
			k--
		}

		if number != 0 || len(stack) != 0 {
			stack = append(stack, number)
		}

	}
	for len(stack) > 0 && k > 0 { // 对于 12345 的这种，最终栈里面就是12345，K不变，直接删除栈的K 后位
		stack = stack[:len(stack)-1]
		k--
	}
	for _, v := range stack {
		ret += string(v + '0')
	}

	if ret == "" { // 对于 20 ， k=2 的这种，最终栈里面是空的
		ret += string('0')
	}

	return ret
}

func RemoveKdigits2(num string, k int) string {
	var stack []byte //byte之间可以直接比较大小
	for i := 0; i < len(num); i++ {
		number := num[i]
		for len(stack) != 0 && stack[len(stack)-1] > number && k > 0 {
			stack = stack[:len(stack)-1]
			k--
		}
		stack = append(stack, number)

	}
	stack = stack[:len(stack)-k]

	s := string(stack)
	s = strings.TrimLeft(s, "0") // 去掉对于 000000xxx这种格式的
	if s == "" {
		return "0"
	}
	return s
}
