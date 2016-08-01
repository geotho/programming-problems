package utils

import "testing"

func TestCountZeroes(t *testing.T) {
	maxLen := 100

	for i := 1; i <= maxLen; i++ {
		for j := 0; j <= i; j++ {
			row := MakeRow(i, j)
			if z := CountZeroes(row); z != j {
				t.Errorf("array=%v, actual=%d, expected=%d", row, z, j)
			}
		}
	}
}
