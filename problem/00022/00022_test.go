package _0022

import "testing"

func TestSample1(t *testing.T) {
	result := GenerateParenthesis(3)
	ans := []string{"((()))", "(()())", "(())()", "()(())", "()()()"}
	count := 0

	for _, s := range ans {
		for i, r := range result {
			if s == r {
				result[i] = ""
				count++
				break
			}
		}
	}

	if count != len(result) {
		t.Error("sample2 miss")
	}
}

func TestSample2(t *testing.T) {
	result := GenerateParenthesis(1)
	ans := "()"

	if len(result) != 1 || ans != result[0] {
		t.Error("sample2 miss")
	}
}
