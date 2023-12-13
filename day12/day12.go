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
	file, err := os.Open("./input-day-12.txt")
	check(err)

	defer file.Close()

	scanner := bufio.NewScanner(file)
	total := 0

	for scanner.Scan() {
		// read line by line
		line := scanner.Text()

		split := strings.Split(line, " ")
		springs := split[0]
		checksums := strings.Split(split[1], ",")

		runes := []rune(springs)
		combinations := [][]rune{runes}
		lenth := len(runes)

		for i := 0; i < lenth; i++ {
			newCombs := [][]rune{}
			for _, comb := range combinations {

				newCombs = append(newCombs, getNextCombinations(i, comb)...)
			}
			combinations = newCombs
		}

		fmt.Println("made combs for " + line)
		// fmt.Println(Map(combinations, func(slc []rune) string { return string(slc) }))

		subtotal := 0
		for _, comp := range combinations {

			if isValidComp(comp, checksums) {
				subtotal++
			}
		}

		fmt.Println("Theres this amny valid arrangements for " + line)
		fmt.Println(subtotal)

		total = total + subtotal
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("total")
	fmt.Println(total)

}

func isValidComp(comp []rune, checksums []string) bool {

	damagedSpringGroups := where(strings.Split(string(comp), "."), instAllWhiteSpace)
	if len(damagedSpringGroups) != len(checksums) {
		return false
	}

	for i, v := range checksums {
		springGroup := damagedSpringGroups[i]
		expectedLen := atoi(v)
		if len(springGroup) != expectedLen {
			return false
		}
	}
	return true
}

func getNextCombinations(i int, runes []rune) [][]rune {
	if runes[i] == '?' {
		runeCopy1 := make([]rune, len(runes))
		copy(runeCopy1, runes)
		runeCopy2 := make([]rune, len(runes))
		copy(runeCopy2, runes)
		runeCopy1[i] = '.'
		runeCopy2[i] = '#'
		return [][]rune{runeCopy1, runeCopy2}
	}

	return [][]rune{runes}
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
