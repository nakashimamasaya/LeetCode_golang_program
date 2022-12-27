package _00036

import "testing"

func TestSample(t *testing.T) {
	boards := [][][]byte{
		{
			{53, 51, 46, 46, 55, 46, 46, 46, 46},
			{54, 46, 46, 49, 57, 53, 46, 46, 46},
			{46, 57, 56, 46, 46, 46, 46, 54, 46},
			{56, 46, 46, 46, 54, 46, 46, 46, 51},
			{52, 46, 46, 56, 46, 51, 46, 46, 49},
			{55, 46, 46, 46, 50, 46, 46, 46, 54},
			{46, 54, 46, 46, 46, 46, 50, 56, 46},
			{46, 46, 46, 52, 49, 57, 46, 46, 53},
			{46, 46, 46, 46, 56, 46, 46, 55, 57},
		},
		{
			{56, 51, 46, 46, 55, 46, 46, 46, 46},
			{54, 46, 46, 49, 57, 53, 46, 46, 46},
			{46, 57, 56, 46, 46, 46, 46, 54, 46},
			{56, 46, 46, 46, 54, 46, 46, 46, 51},
			{52, 46, 46, 56, 46, 51, 46, 46, 49},
			{55, 46, 46, 46, 50, 46, 46, 46, 54},
			{46, 54, 46, 46, 46, 46, 50, 56, 46},
			{46, 46, 46, 52, 49, 57, 46, 46, 53},
			{46, 46, 46, 46, 56, 46, 46, 55, 57},
		},
	}

	ans := []bool{
		true,
		false,
	}

	for i, board := range boards {
		result := isValidSudoku(board)
		if result != ans[i] {
			t.Errorf("miss sample %d\n return %v\n answer %v", i+1, result, ans[i])
		}
	}
}
