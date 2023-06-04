package _00048

// 48. Rotate Image
// https://leetcode.com/problems/rotate-image/

func rotate(matrix [][]int) {
	tmpMat := make([][]int, len(matrix))
	for i, m := range matrix {
		tmpMat[i] = make([]int, len(matrix))
		for j, mm := range m {
			tmpMat[i][j] = mm
		}
	}

	for i, m := range tmpMat {
		for j, mm := range m {
			matrix[j][len(tmpMat)-1-i] = mm
		}
	}
}
