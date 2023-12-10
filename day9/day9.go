package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	file, err := os.Open("./input-day-09.txt")
	check(err)

	defer file.Close()

	scanner := bufio.NewScanner(file)
	extrapolatedValues := []int{}

	for scanner.Scan() {
		// read line by line
		line := scanner.Text()

		data := Map(strings.Split(line, " "), atoi)

		diffLines := [][]int{data}
		//diffLines = append(diffLines, makeDiffLine(data))
		for !every(diffLines[len(diffLines)-1], func(n int) bool { return n == 0 }) {
			diffLines = append(diffLines, makeDiffLine(diffLines[len(diffLines)-1]))
		}
		fmt.Println("made diff")

		diffLines[len(diffLines)-1] = prepend(diffLines[len(diffLines)-1], 0)
		for i := len(diffLines) - 2; i >= 0; i-- {
			newValue := diffLines[i][0] - diffLines[i+1][0]
			diffLines[i] = prepend(diffLines[i], newValue)
		}

		fmt.Println("extrapolated")
		extrapolatedValues = append(extrapolatedValues, diffLines[0][0])

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Sum:")
	fmt.Println(sum(extrapolatedValues))

}

func makeDiffLine(data []int) []int {
	diffLine := []int{}
	for i := 1; i < len(data); i++ {
		diffLine = append(diffLine, data[i]-data[i-1])
	}
	return diffLine
}

func atoi(s string) int {
	lol, _ := strconv.Atoi(s)
	return lol
}

func prepend[T any](data []T, item T) []T {
	return append([]T{item}, data...)
}
func instAllWhiteSpace(x string) bool { return len(strings.TrimSpace(x)) > 0 }

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

func every[T any](values []T, selector func(T) bool) bool {

	return len(values) == len(where(values, selector))

}

func some[T any](values []T, selector func(T) bool) bool {

	return len(where(values, selector)) > 0

}

func sum(numbers []int) int {
	x := 0
	for _, n := range numbers {
		x = x + n
	}
	return x
}
