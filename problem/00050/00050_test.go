package _00050

import "testing"

func TestSample(t *testing.T) {
	x := []float64{
		2,
		2.1,
		2,
	}

	n := []int{
		10,
		3,
		-2,
	}

	ans := []float64{
		1024,
		9.261,
		0.25,
	}

	for i := 0; i < len(x); i++ {
		if result := myPow(x[i], n[i]); float64(int(result*10000))/10000 != ans[i] {
			t.Errorf("miss sample %d\n return %v\n answer %v", i+1, result, ans[i])
		}
	}
}
