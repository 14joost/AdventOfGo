package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
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

func Split(inputString string, splitChar string) []string {
	if splitChar == "" || splitChar == " " {
		return SplitByWhitespace(inputString)
	}
	return strings.Split(inputString, splitChar)
}

func SplitByWhitespace(inputString string) []string {
	return strings.Fields(inputString)
}

func StringSliceToIntSlice(stringSlice []string) []int {
	var intSlice = make([]int, 0)
	for _, s := range stringSlice {
		i, _ := strconv.Atoi(s)
		intSlice = append(intSlice, i)
	}
	return intSlice
}

func RemoveElementFromSliceByIndex[T any](slice []T, s int) []T {
	x := append(slice[:s], slice[s+1:]...)
	return x
}

func RegexFetcher(line string, regEx string) []string {
	return regexp.MustCompile(regEx).FindAllString(line, -1)
}
