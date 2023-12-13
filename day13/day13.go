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
	file, err := os.Open("./input-day-13.txt")
	check(err)

	defer file.Close()

	scanner := bufio.NewScanner(file)

	patterns := [][][]string{}
	currentPattern := [][]string{}
	for scanner.Scan() {
		// read line by line
		line := scanner.Text()

		if !instAllWhiteSpace(line) {
			patterns = append(patterns, currentPattern)
			currentPattern = [][]string{}
		} else {

			split := strings.Split(line, "")
			currentPattern = append(currentPattern, split)
		}
	}

	fmt.Println("Found patterns x", len(patterns))

	total := 0
	for _, pattern := range patterns {
		horizontalMatchIndex := -1
		// find horizontal mirroring
		for i := 1; i < len(pattern); i++ {

			line1 := i - 1
			line2 := i
			foundMisMatch := false

			for line1 >= 0 && line2 < len(pattern) {
				if !isRowsSame(pattern, line1, line2) {
					foundMisMatch = true
					break
				}
				line1--
				line2++
			}
			if !foundMisMatch {
				// We found 2 lines that seemingly match!
				// Verify that the other lines match as well
				horizontalMatchIndex = i
				total = total + horizontalMatchIndex*100

			}
		}

		if horizontalMatchIndex == -1 {

			// find vertical mirroring
			for i := 1; i < len(pattern[0]); i++ {

				line1 := i - 1
				line2 := i
				foundMisMatch := false

				for line1 >= 0 && line2 < len(pattern[0]) {
					if !isColsSame(pattern, line1, line2) {
						foundMisMatch = true
						break
					}
					line1--
					line2++
				}
				if !foundMisMatch {
					// We found 2 lines that seemingly match!
					// TODO: Verify that the other lines match as well

					total = total + i
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("total")
	fmt.Println(total)

}

func isRowsSame(pattern [][]string, rowIndex1 int, rowIndex2 int) bool {
	row1 := pattern[rowIndex1]
	row2 := pattern[rowIndex2]
	if len(row1) != len(row2) {
		return false
	}

	for i := 0; i < len(row1); i++ {
		item1 := row1[i]
		item2 := row2[i]
		if item1 != item2 {
			return false
		}
	}
	return true
}

func isColsSame(pattern [][]string, colIndex1 int, colIndex2 int) bool {
	for i := 0; i < len(pattern); i++ {
		item1 := pattern[i][colIndex1]
		item2 := pattern[i][colIndex2]
		if item1 != item2 {
			return false
		}
	}
	return true
}

func atoi(s string) int {
	lol, _ := strconv.Atoi(s)
	return lol
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
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
