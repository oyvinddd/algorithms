package dac

func MergeAndCount(a []int, b []int) ([]int, int) {
	m := make([]int, len(a)+len(b))
	c, pa, pb := 0, 0, 0
	for c < len(m) {
		if pa < len(a) && a[pa] < b[pb] {
			m[c] = a[pa]
			pa++
		} else if pb < len(b) {
			m[c] = b[pb]
			pb++
		}
		c++
	}
	return m, 0
}
