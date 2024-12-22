package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func FileToStringSlice(day int, test bool) []string {
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

	return inputList
}

func SplitByWhitespace(inputString string) []string {
	var stringSliceNoWhitespace = make([]string, 0)
	for _, s := range strings.Split(inputString, " ") {
		if s != " " && s != "" {
			stringSliceNoWhitespace = append(stringSliceNoWhitespace, s)
		}
	}
	return stringSliceNoWhitespace
}

func StringSliceToIntSlice(stringSlice []string) []int {
	var intSlice = make([]int, 0)
	for _, s := range stringSlice {
		i, _ := strconv.Atoi(s)
		intSlice = append(intSlice, i)
	}
	return intSlice
}
