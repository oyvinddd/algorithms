package dp

/*
CountingSteps finids the number of steps using recursion.

Running time: O(2^n)
*/
func CountingSteps(n int) int {
	if n == 0 {
		return 1
	}
	if n < 0 {
		return 0
	}
	return CountingSteps(n-1) + CountingSteps(n-2)
}

/*
CountingStepsOpt finds the number of steps by using dynamic programming.

Running time: O(n)
*/
func CountingStepsOpt(n int) int {
	mem := make([]int, n+1)
	mem[0] = 1 // base case #1
	mem[1] = 1 // base case #2
	for i := 2; i <= n; i++ {
		mem[i] = mem[i-1] + mem[i-2]
	}
	return mem[n]
}
