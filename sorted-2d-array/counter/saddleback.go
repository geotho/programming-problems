package counter

type Saddleback struct{}

var _ Counter = (*Saddleback)(nil)

// Count starts in the top right corner. If it sees a one, it moves left.
// If it sees a 0, it adds the number of zeroes in the row and moves down.
func (s Saddleback) Count(arr [][]bool) int {
	x := len(arr) - 1
	y := 0
	total := 0
	for x >= 0 && y < len(arr) {
		if arr[y][x] {
			x--
		} else {
			total += x + 1
			y++
		}
	}
	return total
}
