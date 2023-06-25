package main

func getRow(rowIndex int) []int {
	prev := []int{1}
	row := make([]int, 0)
	for i := 1; i < rowIndex+1; i++ {
		row = row[:0]
		row = append(row, 1)
		for j := 1; j < len(prev); j++ {
			row = append(row, prev[j]+prev[j-1])
		}
		row = append(row, 1)
		prev = make([]int, len(row))
		copy(prev, row)
	}
	return prev
}
