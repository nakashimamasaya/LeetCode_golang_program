package _00605

// 605. Can Place Flowers
// https://leetcode.com/problems/can-place-flowers/?envType=study-plan-v2&id=leetcode-75

func canPlaceFlowers(flowerbed []int, n int) bool {
	flowerbed = append(append([]int{0}, flowerbed...), []int{0}...)

	for i := 1; i < len(flowerbed)-1 && n > 0; i++ {
		if flowerbed[i] == 0 && flowerbed[i-1] == 0 && flowerbed[i+1] == 0 {
			flowerbed[i], n = 1, n-1
		}
	}
	return n == 0
}
