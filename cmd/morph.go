package main

import (
	"flag"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"log"
	"morphology"
	"os"
	"strings"
)

func main() {
	input := flag.String("in", "/home/wotan/Pictures/lena_thresh.png",
		"image thresholded to be processed")
	output := flag.String("out", "/home/wotan/Pictures/lena_output.png",
		"image output")
	operation := flag.String("operation", "dilate",
		"morphology operation to be applied. Can be dilate or erode")
	kernel := flag.Int("k", 2, "kernel size")
	flag.Parse()
	log.Printf("in: %s, out: %s, operation: %s, kernel: %d", *input, *output, *operation, *kernel)

	file, err := os.Open(*input)
	defer file.Close()

	if err != nil {
		log.Fatal(err)
	}
	var img image.Image
	if strings.HasSuffix(*input, "jpg") || strings.HasSuffix(*input, "jpeg") {
		img, err = jpeg.Decode(file)
		if err != nil {
			log.Fatal(err)
		}
	} else if strings.HasSuffix(*input, "png") {
		img, err = png.Decode(file)
		if err != nil {
			log.Fatal(err)
		}
	}
	bounds := img.Bounds()
	src := make([][]int, bounds.Max.X)
	for i := range src {
		src[i] = make([]int, bounds.Max.Y)
	}

	for x := 0; x < bounds.Max.X; x++ {
		for y := 0; y < bounds.Max.Y; y++ {
			r, _, _, _ := img.At(x, y).RGBA()
			switch r {
			case 65535:
				src[x][y] = 255
			default:
				src[x][y] = 0
			}
		}
	}
	switch *operation {
	case "erode":
		log.Println("applying erode")
		src = morphology.Erode(src, *kernel)
	case "dilate":
		log.Println("applying dilate")
		src = morphology.Dilate(src, *kernel)
	default:
		log.Fatal("Operation not recognized")
	}

	outputFile, err := os.Create(*output)
	defer outputFile.Close()

	if err != nil {
		log.Fatal(err)
	}
	imgOut := image.NewGray(image.Rect(bounds.Min.X, bounds.Min.Y, bounds.Max.X, bounds.Max.Y))
	for x := 0; x < bounds.Max.X; x++ {
		for y := 0; y < bounds.Max.Y; y++ {
			switch src[x][y] {
			case 255:
				imgOut.Set(x, y, color.Gray{Y: 255})
			}
		}
	}
	if strings.HasSuffix(*output, "jpg") || strings.HasSuffix(*output, "jpeg") {
		err = jpeg.Encode(outputFile, imgOut, nil)
	} else if strings.HasSuffix(*output, "png") {
		err = png.Encode(outputFile, imgOut)
	}
	if err != nil {
		log.Fatal(err)
	}

}
