package _02614

import (
	"sort"
)

// 2614. Prime In Diagonal
// https://leetcode.com/problems/prime-in-diagonal/

func diagonalPrime(nums [][]int) int {
	primeNumber := make([]bool, 4000001)
	primeNumber[1] = true
	for i := 2; i < len(primeNumber); i++ {
		for j := i * 2; !primeNumber[i] && j < len(primeNumber); j += i {
			primeNumber[j] = true
		}
	}

	n := []int{}
	for i, j := 0, len(nums)-1; i < len(nums); i, j = i+1, j-1 {
		n = append(n, nums[i][i], nums[i][j])
	}

	sort.Ints(n)
	for i := len(n) - 1; i >= 0; i-- {
		if !primeNumber[n[i]] {
			return n[i]
		}
	}
	return 0
}
