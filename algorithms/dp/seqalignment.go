package dp

func SequenceAlignment(x string, y string) int {
	d, g := 1, 2 // mismtach and gap penalties
	m, n := len(x), len(y)
	matrix := matrix(m, n)
	for rows := range matrix {
		matrix[rows][0] = rows * d
		for cols := range matrix[rows] {
			matrix[0][cols] = cols * d
		}
	}

	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[i-1]); j++ {
			t1 := d + matrix[i-1][j-1]
			t2 := g + matrix[i-1][j]
			t3 := g + matrix[i][j-1]
			min := min(t1, min(t2, t3))
			matrix[i][j] = min
		}
	}
	// Print matrix:
	// for rows := range m {
	// 	fmt.Println()
	// 	for cols := range m[rows] {
	// 		fmt.Printf("[ %v ] ", m[rows][cols])
	// 	}
	// }
	return matrix[m][n]
}

func min(x int, y int) int {
	if x < y {
		return x
	}
	return y
}
