package main

import (
	"testing"
)

func TestTwoSum(t *testing.T) {
	nums := []int{2, 11, 7, 15}
	target := 9
	res := TwoSum(nums, target)
	t.Log(res)
}

func TestLengthOfLongestSubstring(t *testing.T) {
	s := "pwwkew"
	ret := LengthOfLongestSubstring(s)
	t.Log(ret)
}

func TestGetMergenums(t *testing.T) {
	m1 := []int{2, 4, 7, 9}
	m2 := []int{1, 3, 6, 8}

	ret := GetMergeSortnums(m1, m2, 4, 4)
	t.Log(ret)
}

func TestFindMedianSortedArrays(t *testing.T) {
	m1 := []int{3}
	m2 := []int{1, 2, 4}
	ret := FindMedianSortedArrays(m1, m2)
	t.Log(ret)
}

func TestLongestPalindrome(t *testing.T) {
	s := "babad"
	ret := LongestPalindrome(s)
	t.Log(ret)
}

func TestThreeSum(t *testing.T) {
	nums := []int{-1, 0, 1, 2, -1, -4, -2, -3, 3, 0, 4}
	ret := threeSum(nums)
	t.Log(ret)
}

//1,1,2,3,5,8,13
func TestFibo(t *testing.T) {
	ret := Fibo(7)
	t.Log(ret)
}

func TestNumWays(t *testing.T) {
	ret := NumWays(100)
	t.Log(ret)
}

func TestZConvert(t *testing.T) {
	ret := Convert("AB", 1)
	t.Log(ret)
}

func TestRemoveKdigits(t *testing.T) {
	ret := RemoveKdigits2("20", 2)
	t.Log(ret)
}

func TestMaxNumber(t *testing.T) {
	a := []int{5, 0, 2, 1, 0, 1, 0, 3, 9, 1, 2, 8, 0, 9, 8, 1, 4, 7, 3}
	a2 := []int{7, 6, 7, 1, 0, 1, 0, 5, 6, 0, 5, 0}
	ret := maxNumber(a, a2, 31)
	t.Log(ret)
}

func TestCountPrimes(t *testing.T) {
	ret := countPrimes(10)
	t.Log(ret)
}

func TestFindMaxConsecutiveOnes(t *testing.T) {
	a := []int{1}
	ret := findMaxConsecutiveOnes(a)
	t.Log(ret)
}

func TestFindPoisonedDuration(t *testing.T) {
	a := []int{1, 2, 3, 4}
	ret := findPoisonedDuration(a, 10000)
	t.Log(ret)
}

func TestThirdMax(t *testing.T) {
	a := []int{1, 2}
	ret := thirdMax(a)
	t.Log(ret)
}

func TestIsPossible(t *testing.T) {
	a := []int{1, 3, 4, 5, 6, 7, 8}
	ret := isPossible(a)
	t.Log(ret)
}

func TestFindErrorNums(t *testing.T) {
	a := []int{1, 2, 2, 4}
	ret := findErrorNums(a)
	t.Log(ret)
}

func TestFindShortestSubArray(t *testing.T) {
	a := []int{1, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2}
	ret := findShortestSubArray(a)
	t.Log(ret)
}

func TestFindDisappearedNumbers(t *testing.T) {
	a := []int{1, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2}
	ret := findDisappearedNumbers(a)
	t.Log(ret)
}

func TestFindDuplicates(t *testing.T) {
	a := []int{4, 3, 2, 7, 8, 2, 3, 1}
	ret := findDuplicates(a)
	t.Log(ret)
}

func TestFirstMissingPositive(t *testing.T) {
	a := []int{1, 2, 3}
	ret := firstMissingPositive(a)
	t.Log(ret)
}

func TestHIndex(t *testing.T) {
	a := []int{100}
	ret := hIndex(a)
	t.Log(ret)
}

func TestMinMoves(t *testing.T) {
	a := []int{1, 2, 3, 4}
	ret := minMoves(a)
	t.Log(ret)
}

func TestCheckPossibility(t *testing.T) {
	a := []int{3, 4, 2, 3}
	ret := checkPossibility(a)
	t.Log(ret)
}

func TestMoveZeroes(t *testing.T) {
	a := []int{0, 0, 0, 3, 12, 0}
	moveZeroes(a)
	t.Log(a)
}

func TestRemoveElement(t *testing.T) {
	a := []int{3, 2, 2, 3}
	ret := removeElement(a, 3)
	t.Log(ret)
}

func TestMinSubArrayLen(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	ret := minSubArrayLen(11, a)
	t.Log(ret)
}

func TestSpiralOrder(t *testing.T) {
	a := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	ret := spiralOrder(a)
	t.Log(ret)
}

func TestGenerateMatrix(t *testing.T) {
	a := 3
	ret := generateMatrix(a)
	t.Log(ret)
}

func TestLongestCommonPrefix(t *testing.T) {
	strs := []string{"dog", "racecar", "car"}
	ret := longestCommonPrefix(strs)
	t.Log(ret)
}

func TestTopKFrequent(t *testing.T) {
	nums := []int{1}
	ret := topKFrequent(nums, 1)
	t.Log(ret)
}

func TestPeakIndexInMountainArray(t *testing.T) {
	nums := []int{24, 69, 100, 99, 79, 78, 67, 36, 26, 19}
	ret := peakIndexInMountainArray(nums)
	t.Log(ret)
}

func TestStoneGame(t *testing.T) {
	n := []int{5, 3, 4, 5}
	ret := stoneGame(n)
	t.Log(ret)
}

func TestReadBinaryWatch(t *testing.T) {
	ret := readBinaryWatch(1)
	t.Log(ret)
}

func TestMaximumElementAfterDecrementingAndRearranging(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	ret := maximumElementAfterDecrementingAndRearranging(arr)
	t.Log(ret)
}
