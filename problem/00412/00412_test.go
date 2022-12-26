package _00412

import (
	"reflect"
	"testing"
)

func TestSamp(t *testing.T) {
	ns := []int{
		3,
		5,
		15,
	}

	ans := [][]string{
		{"1", "2", "Fizz"},
		{"1", "2", "Fizz", "4", "Buzz"},
		{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"},
	}

	for i, n := range ns {
		result := fizzBuzz(n)
		if !reflect.DeepEqual(result, ans[i]) {
			t.Errorf("miss sample %d\n return %v\n answer %v", i+1, result, ans[i])
		}
	}
}
