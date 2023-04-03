package _02600

import "testing"

type Input struct {
	NumOnes    int
	NumZeros   int
	NumNegOnes int
	k          int
}

func TestSample(t *testing.T) {
	input := []Input{
		{3, 2, 0, 2},
		{3, 2, 0, 4},
	}

	ans := []int{
		2,
		3,
	}

	for i, v := range input {
		if result := kItemsWithMaximumSum(v.NumOnes, v.NumZeros, v.NumNegOnes, v.k); result != ans[i] {
			t.Errorf("miss sample %d\n return %v\n answer %v", i+1, result, ans[i])
		}
	}
}
