package morphology

// Dilate turns on pixels which were near pixels that were on originally,
// thereby tickening items in the image
func Dilate(src [][]int, k int) [][]int {
	src = Manhattan(src, 255)
	for i := 0; i < len(src); i++ {
		for j := 0; j < len(src[i]); j++ {
			if src[i][j] < k {
				src[i][j] = 255
			} else {
				src[i][j] = 0
			}
		}
	}
	return src
}

// Erode it turns off pixels which were near pixels thar were off originally,
// thereby eating items in the image
func Erode(src [][]int, k int) [][]int {
	src = Manhattan(src, 0)
	for i := 0; i < len(src); i++ {
		for j := 0; j < len(src[i]); j++ {
			if src[i][j] < k {
				src[i][j] = 0
			} else {
				src[i][j] = 255
			}
		}
	}
	return src
}
