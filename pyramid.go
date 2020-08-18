package pyramid

import (
	"fmt"
	"math"
)

// --- Directions
// Write a function that accepts a positive number N.
// The function should print out a pyramid shape
// with N levels using the # character.  Make sure the
// pyramid has spaces on both the left *and* right hand sides
// --- Examples
//   pyramid(1)
//       '#'
//   pyramid(2)
//       ' # '
//       '###'
//   pyramid(3)
//       '  #  '
//       ' ### '
//       '#####'
func pyramidRecursion(n int, rowOptional ...int) (str string) {
	row := 0
	cols := 2*n - 1
	midpoint := int(math.Floor(float64(cols) / 2.0))

	if len(rowOptional) > 0 {
		row = rowOptional[0]
	}

	if row >= n {
		return str
	}

	for i := 0; i < cols; i++ {
		//if midpoint-row <=
		if i >= midpoint-row && i <= midpoint+row {
			//if i <= row {
			str += "#"
		} else {
			str += " "
		}
	}
	if row < (n - 1) {
		str += "\n"
	}
	str += pyramidRecursion(n, row+1)
	return

}

func pyramid(n int) string {
	cols := 2*n - 1
	midpoint := int(math.Floor(float64(cols) / 2.0))
	str := ""
	for i := 0; i < n; i++ {
		for j := 0; j < cols; j++ {
			if j >= midpoint-i && j <= midpoint+i {
				str += "#"
			} else {
				str += " "
			}
		}
		if i < n-1 {
			str += "\n"
		}
	}
	fmt.Println(str)
	return str
}
