package dp

// LongestCommonSubstring ...
func LongestCommonSubstring(a string, b string) string {
	la, lb, maxLen, endIdx := len(a), len(b), 0, len(a)
	m := matrix(la+1, lb+1)
	for i := 1; i <= la; i++ {
		for j := 1; j <= lb; j++ {
			if a[i-1:i] == b[j-1:j] {
				m[i][j] = m[i-1][j-1] + 1
				if m[i][j] > maxLen {
					maxLen = m[i][j]
					endIdx = i
				}
			}
		}
	}
	return a[endIdx-maxLen : endIdx]
}
