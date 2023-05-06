package _00661

// 661. Image Smoother
// https://leetcode.com/problems/image-smoother/

func imageSmoother(img [][]int) [][]int {
	ans := make([][]int, len(img))
	x := []int{-1, 0, 1, -1, 0, 1, -1, 0, 1}
	y := []int{-1, -1, -1, 0, 0, 0, 1, 1, 1}

	for i := 0; i < len(img); i++ {
		ans[i] = make([]int, len(img[i]))
		for j, count := 0, 0; j < len(img[i]); j, count = j+1, 0 {
			for k := 0; k < 9; k++ {
				if i+x[k] < 0 || i+x[k] >= len(img) ||
					j+y[k] < 0 || j+y[k] >= len(img[i]) {
					continue
				}
				ans[i][j] += img[i+x[k]][j+y[k]]
				count++
			}
			ans[i][j] /= count
		}
	}
	return ans
}
