package challenges

import (
	"AdventOfGo/utils"
	"fmt"
	"math"
	"strings"
)

func Day4Part1(horizontalSlice []string) {
	//count := 0
	xMax := len(horizontalSlice[0])
	yMax := len(horizontalSlice)
	verticalSlice := make([]string, len(horizontalSlice[0]))
	diagonalSlice := make([]string, len(horizontalSlice)+len(verticalSlice)-1)

	for i := range verticalSlice {
		for j, horizontalLetter := range utils.SplitEachChar(horizontalSlice[i]) {
			verticalSlice[j] += horizontalLetter
		}
	}

	for z := range diagonalSlice {
		x1 := int(math.Max(float64(xMax-1-z), 0))
		y1 := int(math.Max(float64(z-yMax+1), 0))
		x2 := int(math.Min(float64(2*(xMax-1)-z), float64(xMax-1)))
		y2 := int(math.Min(float64(z), float64(yMax-1)))
		fmt.Printf("1: %d, %d\n", x1, y1)
		fmt.Printf("2: %d, %d\n", x2, y2)
		coordinates := getDiagonalCoordinates(x1, y1, x2, y2)
		utils.PrintSliceAny(coordinates)

	}

	fmt.Printf("Horizontal slice:  \n")
	utils.PrintSlice(horizontalSlice)
	fmt.Println()
	fmt.Printf("Vertical slice: \n")
	utils.PrintSlice(verticalSlice)
	fmt.Println()
	//fmt.Printf("Diagonal slice: \n")
	//utils.PrintSlice(diagonalSlice)
	//fmt.Println()

}

func Day4Part2(stringSlice []string) {
	//score := 0
	//for _, line := range stringSlice {
	//	score += checkLineAnagram(line)
	//}
	//println("Score: ", score)
}

func checkLToR(horizontal string) int {
	return strings.Count(horizontal, "XMAS")
}

func checkRToL(horizontal string) int {
	return strings.Count(horizontal, "SAMX")
}

func getDiagonalCoordinates(x1 int, y1 int, x2 int, y2 int) []utils.Coordinate {
	coordinates := make([]utils.Coordinate, 0)
	coordinates = append(coordinates, utils.Coordinate{X: x1, Y: y1})
	coordinates = append(coordinates, utils.Coordinate{X: x2, Y: y2})

	for i := 0; i < y2; i++ {
		coordinates = append(coordinates, utils.Coordinate{X: x1 - i, Y: y1 + i})
	}
	return utils.RemoveDuplicates(coordinates)
}
