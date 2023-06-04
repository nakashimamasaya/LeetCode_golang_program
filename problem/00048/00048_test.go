package _00048

import (
	"reflect"
	"testing"
)

func TestSample(t *testing.T) {
	matrix := [][][]int{
		{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
		{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}},
	}

	ans := [][][]int{
		{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}},
		{{15, 13, 2, 5}, {14, 3, 4, 1}, {12, 6, 8, 9}, {16, 7, 10, 11}},
	}

	for i, a := range ans {
		if rotate(matrix[i]); !reflect.DeepEqual(matrix[i], a) {
			t.Errorf("miss sample %d\n return %v\n answer %v", i+1, matrix[i], a)
		}
	}
}
