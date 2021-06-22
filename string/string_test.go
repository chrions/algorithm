package string

import (
	"testing"
)

func TestReverseString(t *testing.T) {
	s := "hello"
	b := []byte(s)
	reverseString(b)
	t.Log(string(b))
}

func TestReverseStr(t *testing.T) {
	s := "abcd"
	ret := reverseStr(s, 2)
	t.Log(ret)
}

func TestReverseWords(t *testing.T) {
	s := "  hello world  "
	ret := reverseWords(s)
	t.Log(ret)
}

func TestReplaceSpace(t *testing.T) {
	s := "We are happy."
	ret := replaceSpace(s)
	t.Log(ret)
}

func TestStrStr(t *testing.T) {
	haystack := "a"
	needle := "a"
	ret := strStr(haystack, needle)
	t.Log(ret)
}

func TestGetNexts(t *testing.T) {
	ret := getNexts("abdabcabc")
	t.Log(ret)
}

func TestStrStr2(t *testing.T) {
	haystack := "aabaabaafa"
	needle := "aabaaf"
	ret := strStr2(haystack, needle)
	t.Log(ret)
}

func TestRepeatedSubstringPattern(t *testing.T) {
	s := "abac"
	ret := repeatedSubstringPattern(s)
	t.Log(ret)
}

func TestIsNumber(t *testing.T) {
	s := ".-4"
	ret := isNumber(s)
	t.Log(ret)
}

func TestSmallestGoodBase(t *testing.T) {
	ret := smallestGoodBase("1000000000000000000")
	t.Log(ret)
}
