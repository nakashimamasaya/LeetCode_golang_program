package _00190

import (
	"strconv"
	"testing"
)

func TestSample(t *testing.T) {
	nums := []string{
		"00000010100101000001111010011100",
		"11111111111111111111111111111101",
	}

	ans := []uint32{
		964176192,
		3221225471,
	}

	for i := 0; i < len(nums); i++ {
		num, _ := strconv.ParseUint(nums[i], 2, 32)
		result := reverseBits(uint32(num))
		if result != ans[i] {
			t.Errorf("miss sample %d\n return %v\n answer %v", i+1, result, ans[i])
		}
	}
}
