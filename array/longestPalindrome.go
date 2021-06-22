package main

//最长的回文子串
//输入: "babad"
//输出: "bab"
//注意: "aba" 也是一个有效答案。

//力扣5 中心扩散法
func LongestPalindrome(s string) string {
	if s == "" {
		return ""
	}
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		left1, right1 := expandAroundCenter(s, i, i)   //奇数的回文，like  aba
		left2, right2 := expandAroundCenter(s, i, i+1) //偶数的回文，like  abba
		if right1-left1 > end-start {                  //求最大值
			start, end = left1, right1
		}
		if right2-left2 > end-start { //求最大值
			start, end = left2, right2
		}
	}
	return s[start : end+1]
}

func expandAroundCenter(s string, left, right int) (int, int) {
	for ; left >= 0 && right < len(s) && s[left] == s[right]; left, right = left-1, right+1 { //从中心往两边扩散
	}
	return left + 1, right - 1
}
