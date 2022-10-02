package day20

import (
	"errors"
	"fmt"
	"strings"
)

type pixel uint8

func parseInput(input string) ([]pixel, [][]pixel) {
	blocks := strings.Split(input, "\n\n")

	algoStr := blocks[0]
	algo := make([]pixel, len(algoStr))

	for i, char := range algoStr {
		if char == '#' {
			algo[i] = pixel(1)
		} else {
			algo[i] = pixel(0)
		}
	}

	imageStr := strings.Split(blocks[1], "\n")
	image := make([][]pixel, len(imageStr))

	for i, line := range imageStr {
		imageLine := make([]pixel, len(line))

		for j, char := range line {
			if char == '#' {
				imageLine[j] = pixel(1)
			} else {
				imageLine[j] = pixel(0)
			}
		}

		image[i] = imageLine
	}

	return algo, image
}

func increaseImage(image [][]pixel, fill pixel) [][]pixel {
	height, width := len(image), len(image[0])

	newImage := make([][]pixel, height+2)

	newImage[0] = make([]pixel, width+2)
	for i := range newImage[0] {
		newImage[0][i] = fill
	}

	newImage[height+1] = make([]pixel, width+2)
	for i := range newImage[height+1] {
		newImage[height+1][i] = fill
	}

	for i, line := range image {
		newImageLine := make([]pixel, width+2)
		newImageLine[0] = fill
		newImageLine[width+1] = fill

		for j, pix := range line {
			newImageLine[j+1] = pix
		}

		newImage[i+1] = newImageLine
	}

	return newImage
}

func findAlgoIndex(image [][]pixel, i, j int) int {
	algoIndex := 0
	powOf2 := 1

	for k := i + 1; k >= i-1; k-- {
		for l := j + 1; l >= j-1; l-- {
			algoIndex += int(image[k][l]) * powOf2
			powOf2 *= 2
		}
	}

	return algoIndex
}

func nextStep(algo []pixel, image [][]pixel, fill, nextFill pixel) [][]pixel {
	height, width := len(image), len(image[0])

	increased := increaseImage(image, fill)
	newImage := increaseImage(image, nextFill)

	for i := 1; i <= height; i++ {
		for j := 1; j <= width; j++ {
			algoIndex := findAlgoIndex(increased, i, j)
			newImage[i][j] = algo[algoIndex]
		}
	}

	return newImage
}

func applyAlgo(algo []pixel, baseImage [][]pixel, times int) [][]pixel {
	if algo[0] == pixel(1) && algo[len(algo)-1] == pixel(1) {
		panic(errors.New("invalid algorithm"))
	}

	image := increaseImage(baseImage, pixel(0))
	fill := pixel(0)

	for step := 0; step < times; step++ {
		var nextFill pixel

		if fill == pixel(0) {
			nextFill = algo[0]
		} else {
			nextFill = algo[len(algo)-1]
		}

		image = nextStep(algo, image, fill, nextFill)
		fill = nextFill
	}

	return image
}

func countPixels(image [][]pixel) int {
	count := 0

	for _, line := range image {
		for _, pix := range line {
			count += int(pix)
		}
	}

	return count
}

func Run(input string) {
	algo, baseImage := parseInput(input)

	fmt.Println("Lit pixels after 2 times:", countPixels(applyAlgo(algo, baseImage, 2)))
	fmt.Println("Lit pixels after 2 times:", countPixels(applyAlgo(algo, baseImage, 50)))
}
