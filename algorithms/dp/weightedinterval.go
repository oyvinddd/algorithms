package dp

// Job the jobs that needs to be executed
type Job struct {
	name   string
	start  int
	end    int
	weight int
}

/*
WeightedIntervalScheduling executes the algorithm and makes
use of memoization (a []int) to solve the provlem. Note that
the set of jobs needs to be sorted according to end time for
the algorithm to work properly.

Running time: O(n)
*/
func WeightedIntervalScheduling(m []Job) int {
	a := make([]int, len(m))
	// initialize opt array with negative values
	for i := 0; i < len(a); i++ {
		a[i] = -1
	}
	return wis(m, a, 1) // TODO: start from n?
}

func wis(m []Job, a []int, i int) int {
	if i == 0 {
		return 0
	}
	if a[i] > -1 {
		return a[i]
	}
	a[i] = max(m[i].weight+wis(m, a, i), wis(m, a, i-1))
	return a[i]
}
