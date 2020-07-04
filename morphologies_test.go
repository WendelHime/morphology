package morphology

import (
	"math"
	"math/rand"
	"os"
	"testing"
)

func TestDilate(t *testing.T) {
	src := [][]int{
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
	expected := [][]int{
		{
			0, 255, 0,
		},
		{
			255, 255, 255,
		},
		{
			0, 255, 0,
		},
	}

	result := Dilate(src, 2)
	for i := 0; i < len(expected); i++ {
		for j := 0; j < len(expected[i]); j++ {
			if result[i][j] != expected[i][j] {
				t.Errorf("value expected: %d found: %d, i: %d, j: %d", expected[i][j], result[i][j], i, j)
			}
		}
	}
}

func TestErode(t *testing.T) {
	src := [][]int{
		{
			255, 255, 255,
		},
		{
			255, 0, 255,
		},
		{
			255, 255, 255,
		},
	}
	expected := [][]int{
		{
			255, 0, 255,
		},
		{
			0, 0, 0,
		},
		{
			255, 0, 255,
		},
	}

	result := Erode(src, 2)
	for i := 0; i < len(expected); i++ {
		for j := 0; j < len(expected[i]); j++ {
			if result[i][j] != expected[i][j] {
				t.Errorf("value expected: %d found: %d, i: %d, j: %d", expected[i][j], result[i][j], i, j)
			}
		}
	}
}

func benchmarkDilate(src [][]int, kernelSize int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		Dilate(src, kernelSize)
	}
}

func BenchmarkTestDilate2M(b *testing.B)  { benchmarkDilate(bigMat, 5, b) }
func BenchmarkTestDilate64k(b *testing.B) { benchmarkDilate(smallMat, 5, b) }

func benchmarkErode(src [][]int, kernelSize int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		Erode(src, kernelSize)
	}
}

func BenchmarkTestErode2M(b *testing.B)  { benchmarkErode(bigMat, 5, b) }
func BenchmarkTestErode64k(b *testing.B) { benchmarkErode(smallMat, 5, b) }

const (
	bigMatSize   = 1 << 21
	smallMatSize = 1 << 16
)

var (
	bigMat   [][]int
	smallMat [][]int
)

func makeMat(size int) [][]int {
	xMax := int(math.Sqrt(float64(size)))
	yMax := xMax
	mat := make([][]int, xMax)
	for x := 0; x < xMax; x++ {
		mat[x] = make([]int, yMax)
		for y := 0; y < yMax; y++ {
			if rand.Float32() < 0.5 {
				mat[x][y] = 0
			} else {
				mat[x][y] = 255
			}
		}
	}
	return mat
}

func TestMain(m *testing.M) {
	bigMat = makeMat(bigMatSize)
	smallMat = makeMat(smallMatSize)
	os.Exit(m.Run())
}
