package _00038

import "testing"

func TestSample(t *testing.T) {
	n := []int{
		1,
		4,
	}

	ans := []string{
		"1",
		"1211",
	}

	for i, tmp := range n {
		result := countAndSay(tmp)

		if result != ans[i] {
			t.Errorf("miss sample %d\n return %v\n answer %v", i+1, result, ans[i])
		}
	}
}
