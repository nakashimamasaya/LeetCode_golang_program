package _0066

import (
	"testing"
)

func TestSample(t *testing.T) {
	digits := [][]int{
		{1, 2, 3},
		{4, 3, 2, 1},
		{9},
	}
	ans := [][]int{
		{1, 2, 4},
		{4, 3, 2, 2},
		{1, 0},
	}

	for i := 0; i < len(digits); i++ {
		an := PlusOne(digits[i])
		if !arrayEqual(an, ans[i]) {
			t.Errorf("miss sample %d", i+1)
		}
	}
}

func arrayEqual(a []int, b []int) bool {
	result := true
	for i := 0; i < len(a) && i < len(b) && len(a) == len(b); i++ {
		result = result && a[i] == b[i]
	}

	return result && len(a) == len(b)
}
