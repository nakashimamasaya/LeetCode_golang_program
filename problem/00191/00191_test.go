package _00191

import (
	"strconv"
	"testing"
)

func TestSample(t *testing.T) {
	nums := []string{
		"00000000000000000000000000001011",
		"00000000000000000000000010000000",
		"11111111111111111111111111111101",
	}

	ans := []int{
		3,
		1,
		31,
	}

	for i := 0; i < len(nums); i++ {
		num, _ := strconv.ParseUint(nums[i], 2, 32)
		result := hammingWeight(uint32(num))
		if result != ans[i] {
			t.Errorf("miss sample %d\n return %v\n answer %v", i+1, result, ans[i])
		}
	}
}
