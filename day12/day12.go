package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
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
		checksumsa := strings.Split(split[1], ",")
		checksums := append(checksumsa, checksumsa...)
		checksums = append(checksums, checksumsa...)
		checksums = append(checksums, checksumsa...)
		checksums = append(checksums, checksumsa...)

		// runes := []rune(springs)
		runes := []rune(springs + "?" + springs + "?" + springs + "?" + springs + "?" + springs)
		lenth := len(runes)
		intialCombination := []rune(strings.Repeat(".", lenth))
		combinations := [][]rune{intialCombination}

		// Instead of making all the combinations, try just making all valid combinations
		for i, c := range checksums {
			newCombs := [][]rune{}
			for _, comb := range combinations {

				newCombs = append(newCombs, getNextCombinations(c, comb, checksums[i+1:], runes)...)
			}
			combinations = newCombs
		}

		fmt.Println("made combs for " + line)
		fmt.Println(len(combinations))

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

func getNextCombinations(c string, runes []rune, followingChecksums []string, original []rune) [][]rune {

	check := atoi(c)
	lengthOfAfter := sum(Map(followingChecksums, atoi)) + len(followingChecksums) - 1 // The lenghts of checksums, plus one separator for each of them
	start := 0
	lastIdx := lastIndex(runes, '#')
	if lastIdx != -1 {
		start = lastIdx + 2
	}

	end := len(runes) - lengthOfAfter

	result := [][]rune{}
	for i := start; i < end-check; i++ {
		runeCopy1 := make([]rune, len(runes))
		copy(runeCopy1, runes)
		for j := 0; j < check; j++ {
			runeCopy1[i+j] = '#'
		}

		if !violates(runeCopy1, original, i+check) {
			result = append(result, runeCopy1)
		}
	}

	return result
}

func violates(newRunes []rune, originalRunes []rune, checkToIdx int) bool {
	if len(newRunes) != len(originalRunes) || checkToIdx > len(newRunes) {
		fmt.Println("PANIC")
	}

	for i := 0; i < checkToIdx; i++ {
		if originalRunes[i] == '?' {
			continue
		}

		if originalRunes[i] != newRunes[i] {
			return true
		}
	}
	return false
}

func lastIndex[T comparable](runes []T, item T) int {
	runeCopy1 := make([]T, len(runes))
	copy(runeCopy1, runes)
	// reverse
	for i, j := 0, len(runeCopy1)-1; i < j; i, j = i+1, j-1 {
		runeCopy1[i], runeCopy1[j] = runeCopy1[j], runeCopy1[i]
	}
	index := slices.Index(runeCopy1, item)
	if index == -1 {
		return index
	}
	lastindex := len(runeCopy1) - 1 - index
	return lastindex
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
