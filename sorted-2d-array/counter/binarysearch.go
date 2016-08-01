package counter

import "github.com/geotho/programming-problems/sorted-2d-array/utils"

// BinarySearch is a Counter that does a binary search on each row to find
// the number of zeroes in the row.
type BinarySearch struct{}

var _ Counter = (*BinarySearch)(nil)

func (bs BinarySearch) Count(arr [][]bool) int {
	total := 0
	zeroes := len(arr)
	for _, row := range arr {
		// Limit the row by the previous number of zeroes
		zeroes := utils.CountZeroes(row[:zeroes])
		if zeroes == 0 {
			break
		}
		total += zeroes
	}
	return total
}
