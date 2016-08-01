package counter

import (
	"testing"

	"github.com/geotho/programming-problems/sorted-2d-array/utils"
)

var arrays = [][][]bool{
	linearArray(1),
	linearArray(2),
	linearArray(5),
	linearArray(10),
	linearArray(50),
	linearArray(100),
	linearArray(500),
	linearArray(1000),
	linearArray(5000),
	linearArray(10000),
	linearArray(30000),
	linearArray(50000),
	linearArray(70000),
	linearArray(80000),
	linearArray(90000),
	linearArray(100000),
}

func BenchmarkLinear1(b *testing.B) {
	bench(b, Linear{}, arrays[0])
}

func BenchmarkLinear2(b *testing.B) {
	bench(b, Linear{}, arrays[1])
}

func BenchmarkLinear5(b *testing.B) {
	bench(b, Linear{}, arrays[2])
}

func BenchmarkLinear10(b *testing.B) {
	bench(b, Linear{}, arrays[3])
}

func BenchmarkLinear50(b *testing.B) {
	bench(b, Linear{}, arrays[4])
}

func BenchmarkLinear100(b *testing.B) {
	bench(b, Linear{}, arrays[5])
}

func BenchmarkLinear500(b *testing.B) {
	bench(b, Linear{}, arrays[6])
}

func BenchmarkLinear1000(b *testing.B) {
	bench(b, Linear{}, arrays[7])
}

func BenchmarkLinear5000(b *testing.B) {
	bench(b, Linear{}, arrays[8])
}

func BenchmarkLinear10000(b *testing.B) {
	bench(b, Linear{}, arrays[9])
}

func BenchmarkLinear30000(b *testing.B) {
	bench(b, Linear{}, arrays[10])
}

func BenchmarkLinear50000(b *testing.B) {
	bench(b, Linear{}, arrays[11])
}

func BenchmarkLinear70000(b *testing.B) {
	bench(b, Linear{}, arrays[12])
}

func BenchmarkLinear80000(b *testing.B) {
	bench(b, Linear{}, arrays[13])
}

func BenchmarkLinear90000(b *testing.B) {
	bench(b, Linear{}, arrays[14])
}

func BenchmarkLinear100000(b *testing.B) {
	bench(b, Linear{}, arrays[15])
}

func BenchmarkBinarySearch1(b *testing.B) {
	bench(b, BinarySearch{}, arrays[0])
}

func BenchmarkBinarySearch2(b *testing.B) {
	bench(b, BinarySearch{}, arrays[1])
}

func BenchmarkBinarySearch5(b *testing.B) {
	bench(b, BinarySearch{}, arrays[2])
}

func BenchmarkBinarySearch10(b *testing.B) {
	bench(b, BinarySearch{}, arrays[3])
}

func BenchmarkBinarySearch50(b *testing.B) {
	bench(b, BinarySearch{}, arrays[4])
}

func BenchmarkBinarySearch100(b *testing.B) {
	bench(b, BinarySearch{}, arrays[5])
}

func BenchmarkBinarySearch500(b *testing.B) {
	bench(b, BinarySearch{}, arrays[6])
}

func BenchmarkBinarySearch1000(b *testing.B) {
	bench(b, BinarySearch{}, arrays[7])
}

func BenchmarkBinarySearch5000(b *testing.B) {
	bench(b, BinarySearch{}, arrays[8])
}

func BenchmarkBinarySearch10000(b *testing.B) {
	bench(b, BinarySearch{}, arrays[9])
}

func BenchmarkBinarySearch30000(b *testing.B) {
	bench(b, BinarySearch{}, arrays[10])
}

func BenchmarkBinarySearch50000(b *testing.B) {
	bench(b, BinarySearch{}, arrays[11])
}

func BenchmarkBinarySearch70000(b *testing.B) {
	bench(b, BinarySearch{}, arrays[12])
}

func BenchmarkBinarySearch80000(b *testing.B) {
	bench(b, BinarySearch{}, arrays[13])
}

func BenchmarkBinarySearch90000(b *testing.B) {
	bench(b, BinarySearch{}, arrays[14])
}

func BenchmarkBinarySearch100000(b *testing.B) {
	bench(b, BinarySearch{}, arrays[15])
}

func BenchmarkQuadTree1(b *testing.B) {
	bench(b, QuadTree{}, arrays[0])
}

func BenchmarkQuadTree2(b *testing.B) {
	bench(b, QuadTree{}, arrays[1])
}

func BenchmarkQuadTree5(b *testing.B) {
	bench(b, QuadTree{}, arrays[2])
}

func BenchmarkQuadTree10(b *testing.B) {
	bench(b, QuadTree{}, arrays[3])
}

func BenchmarkQuadTree50(b *testing.B) {
	bench(b, QuadTree{}, arrays[4])
}

func BenchmarkQuadTree100(b *testing.B) {
	bench(b, QuadTree{}, arrays[5])
}

func BenchmarkQuadTree500(b *testing.B) {
	bench(b, QuadTree{}, arrays[6])
}

func BenchmarkQuadTree1000(b *testing.B) {
	bench(b, QuadTree{}, arrays[7])
}

func BenchmarkQuadTree5000(b *testing.B) {
	bench(b, QuadTree{}, arrays[8])
}

func BenchmarkQuadTree10000(b *testing.B) {
	bench(b, QuadTree{}, arrays[9])
}

func BenchmarkQuadTree30000(b *testing.B) {
	bench(b, QuadTree{}, arrays[10])
}

func BenchmarkQuadTree50000(b *testing.B) {
	bench(b, QuadTree{}, arrays[11])
}

func BenchmarkQuadTree70000(b *testing.B) {
	bench(b, QuadTree{}, arrays[12])
}

func BenchmarkQuadTree80000(b *testing.B) {
	bench(b, QuadTree{}, arrays[13])
}

func BenchmarkQuadTree90000(b *testing.B) {
	bench(b, QuadTree{}, arrays[14])
}

func BenchmarkQuadTree100000(b *testing.B) {
	bench(b, QuadTree{}, arrays[15])
}

func BenchmarkSaddleback1(b *testing.B) {
	bench(b, Saddleback{}, arrays[0])
}

func BenchmarkSaddleback2(b *testing.B) {
	bench(b, Saddleback{}, arrays[1])
}

func BenchmarkSaddleback5(b *testing.B) {
	bench(b, Saddleback{}, arrays[2])
}

func BenchmarkSaddleback10(b *testing.B) {
	bench(b, Saddleback{}, arrays[3])
}

func BenchmarkSaddleback50(b *testing.B) {
	bench(b, Saddleback{}, arrays[4])
}

func BenchmarkSaddleback100(b *testing.B) {
	bench(b, Saddleback{}, arrays[5])
}

func BenchmarkSaddleback500(b *testing.B) {
	bench(b, Saddleback{}, arrays[6])
}

func BenchmarkSaddleback1000(b *testing.B) {
	bench(b, Saddleback{}, arrays[7])
}

func BenchmarkSaddleback5000(b *testing.B) {
	bench(b, Saddleback{}, arrays[8])
}

func BenchmarkSaddleback10000(b *testing.B) {
	bench(b, Saddleback{}, arrays[9])
}

func BenchmarkSaddleback30000(b *testing.B) {
	bench(b, Saddleback{}, arrays[10])
}

func BenchmarkSaddleback50000(b *testing.B) {
	bench(b, Saddleback{}, arrays[11])
}

func BenchmarkSaddleback70000(b *testing.B) {
	bench(b, Saddleback{}, arrays[12])
}

func BenchmarkSaddleback80000(b *testing.B) {
	bench(b, Saddleback{}, arrays[13])
}

func BenchmarkSaddleback90000(b *testing.B) {
	bench(b, Saddleback{}, arrays[14])
}

func BenchmarkSaddleback100000(b *testing.B) {
	bench(b, Saddleback{}, arrays[15])
}

func bench(b *testing.B, c Counter, arr [][]bool) {
	for i := 0; i < b.N; i++ {
		c.Count(arr)
	}
}

func linearArray(n int) [][]bool {
	arr := make([][]bool, 0, n)
	for i := 0; i < n; i++ {
		arr = append(arr, utils.MakeRow(n, n-i))
	}
	return arr
}
