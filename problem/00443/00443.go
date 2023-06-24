package _00443

import "strconv"

// 443. String Compression
// https://leetcode.com/problems/string-compression/

func compress(chars []byte) int {
	ans, tmp, count := 0, chars[0], 0

	for _, v := range chars {
		if v != tmp {
			ans = charsCompress(chars, tmp, count, ans)
			tmp, count = v, 1
		} else {
			count++
		}
	}
	ans = charsCompress(chars, tmp, count, ans)

	return ans
}

func charsCompress(chars []byte, b byte, c, index int) int {
	tmpS := string(b)
	if c > 1 {
		tmpS += strconv.Itoa(c)
	}

	for i := 0; i < len(tmpS); i, index = i+1, index+1 {
		chars[index] = tmpS[i]
	}

	return index
}
