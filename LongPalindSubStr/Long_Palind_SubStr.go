package longpalindsubstr

func longestPalindrome(s string) string {

	if len(s) < 2 {
		return s
	}

	start, maxLen := 0, 1
	for i := 0; i < len(s); i++ {
		left, right := expandAroundCenter(s, i, i)
		if right-left+1 > maxLen {
			start = left
			maxLen = right - left + 1
		}

		left, right = expandAroundCenter(s, i, i+1)
		if right-left+1 > maxLen {
			start = left
			maxLen = right - left + 1
		}
	}

	return s[start : start+maxLen]
}

func expandAroundCenter(s string, left, right int) (int, int) {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return left + 1, right - 1
}
