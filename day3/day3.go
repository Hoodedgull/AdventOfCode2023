package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	file, err := os.Open("./input-day-03.txt")
	check(err)

	defer file.Close()

	r, _ := regexp.Compile("\\d+")
	r2, _ := regexp.Compile("[^\\d\\.]")

	scanner := bufio.NewScanner(file)
	numbers := []int{}
	prevline := ""
	currentline := ""
	for scanner.Scan() {
		// read line by line
		nextline := scanner.Text()

		matcheIndices := r.FindAllStringSubmatchIndex(currentline, -1)
		symbolIndices := r2.FindAllStringSubmatchIndex(prevline, -1)
		symbolIndices2 := r2.FindAllStringSubmatchIndex(currentline, -1)
		symbolIndices3 := r2.FindAllStringSubmatchIndex(nextline, -1)

		for _, match := range matcheIndices {
			startIdx := match[0]
			endIdx := match[1]

			prevIdx := startIdx - 1
			if prevIdx < 0 {
				prevIdx = 0
			}
			nextIdx := endIdx + 1
			if nextIdx > len(currentline) {
				nextIdx = len(currentline)
			}
			isPartNum := false
			if includes(symbolIndices, func(symbol []int) bool {
				return symbol[0] >= prevIdx && symbol[0] < nextIdx
			}) {
				isPartNum = true
			}
			if includes(symbolIndices2, func(symbol []int) bool {
				return symbol[0] >= prevIdx && symbol[0] < nextIdx
			}) {
				isPartNum = true
			}
			if includes(symbolIndices3, func(symbol []int) bool {
				return symbol[0] >= prevIdx && symbol[0] < nextIdx
			}) {
				isPartNum = true
			}

			if isPartNum {
				matchedNumber, _ := strconv.Atoi(currentline[startIdx:endIdx])
				numbers = append(numbers, matchedNumber)
			}

		}

		prevline = currentline
		currentline = nextline

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(numbers)
	fmt.Println(sum(numbers))

}

func includes[T any](values []T, selector func(T) bool) bool {
	for _, item := range values {
		if selector(item) {
			return true
		}
	}
	return false

}

func sum(numbers []int) int {
	x := 0
	for _, n := range numbers {
		x = x + n
	}
	return x
}
