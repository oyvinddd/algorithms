package dp

// Item struct - the items we're going to put in our knapsack
type Item struct {
	weight int
	value  int
}

// NewItem convenience constructor for item struct
func NewItem(w int, v int) *Item {
	return &Item{weight: w, value: v}
}

// KnapsackOpt runs the optimal knapsack algorithm
func KnapsackOpt(items []Item, cap int) int {
	// Create and initialize matrix with -1 values
	mem := matrix(len(items)+1, cap+1)
	for rows := range mem {
		for cols := range mem[rows] {
			mem[rows][cols] = -1
		}
	}
	return ksOpt(items, mem, len(items)-1, cap)
}

// Knapsack runs the knapsack recursive algorithm
func Knapsack(items []Item, cap int) int {
	return ks(items, len(items)-1, cap)
}

var result int = 0

/*
A memoized recursive algorithm for the knapsack problem where
we use a two-dimensional (n*c) array to store already computed
values for all of our subproblems.

Running time: O(n*c)
*/
func ksOpt(items []Item, mem [][]int, n int, c int) int {
	if mem[n][c] > -1 {
		return mem[n][c]
	}
	if n == 0 || c == 0 {
		result = 0
	} else if items[n].weight > c {
		result = ks(items, n-1, c)
	} else {
		tmp1 := ks(items, n-1, c)
		tmp2 := items[n].value + ks(items, n-1, c-items[n].weight)
		result = max(tmp1, tmp2)
	}
	mem[n][c] = result
	return result
}

/*
Recursive algorithm to solve the knapsack problem
Running time: O(2^n)
*/
func ks(items []Item, n int, c int) int {
	if n == 0 || c == 0 {
		result = 0
	} else if items[n].weight > c {
		result = ks(items, n-1, c)
	} else {
		tmp1 := ks(items, n-1, c)
		tmp2 := items[n].value + ks(items, n-1, c-items[n].weight)
		result = max(tmp1, tmp2)
	}
	return result
}

// max - helper function
func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

// matrix - returns a matrix
func matrix(rows int, cols int) [][]int {
	m := make([][]int, rows)
	for i := range m {
		m[i] = make([]int, cols)
	}
	return m
}
