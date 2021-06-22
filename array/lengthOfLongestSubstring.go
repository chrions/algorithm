package main

//给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
//输入: "cababcbb"
//cab ba  ab abc cb b
//输出: 3
//解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
//解法：滑动指针，左右指针从0开始，右指针遍历，遇到相同的值,左指针移动到之前那个值的后一位，所以要用map来存下标
func LengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	if len(s) == 1 {
		return 1
	}
	m := map[byte]int{}
	length := 0
	left := 0
	right := 0
	for ; left < len(s) && right < len(s); right++ {

		if v, ok := m[s[right]]; ok {
			if v >= left && v <= right {
				left = v + 1
			}
		}
		m[s[right]] = right
		length = Max(length, right-left+1)
	}
	return length
}

func Max(length1, length2 int) int {
	if length1 > length2 {
		return length1
	}
	return length2
}
