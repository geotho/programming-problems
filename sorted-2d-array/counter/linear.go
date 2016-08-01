package counter

// Linear is a Counter that iterates through each value in the 2D array and
// increments a counter for each false it finds.
// It early-exits a row when it finds its first true.
type Linear struct{}

var _ Counter = (*Linear)(nil)

func (l Linear) Count(arr [][]bool) int {
	total := 0
	for _, row := range arr {
		for _, v := range row {
			// if we hit a true, we can stop with this row
			if v {
				break
			}
			total++
		}
	}
	return total
}
