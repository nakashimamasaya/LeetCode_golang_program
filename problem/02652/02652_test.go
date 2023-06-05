package _02652

import "testing"

func TestSample(t *testing.T) {
	n := []int{
		7,
		10,
		9,
	}

	ans := []int{
		21,
		40,
		30,
	}

	for i, a := range ans {
		if result := sumOfMultiples(n[i]); result != a {
			t.Errorf("miss sample %d\n return %v\n answer %v", i+1, result, a)
		}
	}
}
