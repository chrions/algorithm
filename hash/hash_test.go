package hash

import "testing"

func TestGroupAnagrams(t *testing.T) {
	a := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	ret := groupAnagrams(a)
	t.Log(ret)
}

func TestIsHappy(t *testing.T) {
	ret := isHappy(19)
	t.Log(ret)
}

func TestFourSumCount(t *testing.T) {
	A := []int{-1, -1}
	B := []int{-1, 1}
	C := []int{-1, 1}
	D := []int{1, -1}

	ret := fourSumCount(A, B, C, D)
	t.Log(ret)
}

func TestCanConstruct(t *testing.T) {
	ret := canConstruct("aaa", "aab")
	t.Log(ret)
}

func TestThreeSum(t *testing.T) {
	nums := []int{-1, 0, 1, 2, -1, -4, -2, -3, 3, 0, 4}
	ret := threeSum(nums)
	t.Log(ret)
}

func TestFourSum(t *testing.T) {
	nums := []int{1, 0, -1, 0, -2, 2}
	ret := fourSum(nums, 0)
	t.Log(ret)
}
