package linearsort

/*
Sort returns a sorted copy of the input slice
Running time: linear time proprtional to the size of the input
*/
func Sort(a []int) []int {
	b := make([]int, len(a))
	for index := range a {
		b[a[index]] = a[index]
	}
	return b
}

func SortRec() {

}
