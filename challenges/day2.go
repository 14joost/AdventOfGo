package challenges

import (
	"AdventOfGo/utils"
	"fmt"
	"math"
)

func Day2Part1(stringSlice []string) {
	score := 0
	for _, line := range stringSlice {
		score += checkLine(utils.StringSliceToIntSlice(utils.SplitByWhitespace(line)))
	}
	fmt.Println("Score: ", score)
}

func Day2Part2(stringSlice []string) {
	score := 0
	for _, line := range stringSlice {

		report := utils.StringSliceToIntSlice(utils.SplitByWhitespace(line))

		if checkLine(report) == 1 {
			score++
			continue
		}

		for i := 0; i < len(report); i++ {
			reportCopy := make([]int, len(report))
			copy(reportCopy, report)
			dampenedReport := utils.RemoveElementFromSliceByIndex(reportCopy, i)
			if checkLine(dampenedReport) == 1 {
				score++
				break
			}
		}

		//score += checkLine()
	}
	fmt.Println("Score: ", score)
}

func checkLine(report []int) int {
	var asc, desc bool
	for i := 0; i < len(report)-1; i++ {
		diff := report[i] - report[i+1]
		absDiff := int(math.Abs(float64(diff)))
		if absDiff < 1 || absDiff > 3 {
			return 0
		}
		if diff < 0 { // ascending
			asc = true
		} else { // descending
			desc = true
		}
		if asc && desc {
			return 0
		}
	}
	return 1
}
