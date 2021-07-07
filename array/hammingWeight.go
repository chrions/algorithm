package main

import "math/bits"

//剑指offer 15
func hammingWeight(num uint32) int {
	return bits.OnesCount32(num)
}
