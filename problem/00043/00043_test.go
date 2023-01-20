package _00043

import "testing"

func TestSample(t *testing.T) {
	num1 := []string{
		"2",
		"123",
	}

	num2 := []string{
		"3",
		"456",
	}

	ans := []string{
		"6",
		"56088",
	}

	for i := 0; i < len(num1); i++ {
		if result := multiply(num1[i], num2[i]); result != ans[i] {
			t.Errorf("miss sample %d\n return %v\n answer %v", i+1, result, ans[i])
		}
	}
}
