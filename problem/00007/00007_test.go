package _00007

import "testing"

func TestSample(t *testing.T) {
	xs := []int{
		123,
		-123,
		120,
	}

	ans := []int{
		321,
		-321,
		21,
	}

	for i := 0; i < len(xs); i++ {
		result := reverse(xs[i])
		if result != ans[i] {
			t.Errorf("miss sample %d\n return %v\n answer %v", i+1, result, ans[i])
		}
	}
}
