package main

import (
	"AdventOfGo/challenges"
	"bufio"
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
	day, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	test, err := strconv.ParseBool(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Solving challenge for day %d! Test input = %t\n", day, test)

	inputFile := fmt.Sprintf("input_day%d", day)
	inputList := make([]string, 0)
	var testStr string
	if test {
		testStr = "_test"
	}
	file, err := os.Open(fmt.Sprintf("inputFiles/%s%s.txt", inputFile, testStr))
	scanner := bufio.NewScanner(file)
	if err != nil {
		log.Fatal(err)
	}
	for scanner.Scan() {
		text := scanner.Text()
		inputList = append(inputList, text)
		//fmt.Println(text)
	}
	challenges.Day1(inputList)
}
