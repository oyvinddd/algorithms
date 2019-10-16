package dac

// MergeAndCount counts the number of inversions between two
// lists and returns the sorted result
func MergeAndCount(a []int, b []int) (int, []int) { //WIP:
	pa, pb, pc, cnt := 0, 0, 0, 0
	lena, lenb := len(a), len(b)
	out := make([]int, lena+lenb)
	for pa < lena && pb < lenb {
		elema, elemb := a[pa], b[pb]
		if elema <= elemb {
			out[pc] = elema
			pa++
		} else {
			out[pc] = elemb
			pb++
		}
		if elemb < elema && pa < lena-1 {
			cnt++
		}
		pc++
	}
	return cnt, out
}
