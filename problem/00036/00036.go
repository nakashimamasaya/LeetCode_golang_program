package _00036

// 36. Valid Sudoku
// https://leetcode.com/problems/valid-sudoku/

func isValidSudoku(board [][]byte) bool {
	const vertical, horizontal, number = 9, 9, 9
	const asciiOne = 49
	verticalValid, horizonValid := [vertical][number]bool{{}}, [horizontal][number]bool{{}}
	subValid := [vertical / 3][horizontal / 3][number]bool{}
	for i := 0; i < vertical; i++ {
		for j := 0; j < horizontal; j++ {
			num := int(board[i][j]) - asciiOne
			if num < 0 {
				continue
			}
			if verticalValid[i][num] || horizonValid[j][num] || subValid[i/3][j/3][num] {
				return false
			}
			verticalValid[i][num], horizonValid[j][num] = true, true
			subValid[i/3][j/3][num] = true
		}
	}
	return true
}
