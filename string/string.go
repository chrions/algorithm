package string

import (
	"sync"
)

func reverseString(s []byte) {

	left := 0
	right := len(s) - 1
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}

//力扣541. 反转字符串 II,每隔k个反转k个，末尾不够k个时全部反转；
func reverseStr(s string, k int) string {
	s1 := []byte(s)
	sy := sync.WaitGroup{}
	cnt := 0
	if len(s1)%k == 0 {
		cnt = len(s1) / k
	} else {
		cnt = len(s1)/k + 1
	}

	sy.Add(cnt)
	flag := 0
	for i := 0; i < len(s1); i += k {
		end := i + k
		if end > len(s1) {
			end = len(s1)
		}
		flag++

		go func(start, end, flag int) {
			if flag%2 == 0 {
				sy.Done()
			} else {
				for start < end {
					s1[start], s1[end] = s1[end], s1[start]
					start++
					end--
				}

				sy.Done()
			}

		}(i, end-1, flag)
	}
	sy.Wait()
	return string(s1)
}

//力扣：剑指 Offer 05. 替换空格
func replaceSpace(s string) string {
	b := make([]byte, 0)

	for i := 0; i < len(s); i++ {
		if string(s[i]) == " " {
			b = append(b, []byte("%20")...)
		} else {
			b = append(b, s[i])
		}
	}
	return string(b)
}

//151. 翻转字符串里的单词
func reverseWords(s string) string {
	array := []byte(s)
	n := make([]byte, 0)
	n2 := make([][]byte, 0)
	s1 := ""
	for i := 0; i < len(array); i++ {
		if array[i] != 32 {
			n = append(n, array[i])
		} else if len(n) > 0 {
			n2 = append(n2, n)
			n = make([]byte, 0)
		}

	}
	if len(n) > 0 {
		n2 = append(n2, n)
	}

	for i := len(n2) - 1; i >= 0; i-- {
		s1 += string(n2[i]) + " "
	}
	return s1[:len(s1)-1]
}

// 28. 实现 strStr()
//暴力解法
func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}

	n := len(needle)
	for i := 0; i <= len(haystack)-n; i++ {
		stack := make([]byte, 0)
		stack = append(stack, haystack[i:i+n]...)
		if string(stack) == needle {
			return i
		}
	}
	return -1
}

// 28. 实现 strStr()
//KMP 算法
func strStr2(haystack string, needle string) int {
	n := len(haystack)
	m := len(needle)
	if m == 0 {
		return 0
	}
	nexts := getNexts(needle)
	j := -1 //next数组里记录的起始位置为-1
	for i := 0; i < n; i++ {
		for haystack[i] != needle[j+1] && j >= 0 {
			j = nexts[j] //不匹配时，从nexts数组中找到下一个匹配位置
		}
		if haystack[i] == needle[j+1] {
			j++ //匹配时，i和j同时向后移动
		}
		if j == m-1 { //如果j指向了模式串needle的末尾，则完全匹配
			return i - m + 1
		}
	}

	return -1
}

func getNexts(pattern string) []int {
	m := len(pattern)
	nexts := make([]int, m)
	// 定义两个指针i和j，j指向前缀终止位置（严格来说是终止位置减一的位置），i指向后缀终止位置（与j同理）。
	j := -1
	nexts[0] = j
	for i := 1; i < m; i++ { // i从1往后遍历
		// j从0往后遍历，当j+1和i不相等时，就是遇到了 前后缀末尾不相同的情况，此时要进行回溯
		for pattern[j+1] != pattern[i] && j >= 0 {
			//next[j]就是记录着j（包括j）之前的子串的相同前后缀的长度。
			//s[i] 与 s[j+1] 不相同，就要找 j+1前一个元素在next数组里的值，就是 nexts[j]
			j = nexts[j]
		}

		if pattern[j+1] == pattern[i] {
			//同时向后移动i 和j 说明找到了相同的前后缀
			j += 1
		}
		//还要将j（前缀的长度）赋给next[i]，next[i]要记录相同前后缀的长度
		nexts[i] = j
	}

	return nexts
}

// 459. 重复的子字符串
func repeatedSubstringPattern(s string) bool {
	if len(s) == 0 {
		return false
	}
	nexts := getNexts(s)
	n := len(s) - (nexts[len(nexts)-1] + 1)
	//当 nexts 数组 nexts[len(nexts)-1] != -1 时
	//nexts[len(nexts)-1] + 1 是字符串的 最长相同前后缀 的长度
	//n 如果能被12整除就说明有重复的子字符串
	if len(s)%n == 0 && nexts[len(nexts)-1] != -1 {
		return true
	}
	return false
}

//65. 有效数字
func isNumber(s string) bool {
	if len(s) == 0 {
		return false
	}
	isNum := false
	isE_or_e := false
	isDot := false
	for i := 0; i < len(s); i++ {
		if s[i] == 'e' || s[i] == 'E' {
			if !isNum || isE_or_e {
				return false
			}
			isNum = false
			isE_or_e = true
		} else if s[i] == '.' {
			if isDot || isE_or_e {
				return false
			}
			isDot = true
		} else if s[i] == '+' || s[i] == '-' {
			if i > 0 && (s[i-1] == '+' || s[i-1] == '-' || (s[i-1] >= '0' && s[i-1] <= '9') || s[i-1] == '.') {
				return false
			}

		} else if s[i] >= '0' && s[i] <= '9' {
			isNum = true
		} else {
			return false
		}
	}
	return isNum
}

//力扣483
func smallestGoodBase(n string) string {

	return ""
}
