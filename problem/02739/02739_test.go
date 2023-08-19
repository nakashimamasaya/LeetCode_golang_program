package _02739

import "testing"

func TestSample(t *testing.T) {
	mainTank := []int{
		5,
		1,
	}

	additionalTank := []int{
		10,
		2,
	}

	ans := []int{
		60,
		10,
	}

	for i, a := range ans {
		if result := distanceTraveled(mainTank[i], additionalTank[i]); result != a {
			t.Errorf("miss sample %d\n return %v\n answer %v", i+1, result, a)
		}
	}
}
