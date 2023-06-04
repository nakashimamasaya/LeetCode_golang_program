package _00168

// 168. Excel Sheet Column Title
// https://leetcode.com/problems/excel-sheet-column-title/

func convertToTitle(columnNumber int) string {
	ans := ""

	for ; columnNumber > 0; columnNumber = (columnNumber - 1) / 26 {
		ans = string(65+rune(columnNumber-1)%26) + ans
	}
	return ans
}
