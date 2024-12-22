package main

import (
	"AdventOfGo/challenges"
	"AdventOfGo/utils"
	"fmt"
	"log"
	"os"
	"strconv"
)

func init() {
	// This function will be executed before the main function
	// It is used to initialize the program
}
func main() {
	var day int
	var test bool
	var err error
	if len(os.Args) < 2 {
		fmt.Println("No arguments provided. Defaulting to day 1 and test input")
		day = 1
		test = true
	} else {
		day, err = strconv.Atoi(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
		test, err = strconv.ParseBool(os.Args[2])
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Printf("Solving challenge for day %d! Test input = %t\n", day, test)

	challenges.Day1(utils.FileToStringSlice(day, test))
}
