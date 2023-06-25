package _00039

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

func TestSample(t *testing.T) {
	candidates := [][]int{
		{2, 3, 6, 7},
		{2, 3, 5},
		{2},
	}

	target := []int{
		7,
		8,
		1,
	}

	ans := [][][]int{
		{{2, 2, 3}, {7}},
		{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}},
		{},
	}

	for i, a := range ans {
		if result := combinationSum(candidates[i], target[i]); !checkSlice(a, result) {
			t.Errorf("miss sample %d\n return %v\n answer %v", i+1, result, a)
		}
	}
}

func checkSlice(a, b [][]int) bool {
	ansF := true
	for _, va := range a {
		sort.Ints(va)
		flag := false
		for _, vb := range b {
			sort.Ints(vb)
			flag = flag || reflect.DeepEqual(va, vb)
		}
		ansF = ansF && flag
		if !flag {
			fmt.Println(va)
		}
	}
	return ansF
}
