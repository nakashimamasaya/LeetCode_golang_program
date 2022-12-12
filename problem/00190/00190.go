package _00190

import "math/bits"

// 190. Reverse Bits
// https://leetcode.com/problems/reverse-bits/

func reverseBits(num uint32) uint32 {
	return bits.Reverse32(num)
}
