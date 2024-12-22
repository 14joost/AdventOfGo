package challenges

import (
	"AdventOfGo/utils"
	"fmt"
	"log"
	"math"
	"sort"
	"strconv"
)

func Day1Part1(inputList []string) {

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

func Day1Part2(inputList []string) {
	//fmt.Print(inputList)
	leftMap := make(map[int]int)
	rightMap := make(map[int]int)
	for _, line := range inputList {
		sliceNoWhitespace := utils.SplitByWhitespace(line)
		left, err := strconv.Atoi(sliceNoWhitespace[0])
		if err != nil {
			log.Fatal(err)
		}
		right, err := strconv.Atoi(sliceNoWhitespace[1])
		if err != nil {
			log.Fatal(err)
		}
		checkAndUpdateMap(leftMap, left)
		checkAndUpdateMap(rightMap, right)
	}

	fmt.Println("Left map: ", leftMap)
	fmt.Println("Right map: ", rightMap)

	similarityScore := 0

	for lKey, lValue := range leftMap {
		similarityScore += lKey * rightMap[lKey] * lValue
	}
	fmt.Println("Similarity score: ", similarityScore)
}

func checkAndUpdateMap(m map[int]int, key int) {
	if _, ok := m[key]; ok {
		m[key]++
	} else {
		m[key] = 1
	}
}
