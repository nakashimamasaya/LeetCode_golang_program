package _00258

import "testing"

func TestSample(t *testing.T) {
	nums := []int{
		38,
		0,
	}

	ans := []int{
		2,
		0,
	}

	for i := 0; i < len(nums); i++ {
		result := addDigits(nums[i])
		if result != ans[i] {
			t.Errorf("miss sample %d\n return %v\n answer %v", i+1, result, ans[i])
		}
	}
}
