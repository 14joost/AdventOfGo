package challenges

import (
	"AdventOfGo/utils"
	"fmt"
)

const regEx string = "mul\\([0-9]*,[0-9]*\\)"

func Day3Part1(stringSlice []string) {
	score := 0
	for _, line := range stringSlice {
		//fmt.Printf("Line %d: %s\n", i, line)
		var muls = utils.RegexFetcher(line, regEx)
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
	// TODO
}

func multiply(a int, b int) int {
	fmt.Printf("Multiplying %d and %d\n", a, b)
	return a * b
}
