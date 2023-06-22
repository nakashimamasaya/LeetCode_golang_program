package _00762

// 762. Prime Number of Set Bits in Binary Representation
// https://leetcode.com/problems/prime-number-of-set-bits-in-binary-representation/

func countPrimeSetBits(left int, right int) int {
	prime := map[byte]bool{}
	prime[1] = true
	for i := byte(2); i < 20; i++ {
		if !prime[i] {
			for j := byte(2); j*i < 20; j++ {
				prime[j*i] = true
			}
		}
	}

	ans := 0
	for ; left <= right; left++ {
		if !prime[countBits(left)] {
			ans++
		}
	}
	return ans
}

func countBits(num int) byte {
	var ans byte
	for ; num > 0; num = num >> 1 {
		ans += byte(num) % 2
	}
	return ans
}
