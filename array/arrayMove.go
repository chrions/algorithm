package main

import "math"

/**
思路：将除了一个元素之外的全部元素+1，等价于将该元素-1，因为我们只对元素的相对大小感兴趣。因此，该问题简化为需要进行的减法次数。
数学法是真的牛逼！学好数学很重要
*/
func minMoves(nums []int) int {
	min := math.MaxInt64
	moves := 0
	for _, v := range nums {
		if v < min {
			min = v
		}
	}

	for _, v := range nums {
		moves += v - min
	}
	return moves
}

//来源：力扣（LeetCode）665
/*要点：
4，2，3
-1，4，2，3
2，3，3，2，4
通过分析上面三个例子可以发现，当我们发现后面的数字小于前面的数字产生冲突后，
[1]有时候需要修改前面较大的数字(比如前两个例子需要修改4)，

[2]有时候却要修改后面较小的那个数字(比如前第三个例子需要修改2)，

那么有什么内在规律吗？是有的，判断修改那个数字其实跟再前面一个数的大小有关系，
首先如果再前面的数不存在（比如例子1）4前面没有数字了，我们直接修改前面的数字为当前的数字2即可。
而当再前面的数字存在，并且小于当前数时（比如例子2）-1小于2，我们还是需要修改前面的数字4为当前数字2；
如果再前面的数大于当前数（比如例子3）3大于2，我们需要修改当前数2为前面的数3。
*/
func checkPossibility(nums []int) bool {
	changeCnt := 0
	for i := 1; i < len(nums); i++ {
		if nums[i-1] <= nums[i] {
			continue
		}
		changeCnt++
		if i-2 >= 0 && nums[i-2] > nums[i] {
			nums[i] = nums[i-1]
		} else {
			nums[i-1] = nums[i]
		}
	}
	return changeCnt <= 1
}

//来源：力扣（LeetCode）283
//0,1,0,3,12
func moveZeroes(nums []int) {
	slow := 0 //慢指针，指向第一个0，用来和fast指向的非0数字交换
	for fast := 0; fast < len(nums); fast++ {
		//如果fast遇到的数字是0，slow停在0的位置上,此时slow指向的是第一个0

		if nums[fast] != 0 { //如果fast遇到第一个非0的数字，和slow交换
			nums[slow], nums[fast] = nums[fast], nums[slow]
			slow++ //此时slow指向的不是0了，往后移一位
		}
	}
}

//来源：力扣（LeetCode）27
func removeElement(nums []int, val int) int {
	i, j, n := 0, 0, len(nums)
	for i < n {
		if nums[i] != val {
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
		i++
	}
	return len(nums[0:j])
}
