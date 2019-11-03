package dp

/*
LCS computes the longest commont subsequence using a recursive approach.

Running time: O(2^n)
*/
func LCS(a string, b string, i int, j int) int {
	if i == len(a) || j == len(b) {
		return 0
	}
	if a[i] == b[j] {
		return 1 + LCS(a, b, i+1, j+1)
	}
	return max(LCS(a, b, i+1, j), LCS(a, b, i, j+1))
}
