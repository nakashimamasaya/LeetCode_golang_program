package _02418

import (
	"reflect"
	"testing"
)

func TestSample(t *testing.T) {
	names := [][]string{
		{"Mary", "John", "Emma"},
		{"Alice", "Bob", "Bob"},
	}

	heights := [][]int{
		{180, 165, 170},
		{155, 185, 150},
	}
	ans := [][]string{
		{"Mary", "Emma", "John"},
		{"Bob", "Alice", "Bob"},
	}

	for i := 0; i < len(names); i++ {
		if result := sortPeople(names[i], heights[i]); !reflect.DeepEqual(ans[i], result) {
			t.Errorf("miss sample %d\n return %v\n answer %v", i+1, result, ans[i])
		}
	}
}
