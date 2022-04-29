package _0027

import "testing"

func TestSample1(t *testing.T) {
	nums := []int{3, 2, 2, 3}
	val := 2
	ans := 2
	result := RemoveElement(nums, val)
	if result != ans {
		t.Error("sample1 miss result")
	}

	for i := 0; i < ans; i++ {
		if nums[i] == val {
			t.Error("sample1 miss nums")
		}
	}
}

func TestSample2(t *testing.T) {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	val := 2
	ans := 5
	result := RemoveElement(nums, val)
	if result != ans {
		t.Error("sample2 miss val")
	}

	for i := 0; i < ans; i++ {
		if nums[i] == val {
			t.Error("sample2 miss nums")
		}
	}

}
