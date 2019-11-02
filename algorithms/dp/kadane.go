package dp

/*
Kadane computes the maximum-sum contigous subsequence.

Running time: O(n) (space complexity: O(n))
*/
func Kadane(a []int) int {
	arr := make([]int, len(a))
	for i := 1; i < len(a); i++ {
		arr[i] = max(arr[i-1], 0) + a[i]
	}
	return arr[len(a)-1]
}
