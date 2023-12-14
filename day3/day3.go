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
	r2, _ := regexp.Compile("\\*")

	scanner := bufio.NewScanner(file)
	numbers := []int{}
	prevline := ""
	currentline := ""
	for scanner.Scan() {
		// read line by line
		nextline := scanner.Text()

		gearIndices := r2.FindAllStringSubmatchIndex(currentline, -1)
		numberIndices := r.FindAllStringSubmatchIndex(prevline, -1)
		numberIndices2 := r.FindAllStringSubmatchIndex(currentline, -1)
		numberIndices3 := r.FindAllStringSubmatchIndex(nextline, -1)

		for _, match := range gearIndices {
			startIdx := match[0]
			endIdx := match[1]

			// prevIdx := startIdx - 1
			// if prevIdx < 0 {
			// 	prevIdx = 0
			// }
			// nextIdx := endIdx + 1
			// if nextIdx > len(currentline) {
			// 	nextIdx = len(currentline)
			// }

			mathcingFunc := func(numidx []int) bool {
				return (numidx[1] == startIdx) ||
					(numidx[0] == endIdx) ||
					(numidx[0] <= startIdx && numidx[1] >= startIdx)
			}

			indexToValueFuncFunc := func(line string) func(indexes []int) int {
				return func(indexes []int) int {
					lol, _ := strconv.Atoi(line[indexes[0]:indexes[1]])
					return lol
				}
			}
			indexesToValuePrev := indexToValueFuncFunc(prevline)
			indexesToValueCurr := indexToValueFuncFunc(currentline)
			indexesToValueNext := indexToValueFuncFunc(nextline)

			numberOfNumbers := []int{}
			numberOfNumbers = append(numberOfNumbers, Map(where(numberIndices, mathcingFunc), indexesToValuePrev)...)

			numberOfNumbers = append(numberOfNumbers, Map(where(numberIndices2, mathcingFunc), indexesToValueCurr)...)
			numberOfNumbers = append(numberOfNumbers, Map(where(numberIndices3, mathcingFunc), indexesToValueNext)...)

			if len(numberOfNumbers) > 2 {
				fmt.Println("wtf")
			}
			if len(numberOfNumbers) == 2 {
				power := numberOfNumbers[0] * numberOfNumbers[1]
				numbers = append(numbers, power)
				fmt.Println(numberOfNumbers)
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

func Map[T any, V any](vs []T, f func(T) V) []V {
	vsm := make([]V, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

func where[T any](values []T, selector func(T) bool) []T {
	result := []T{}
	for _, item := range values {
		if selector(item) {
			result = append(result, item)
		}
	}
	return result

}

func sum(numbers []int) int {
	x := 0
	for _, n := range numbers {
		x = x + n
	}
	return x
}
