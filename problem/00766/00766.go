package _00766

// 766. Toeplitz Matrix
// https://leetcode.com/problems/toeplitz-matrix/

func isToeplitzMatrix(matrix [][]int) bool {
	for i := len(matrix) - 1; i > 0; i-- {
		for j := len(matrix[0]) - 1; j > 0; j-- {
			if matrix[i-1][j-1] != matrix[i][j] {
				return false
			}
		}
	}
	return true
}
