package utils

import "math/rand"

// CountZeroes counts the number of zeroes in a row using a recursive binary search.
func CountZeroes(row []bool) int {
	if row[0] {
		return 0
	}

	if !row[len(row)-1] {
		return len(row)
	}

	hi := len(row)
	lo := 0

	for {
		mid := (hi - 1 + lo) / 2
		if !row[mid] {
			if row[mid+1] {
				return mid + 1
			}
			lo = mid + 1
			continue
		}
		hi = mid
	}
}

// GenArray creates a random n*n array where all the falses are in the top left
// hand corner and the remainder is filled with trues.
func GenArray(n int) [][]bool {
	arr := make([][]bool, 0, n)
	falses := rand.Intn(n)

	for i := 0; i < n; i++ {
		row := MakeRow(n, falses)
		arr = append(arr, row)
		falses = rand.Intn(falses + 1)
	}

	return arr
}

// MakeRow makes a row of length `length` filled with `zeroes` falses.
func MakeRow(length, zeroes int) []bool {
	row := make([]bool, length, length)
	for j := zeroes; j < length; j++ {
		row[j] = true
	}
	return row
}

// Str returns the matrix as a string of 1s and 0s.
func Str(arr [][]bool) string {
	str := ""
	for _, row := range arr {
		for _, v := range row {
			if v {
				str += "1"
			} else {
				str += "0"
			}
		}
		str += "\n"
	}
	return str
}
