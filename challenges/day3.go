package challenges

import (
	"AdventOfGo/utils"
	"fmt"
)

const mulRegEx string = "mul\\([0-9]*,[0-9]*\\)"

func Day3Part1(stringSlice []string) {
	score := 0
	for _, line := range stringSlice {
		//fmt.Printf("Line %d: %s\n", i, line)
		var muls = utils.RegexFetcher(line, mulRegEx)
		fmt.Printf("All multiply matches: %s. \n", muls)

		for _, mul := range muls {
			strippedMul := mul[4 : len(mul)-1]
			fmt.Printf("Mul: %s\n", strippedMul)
			values := utils.StringSliceToIntSlice(utils.Split(strippedMul, ","))
			score += multiply(values[0], values[1])
		}
	}

	fmt.Println("Score: ", score)
}

func Day3Part2(stringSlice []string) {
	score := 0
	do := "do()"
	dont := "don't()"

	for _, line := range stringSlice {
		//fmt.Printf("Line %d: %s\n", i, line)
		doIndices := utils.SubstringIndexFetcher(line, do)
		doIndices = append([]int{0}, doIndices...)
		dontIndices := utils.SubstringIndexFetcher(line, dont)
		fmt.Printf("Do indices: %d. \n", doIndices)
		fmt.Printf("Don't indices: %d. \n", dontIndices)
		doDontSlice := trimDoDont(doIndices, dontIndices)
		fmt.Printf("DoDont slice: %d. \n", doDontSlice)
		//var muls = utils.RegexFetcher(line, mulRegEx)
		//fmt.Printf("All multiply matches: %s. \n", muls)
		//
		//for _, mul := range muls {
		//	strippedMul := mul[4 : len(mul)-1]
		//	fmt.Printf("Mul: %s\n", strippedMul)
		//	values := utils.StringSliceToIntSlice(utils.Split(strippedMul, ","))
		//	score += multiply(values[0], values[1])
		//}
	}

	fmt.Println("Score: ", score)
}

func multiply(a int, b int) int {
	fmt.Printf("Multiplying %d and %d\n", a, b)
	return a * b
}

func trimDoDont(doSlice []int, dontSlice []int) []int {
	doDontSlice := make([]int, 0)
	doDontSlice = append(doDontSlice, 0)
	doCounter, dontCounter := 1, 0
	done, do := false, false
	for !done {
		if do {
			if doCounter == len(doSlice) {
				do = false
			} else if doSlice[doCounter] > doDontSlice[len(doDontSlice)-1] && dontCounter < len(dontSlice) {
				doDontSlice = append(doDontSlice, doSlice[doCounter])
				doCounter++
				do = false
			}
		} else {
			if dontCounter == len(dontSlice) {
				do = true
			} else if dontSlice[dontCounter] > doDontSlice[len(doDontSlice)-1] && doCounter < len(doSlice) {
				doDontSlice = append(doDontSlice, dontSlice[dontCounter])
				dontCounter++
				do = true
			}
		}
		if doCounter >= len(doSlice) && dontCounter >= len(dontSlice) {
			done = true
		}
	}
	return doDontSlice
}
