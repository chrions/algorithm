package main

//青蛙跳，每次跳1阶或者2阶
//与斐波拉契相似,唯一不同是0个台阶也有一种走法
//f(0) = 1, f(1) = 1, f(2) = (f1) + f(0), f(n) = f(n-1) + f(n-2)
func NumWays(n int) int {
	if n <= 1 {
		return 1
	}
	a, b := 1, 1
	ret := 0
	for i := 2; i <= n; i++ {
		ret = (a + b) % 1000000007
		a, b = b, ret
	}
	return ret
}
