package greedy

// Job struct
type Job struct {
	Name  string
	Start int
	End   int
}

// NewJob convenience constructor
func NewJob(name string, start int, end int) *Job {
	return &Job{Name: name, Start: start, End: end}
}

// IntervalScheduling returns an optimal set of compatible jobs
func IntervalScheduling(jobs []Job) []Job {
	// TODO: sort the input list of jobs according to finish time
	// Choose any nlogn algorithm for the sorting
	optimal := make([]Job, len(jobs))

	cnt, cnto := 0, 0
	for cnt < len(jobs) {
		job := jobs[cnt]
		cnt++
		if cnto > 0 && !compatible(optimal[cnto-1], job) {
			continue
		}
		optimal[cnto] = job
		cnto++
	}
	return optimal
}

func compatible(j1 Job, j2 Job) bool {
	return j1.End < j2.Start
}

// TODO: use this qs at the start of the greedy algorithms for O(nlogn) sorting
/*
func quicksort(a []int) []int {
	if len(a) < 2 {
		return a
	}

	left, right := 0, len(a)-1

	pivot := rand.Int() % len(a)

	a[pivot], a[right] = a[right], a[pivot]

	for i, _ := range a {
		if a[i] < a[right] {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}

	a[left], a[right] = a[right], a[left]

	quicksort(a[:left])
	quicksort(a[left+1:])

	return a
}
*/
