package greedy

import (
	"sort"
)

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
// Running time: O(nlogn)
//	- sorting
//	- iterating n times (constant time operations each time)
func IntervalScheduling(jobs []Job) []Job {
	// Sort the slice according to end time O(nlogn)
	sort.Slice(jobs, func(i, j int) bool {
		return jobs[i].End < jobs[j].End
	})
	optimal := make([]Job, len(jobs))
	cnt, cnto := 0, 0
	for cnt < len(jobs) {
		job := jobs[cnt]
		cnt++
		if cnto > 0 && !isCompatible(optimal[cnto-1], job) {
			continue
		}
		optimal[cnto] = job
		cnto++
	}
	return optimal
}

func isCompatible(this Job, that Job) bool {
	return this.End <= that.Start
}
