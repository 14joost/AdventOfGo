package main

import (
	"AdventOfGo/challenges"
	"AdventOfGo/utils"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	var day, part int
	var test bool
	var err error
	if len(os.Args) < 4 {
		day = 2
		part = 2
		test = true
		fmt.Printf("No sufficient arguments provided. Defaulting to day %d, puzzle %d and test input = %t\n", day, part, test)
	} else {
		day, err = strconv.Atoi(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
		part, err = strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatal(err)
		}
		test, err = strconv.ParseBool(os.Args[3])
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Printf("Solving challenge %d for day %d! Test input = %t\n", part, day, test)

	challengePicker(day, part, test)
}

func challengePicker(day int, part int, test bool) {

	switch day {
	case 1:
		if part == 1 {
			challenges.Day1Part1(utils.FileToStringSlice(day, test))
		} else if part == 2 {
			challenges.Day1Part2(utils.FileToStringSlice(day, test))
		}
	case 2:
		if part == 1 {
			challenges.Day2Part1(utils.FileToStringSlice(day, test))
		} else if part == 2 {
			challenges.Day2Part2(utils.FileToStringSlice(day, test))
		}
	default:
		fmt.Println("No challenges for that day")
	}
}
