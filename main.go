package main

import (
	_0027 "LeetCode/problem/00027"
	_0028 "LeetCode/problem/00028"
	_0035 "LeetCode/problem/00035"
	_0053 "LeetCode/problem/00053"
	_0058 "LeetCode/problem/00058"
	_0066 "LeetCode/problem/00066"
	_0067 "LeetCode/problem/00067"
	_0069 "LeetCode/problem/00069"
	_0070 "LeetCode/problem/00070"
	_0083 "LeetCode/problem/00083"
	_0136 "LeetCode/problem/00136"
	_0171 "LeetCode/problem/00171"
	_0206 "LeetCode/problem/00206"
	_0217 "LeetCode/problem/00217"
	_0219 "LeetCode/problem/00219"
	_0704 "LeetCode/problem/00704"
	_01046 "LeetCode/problem/01046"
	"fmt"
)

func main() {
	// 27. Remove Element
	separate("27")
	fmt.Println(_0027.RemoveElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2))

	// 28. Implement strStr()
	separate("28")
	fmt.Println(_0028.StrStr("hello", "ll"))

	// 35. Search Insert Position
	separate("35")
	fmt.Println(_0035.SearchInsert([]int{1, 3, 5, 6}, 5))

	// 53. Maximum Subarray
	separate("53")
	fmt.Println(_0053.MaxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))

	// 58. Length of Last Word
	separate("58")
	fmt.Println(_0058.LengthOfLastWord("   fly me   to   the moon  "))

	// 66. Plus One
	separate("66")
	fmt.Println(_0066.PlusOne([]int{9, 9, 9}))

	// 67. Add Binary
	separate("67")
	fmt.Println(_0067.AddBinary("11", "1"))

	// 69. Sqrt(x)
	separate("69")
	fmt.Println(_0069.MySqrt(0))

	// 70. Climbing Stairs
	separate("70")
	fmt.Println(_0070.ClimbStairs(4))

	// 83. Remove Duplicates from Sorted List
	separate("83")
	fmt.Println(_0083.DeleteDuplicates(_0083.SetSample()))

	// 136. Single Number
	separate("136")
	fmt.Println(_0136.SingleNumber([]int{4, 1, 2, 1, 2}))

	// 171. Excel Sheet Column Number
	separate("171")
	fmt.Println(_0171.TitleToNumber("AB"))

	// 206. Reverse Linked List
	separate("206")
	fmt.Println(_0206.ReverseList(_0206.SetSample()).Val)

	// 217. Contains Duplicate
	separate("217")
	fmt.Println(_0217.ContainsDuplicate([]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}))

	// 219. Contains Duplicate II
	separate("219")
	fmt.Println(_0219.ContainsNearbyDuplicate([]int{1, 2, 3, 1, 2, 3}, 2))

	// 704. Binary Search
	separate("704")
	fmt.Println(_0704.Search([]int{-1, 0, 3, 5, 9, 12}, 9))

	// 1046. Last Stone Weight
	separate("1046")
	fmt.Println(_01046.LastStoneWeight([]int{2, 7, 4, 1, 8, 1}))
}

func separate(num string) {
	fmt.Println(num + "----------------------------------")
}
