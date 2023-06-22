package _00762

import "testing"

func TestSample(t *testing.T) {
	left := []int{
		6,
		10,
	}

	right := []int{
		10,
		15,
	}

	ans := []int{
		4,
		5,
	}

	for i, a := range ans {
		if result := countPrimeSetBits(left[i], right[i]); result != a {
			t.Errorf("miss sample %d\n return %v\n answer %v", i+1, result, a)
		}
	}
}
