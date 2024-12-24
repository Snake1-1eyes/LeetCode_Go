package longsubstr

func lengthOfLongestSubstring(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}

	lastIndex := make(map[byte]int)
	maxLen := 0
	left := 0

	for right := 0; right < n; right++ {
		if idx, ok := lastIndex[s[right]]; ok && idx >= left {
			left = idx + 1
		}
		lastIndex[s[right]] = right
		if currentLen := right - left + 1; currentLen > maxLen {
			maxLen = currentLen
		}
	}
	return maxLen
}
