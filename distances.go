package morphology

import "math"

// Manhattan create a matrix with the distance based on value
func Manhattan(src [][]int, value int) [][]int {
	// transverse from top left to bottom right
	for i := 0; i < len(src); i++ {
		for j := 0; j < len(src[i]); j++ {
			if src[i][j] == value {
				// first pass and pixel was on, it gets a zero
				src[i][j] = 0
			} else {
				// pixel was off
				// It is at most the sum of the lengths of the array
				// away from a pixel that is on
				src[i][j] = len(src) + len(src[i])
				// or one more than the pixel to the north
				if i > 0 {
					src[i][j] = int(math.Min(float64(src[i][j]), float64(src[i-1][j]+1)))
				}
				// or one more than the pixel to the west
				if j > 0 {
					src[i][j] = int(math.Min(float64(src[i][j]), float64(src[i][j-1]+1)))
				}
			}
		}
	}
	// traverse from bottom right to top left
	for i := len(src) - 1; i >= 0; i-- {
		for j := len(src[i]) - 1; j >= 0; j-- {
			// either what we had on the first pass
			// or one more than the pixel to the south
			if i+1 < len(src) {
				src[i][j] = int(math.Min(float64(src[i][j]), float64(src[i+1][j]+1)))
			}
			// or one more than the pixel to the east
			if j+1 < len(src[i]) {
				src[i][j] = int(math.Min(float64(src[i][j]), float64(src[i][j+1]+1)))
			}
		}
	}
	return src
}
