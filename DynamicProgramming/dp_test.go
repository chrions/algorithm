package DynamicProgramming

import "testing"

//1,1,2,3,5,8,13
func TestFibo(t *testing.T) {
	ret := fib(7)
	t.Log(ret)
}

func TestClimbStairs(t *testing.T) {
	ret := climbStairs(3)
	t.Log(ret)
}

func TestJumpFloor(t *testing.T) {
	ret := jumpFloor(6)
	t.Log(ret)
}

func TestMinCostClimbingStairs(t *testing.T) {
	cost := []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}
	ret := minCostClimbingStairs(cost)
	t.Log(ret)
}
