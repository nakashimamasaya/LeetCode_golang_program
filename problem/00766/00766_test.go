package _00766

import "testing"

func TestSample(t *testing.T) {
	matrix := [][][]int{
		{{1, 2, 3, 4}, {5, 1, 2, 3}, {9, 5, 1, 2}},
		{{1, 2}, {2, 2}},
	}

	ans := []bool{
		true,
		false,
	}

	for i, a := range ans {
		if result := isToeplitzMatrix(matrix[i]); result != a {
			t.Errorf("miss sample %d\n return %v\n answer %v", i+1, result, a)
		}
	}
}
