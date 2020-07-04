package morphology

import "testing"

func TestManhattan(t *testing.T) {
	a := [][]int{
		{
			0, 0, 0,
		},
		{
			0, 255, 0,
		},
		{
			0, 0, 0,
		},
	}

	result := Manhattan(a, 255)
	expected := [][]int{
		{
			2, 1, 2,
		},
		{
			1, 0, 1,
		},
		{
			2, 1, 2,
		},
	}
	for i := 0; i < len(expected); i++ {
		for j := 0; j < len(expected[i]); j++ {

			if expected[i][j] != result[i][j] {
				t.Errorf("value expected: %d found: %d, i: %d, j: %d", expected[i][j], result[i][j], i, j)
			}
		}
	}
}
