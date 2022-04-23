package _0026

import "testing"

func TestSample1(t *testing.T) {
	result := RemoveDuplicates([]int{1, 1, 2})
	if result != 2 {
		t.Error("sample1 miss")
	}
}

func TestSample2(t *testing.T) {
	result := RemoveDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4})
	if result != 5 {
		t.Error("sample2 miss")
	}
}
