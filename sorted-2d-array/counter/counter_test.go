package counter

import (
	"testing"

	"github.com/geotho/programming-problems/sorted-2d-array/utils"
)

var (
	linear = Linear{}
	binary = BinarySearch{}
	quad   = QuadTree{}
	saddle = Saddleback{}
)

func TestCounters(t *testing.T) {
	itr := 1
	maxSize := 1000

	for i := 1; i <= maxSize; i++ {
		for j := 0; j < itr; j++ {
			arr := utils.GenArray(i)
			linearCount := linear.Count(arr)
			binaryCount := binary.Count(arr)
			quadCount := quad.Count(arr)
			saddleCount := saddle.Count(arr)

			if linearCount != binaryCount {
				t.Errorf("binary count wrong: array=\n%s, binary=%d, linear=%d", utils.Str(arr), binaryCount, linearCount)
				return
			}
			if linearCount != quadCount {
				t.Errorf("quad count wrong: array=\n%s, quad=%d, linear=%d", utils.Str(arr), quadCount, linearCount)
				return
			}
			if linearCount != saddleCount {
				t.Errorf("saddle count wrong: array=\n%s, saddle=%d, linear=%d", utils.Str(arr), saddleCount, linearCount)
				return
			}
		}
	}
}
