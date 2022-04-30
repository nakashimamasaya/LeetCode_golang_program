package _0028

import "testing"

func TestSample1(t *testing.T) {
	haystack := "hello"
	needle := "ll"
	ans := 2

	result := StrStr(haystack, needle)

	if result != ans {
		t.Error("sample1 miss")
	}
}

func TestSample2(t *testing.T) {
	hayStack := "aaaaa"
	needle := "bba"
	ans := -1

	result := StrStr(hayStack, needle)

	if result != ans {
		t.Error("sample2 miss")
	}
}
