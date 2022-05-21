package _0012

import (
	"testing"
)

func TestSample(t *testing.T) {
	nums := []int{
		3,
		58,
		1994,
	}

	ans := []string{
		"III",
		"LVIII",
		"MCMXCIV",
	}

	for i := 0; i < len(ans); i++ {
		result := intToRoman(nums[i])
		if result != ans[i] {
			t.Errorf("miss sample %v\n return %v \n answer %v", i+1, result, ans[i])
		}
	}
}

func TestInput(t *testing.T) {
	num := 9
	ans := "IX"

	result := intToRoman(num)
	if result != ans {
		t.Errorf("miss Input 9\n return %v\n answer %v\n", result, ans)
	}
}
