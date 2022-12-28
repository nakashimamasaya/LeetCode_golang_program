package _00228

import (
	"reflect"
	"testing"
)

func TestSample(t *testing.T) {
	nums := [][]int{
		{0, 1, 2, 4, 5, 7},
		{0, 2, 3, 4, 6, 8, 9},
	}

	ans := [][]string{
		{"0->2", "4->5", "7"},
		{"0", "2->4", "6", "8->9"},
	}

	for i, n := range nums {
		result := summaryRanges(n)
		if !reflect.DeepEqual(ans[i], result) {
			t.Errorf("miss sample %d\n return %v\n answer %v", i+1, result, ans[i])

		}
	}
}
