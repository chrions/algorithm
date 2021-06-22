package Stack

import "testing"

func TestIsValid(t *testing.T) {
	ret := IsValid("{()}")
	t.Log(ret)
}

func TestCalculate(t *testing.T) {
	ret := calculate("-(3+(4-1))")
	t.Log(ret)
}

func TestRemoveDuplicates(t *testing.T) {
	ret := removeDuplicates("abbaca")
	t.Log(ret)
}

func TestMaxSlidingWindow(t *testing.T) {
	nums := []int{7, 2, 4}
	ret := maxSlidingWindow(nums, 2)
	t.Log(ret)
}
