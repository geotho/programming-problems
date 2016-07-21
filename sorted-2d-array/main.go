package main

import (
	"fmt"
	"math/rand"
	"time"
)

// genArray creates a random n*n array where all the 0s are in the top left
// hand corner and the remainder is filled with 1s.
func genArray(n int) [][]int {
	arr := make([][]int, 0, n)
	falses := rand.Intn(n)

	for i := 0; i < n; i++ {
		row := makeRow(n, falses)
		arr = append(arr, row)
		falses = rand.Intn(falses + 1)
	}

	return arr
}

func makeRow(length, zeroes int) []int {
	row := make([]int, length, length)
	for j := zeroes; j < length; j++ {
		row[j] = 1
	}
	return row
}

func countZeroes(row []int) int {
	if row[0] == 1 {
		return 0
	}

	if row[len(row)-1] == 0 {
		return len(row)
	}

	return countZeroesR(row, len(row), 0)
}

func countZeroesR(row []int, hi, lo int) int {
	mid := (hi - 1 + lo) / 2

	if row[mid] == 0 {
		if row[mid+1] == 1 {
			return mid + 1
		}
		return countZeroesR(row, hi, mid+1)
	}

	return countZeroesR(row, mid, lo)
}

func main() {
	rand.Seed(time.Now().Unix())
	arr := genArray(40)
	for _, row := range arr {
		fmt.Println(row)
	}
}
