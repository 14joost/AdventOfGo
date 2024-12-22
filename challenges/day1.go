package challenges

import (
	"AdventOfGo/utils"
	"fmt"
	"math"
	"sort"
)

func Day1(inputList []string) {

	//fmt.Print(inputList)
	var leftSlice, rightSlice []string
	for _, line := range inputList {
		sliceNoWhitespace := utils.SplitByWhitespace(line)
		left := sliceNoWhitespace[0]
		right := sliceNoWhitespace[1]
		leftSlice = append(leftSlice, left)
		rightSlice = append(rightSlice, right)
	}
	fmt.Println("Left slice: ", leftSlice)
	fmt.Println("Right slice: ", rightSlice)

	leftSliceInt := utils.StringSliceToIntSlice(leftSlice)
	rightSliceInt := utils.StringSliceToIntSlice(rightSlice)

	sort.Sort(sort.IntSlice(leftSliceInt))
	sort.Sort(sort.IntSlice(rightSliceInt))

	fmt.Println("Left slice int: ", leftSliceInt)
	fmt.Println("Right slice int: ", rightSliceInt)

	var diffSum int
	for i := 0; i < len(leftSliceInt); i++ {
		//diffSum += math.Abs(rightSliceInt[i] - leftSliceInt[i])
		diffSum += int(math.Abs(float64(rightSliceInt[i] - leftSliceInt[i])))
	}
	fmt.Println("Sum of differences: ", diffSum)
}
