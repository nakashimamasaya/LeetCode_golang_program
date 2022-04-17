package _0067

// 67. Add Binary
// https://leetcode.com/problems/add-binary/

func AddBinary(a string, b string) string {
	a = reverseStr(a)
	b = reverseStr(b)
	var ans []byte
	var aN, bN, carry byte

	carry = 0
	for i := 0; i < len(a) || i < len(b); i++ {
		aN, bN = 48, 48
		if i < len(a) {
			aN = a[i]
		}
		if i < len(b) {
			bN = b[i]
		}

		calc := aN + bN + carry
		ans = append(ans, calc%2+48)
		carry = calc / 98
	}
	if carry == 1 {
		ans = append(ans, 49)
	}

	return reverseStr(string(ans))
}

func reverseStr(s string) string {
	tmp := []byte(s)

	for i, j := 0, len(tmp)-1; i < j; i, j = i+1, j-1 {
		tmp[i], tmp[j] = tmp[j], tmp[i]
	}

	return string(tmp)
}
