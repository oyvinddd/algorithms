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
