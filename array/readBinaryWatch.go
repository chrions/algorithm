package main

import (
	"fmt"
	"math/bits"
)

//力扣401
func readBinaryWatch(turnedOn int) []string {
	ans := []string{}
	for h := uint8(0); h < 12; h++ {
		for m := uint8(0); m < 60; m++ {
			if bits.OnesCount8(h)+bits.OnesCount8(m) == turnedOn {
				ans = append(ans, fmt.Sprintf("%d:%02d", h, m))
			}
		}
	}
	return ans
}
