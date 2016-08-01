package counter

import (
	"github.com/geotho/programming-problems/sorted-2d-array/utils"
)

// QuadTree is a Counter that partitions the matrix into quadrants. The top left
// contains all 0s, the bottom right contains all 1s and the other two are recursed into.
type QuadTree struct{}

type point struct{ x, y int }

var _ Counter = (*QuadTree)(nil)

// Count counts 0s in the matrix by recursively splitting the matrix into quadrants.
func (qt QuadTree) Count(arr [][]bool) int {
	n := len(arr)
	return countR(arr, point{0, 0}, point{n - 1, n - 1})
}

func countR(arr [][]bool, topLeft, bottomRight point) int {
	if topLeft.x < 0 || topLeft.y < 0 {
		return 0
	}

	if n := len(arr); bottomRight.x > n || bottomRight.y > n {
		return 0
	}

	if bottomRight.x < topLeft.x || bottomRight.y < topLeft.y {
		return 0
	}

	if topLeft == bottomRight {
		if arr[topLeft.y][topLeft.x] {
			return 0
		}
		return 1
	}

	if topLeft.y == bottomRight.y {
		return utils.CountZeroes(arr[topLeft.y][topLeft.x : bottomRight.x+1])
	}

	mid := (topLeft.y + bottomRight.y) / 2
	midRow := arr[mid]

	zeroes := utils.CountZeroes(midRow) - 1
	m := point{zeroes, mid}

	return (m.x-topLeft.x+1)*(m.y-topLeft.y+1) +
		countR(arr, point{topLeft.x, m.y + 1}, point{m.x, bottomRight.y}) +
		countR(arr, point{m.x + 1, topLeft.y}, point{bottomRight.x, m.y - 1})
}
