package hash

import (
	"fmt"
	"math"
	"sort"
	"strconv"
)

//力扣：49. 字母异位词分组
func groupAnagrams(strs []string) [][]string {
	m := map[string][]string{}
	for _, v := range strs {
		str := []byte(v)
		sort.Slice(str, func(i, j int) bool {
			return str[i] > str[j]
		})
		m[string(str)] = append(m[string(str)], v)
	}
	ret := [][]string{}
	for _, v := range m {
		ret = append(ret, v)
	}
	return ret
}

//力扣：242. 有效的字母异位词
func isAnagram(s string, t string) bool {
	s1 := []byte(s)
	t1 := []byte(t)
	sort.Slice(s1, func(i, j int) bool {
		return s1[i] < s1[j]
	})
	sort.Slice(t1, func(i, j int) bool {
		return t1[i] < t1[j]
	})

	return string(s1) == string(t1)
}

//力扣：349. 两个数组的交集
func intersection(nums1 []int, nums2 []int) []int {
	m := map[int]int{}
	ret := make([]int, 0)
	for _, v := range nums1 {
		m[v] = v
	}
	m2 := map[int]int{}
	for _, v := range nums2 {
		m2[v] = v
	}
	if len(m) > len(m2) {
		m, m2 = m2, m
	}
	for k, v := range m2 {
		if _, ok := m[k]; ok {
			ret = append(ret, v)
		}
	}
	return ret
}

//力扣：202. 快乐数
func isHappy(n int) bool {
	m := map[float64]struct{}{}
	s := fmt.Sprintf("%d", n)
	for {
		sum := float64(0)
		for i := 0; i < len(s); i++ {
			x, _ := strconv.ParseFloat(string(s[i]), 64)
			sum += math.Pow(x, 2)
		}
		if sum == float64(1) {
			return true
		}
		if _, ok := m[sum]; ok {
			return false
		}
		m[sum] = struct{}{}
		s = fmt.Sprintf("%d", int(sum))
	}
}

//力扣：454. 四数相加 II
func fourSumCount(A []int, B []int, C []int, D []int) int {
	m := map[int]int{}
	ret := 0
	for _, a := range A {
		for _, b := range B {
			m[a+b]++
		}
	}

	for _, c := range C {
		for _, d := range D {
			//如果 m[-(c + d)] 存在值，就说明a+b+c+d=0的次数是之前m[a+b]的次数
			ret += m[-(c + d)]
		}
	}
	return ret
}

//力扣：383. 赎金信
func canConstruct(ransomNote string, magazine string) bool {
	m := map[int32]int{}
	cnt := 0
	for _, v := range magazine {
		m[v]++
	}
	for _, v := range ransomNote {
		if c, ok := m[v]; ok {
			if c > 0 {
				m[v]--
				cnt++
			}

		}
	}
	return cnt == len(ransomNote)
}

//力扣 15。三数之和
func threeSum(nums []int) [][]int {
	if len(nums) == 0 {
		return nil
	}
	ret := [][]int{}
	sort.Ints(nums)

	for i := 0; i < len(nums); i++ {
		right := len(nums) - 1
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left := i + 1
		for left < right {
			if left > i+1 && nums[left] == nums[left-1] {
				left++
				continue
			}
			if right < len(nums)-1 && nums[right] == nums[right+1] {
				right--
				continue
			}
			if nums[left]+nums[i]+nums[right] > 0 {
				right--
				continue
			}
			if nums[left]+nums[i]+nums[right] < 0 {
				left++
				continue
			}
			if nums[left]+nums[i]+nums[right] == 0 {
				s := []int{nums[left], nums[i], nums[right]}
				ret = append(ret, s)
				right--
				left++
			}
		}
	}
	return ret
}

//力扣：18. 四数之和
func fourSum(nums []int, target int) [][]int {
	if len(nums) == 0 || len(nums) < 4 {
		return nil
	}
	ret := [][]int{}
	sort.Ints(nums)

	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums); j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			right := len(nums) - 1
			left := j + 1
			for left < right {
				if left > j+1 && nums[left] == nums[left-1] {
					left++
					continue
				}
				if right < len(nums)-1 && nums[right] == nums[right+1] {
					right--
					continue
				}
				if nums[left]+nums[right]+nums[i]+nums[j] == target {
					temp := []int{nums[left], nums[right], nums[i], nums[j]}
					ret = append(ret, temp)
					left++
					right--
				}
				if nums[left]+nums[right]+nums[i]+nums[j] > target {
					right--
					continue
				}
				if nums[left]+nums[right]+nums[i]+nums[j] < target {
					left++
					continue
				}
			}

		}
	}
	return ret
}
