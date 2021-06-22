package main

import "strings"

//将一个给定字符串根据给定的行数，以从上往下、从左到右进行 Z 字形排列。
//
//比如输入字符串为 "LEETCODEISHIRING" 行数为 3 时，排列如下：
//
//L   C   I   R
//E T O E S I I G
//E   D   H   N
//之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："LCIRETOESIIGEDHN"。

//解题重点：多画几个图，推导出循环周期为 n = (2 * numRows - 2) （2倍行数 - 头尾2个）。
//对于字符串索引值 i，计算 x = i % n 确定在循环周期中的位置。
//行号rows[y]  y = min(x,n-x)
//当行数为1的时候，直接返回s
func Convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	rows := make([]string, numRows)
	n := numRows*2 - 2
	for i, v := range s {
		x := i % n
		rows[min(x, n-x)] += string(v)
	}
	return strings.Join(rows, "")
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
