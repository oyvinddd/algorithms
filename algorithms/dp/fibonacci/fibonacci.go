package dp

// https://en.wikipedia.org/wiki/Fibonacci_number

/*
Fib computes the n first Fibonacci numbers recursively (classic approach)
Running time: O(2^n)
*/
func Fib(n int) int {
	if n <= 2 {
		return 1
	}
	return Fib(n-1) + Fib(n-2)
}

/*
FibOpt computes the n first Fibonacci numbers using memoization (DP approach)
(see implementation details in the funciton "fo" below)
Running time: (2n + 1) * O(1) => O(n)
*/
func FibOpt(n int) int {
	return fo(n, make([]int, n+1))
}

// fo computes the n first Fibonacci numbers using memoization (dp)
func fo(n int, m []int) int {
	if m[n] > 0 {
		return m[n]
	}
	var res int
	if n == 1 || n == 2 {
		res = 1
	} else {
		res = fo(n-1, m) + fo(n-2, m)
		m[n] = res
	}
	return res
}
