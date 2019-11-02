package dp

// Item struct - the items we're going to put in our knapsack
type Item struct {
	weight int
	value  int
}

var mem [][]int
var result int = 0

// NewItem convenience constructor for item struct
func NewItem(w int, v int) *Item {
	return &Item{weight: w, value: v}
}

// KnapsackOpt runs the optimal knapsack algorithm
func KnapsackOpt(items []Item, cap int) int {
	// Create and initialize matrix with -1 values
	mem = matrix(len(items)+1, cap+1)
	for rows := range mem {
		for cols := range mem[rows] {
			mem[rows][cols] = -1
		}
	}
	return ksOpt(items, len(items)-1, cap)
}

// Knapsack runs the knapsack recursive algorithm
func Knapsack(items []Item, cap int) int {
	return ks(items, len(items), cap)
}

/*
A memoized recursive algorithm for the knapsack problem where
we use a two-dimensional (n*c) array to store already computed
values for all of our subproblems.

Running time: O(n*c)
*/
func ksOpt(items []Item, n int, c int) int {
	if mem[n][c] > -1 {
		return mem[n][c]
	}
	if n == 0 || c == 0 {
		result = 0
	} else if items[n].weight > c {
		result = ksOpt(items, n-1, c)
	} else {
		tmp1 := ksOpt(items, n-1, c)
		tmp2 := items[n].value + ksOpt(items, n-1, c-items[n].weight)
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
		return 0
	}
	if items[n-1].weight > c {
		return ks(items, n-1, c)
	}
	tmp1 := ks(items, n-1, c)
	tmp2 := items[n-1].value + ks(items, n-1, c-items[n].weight)
	return max(tmp1, tmp2)
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
