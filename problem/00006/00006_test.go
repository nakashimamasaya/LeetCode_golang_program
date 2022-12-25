package _00006

import "testing"

func TestSample(t *testing.T) {
	ss := []string{
		"PAYPALISHIRING",
		"PAYPALISHIRING",
		"A",
	}

	numRows := []int{
		3,
		4,
		1,
	}

	ans := []string{
		"PAHNAPLSIIGYIR",
		"PINALSIGYAHRPI",
		"A",
	}

	for i := 0; i < len(ss); i++ {
		result := convert(ss[i], numRows[i])
		if result != ans[i] {
			t.Errorf("miss sample %d\n return %v\n answer %v", i+1, result, ans[i])
		}
	}
}
