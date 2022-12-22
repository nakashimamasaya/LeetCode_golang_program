package _00242

import (
	"reflect"
	"sort"
)

// 242. Valid Anagram
// https://leetcode.com/problems/valid-anagram/

func isAnagram(s string, t string) bool {
	return reflect.DeepEqual(stringSort(s), stringSort(t))
}

func stringSort(s string) []string {
	var slice []string
	for _, tmp := range s {
		slice = append(slice, string(tmp))
	}
	sort.Strings(slice)
	return slice
}
