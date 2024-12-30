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

func trimDoDont(input1, input2 []int) []int {
	var output []int
	i, j := 0, 0 // Indices for input1 and input2
	len1, len2 := len(input1), len(input2)
	turn := 0  // 0 for input1's turn, 1 for input2's turn
	prev := -1 // Initialize previous value

	for {
		if turn%2 == 0 { // Input1's turn
			// Find the next element in input1 greater than 'prev'
			for i < len1 && input1[i] <= prev {
				i++
			}
			if i >= len1 {
				// Cannot pick from input1; stop the process
				break
			}
			output = append(output, input1[i])
			prev = input1[i]
			i++
		} else { // Input2's turn
			// Find the next element in input2 greater than 'prev'
			for j < len2 && input2[j] <= prev {
				j++
			}
			if j >= len2 {
				// Cannot pick from input2; stop the process
				break
			}
			output = append(output, input2[j])
			prev = input2[j]
			j++
		}
		turn++
	}
	return output
}
