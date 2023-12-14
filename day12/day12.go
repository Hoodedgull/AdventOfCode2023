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
	file, err := os.Open("./sample-day-12.txt")
	check(err)

	defer file.Close()

	scanner := bufio.NewScanner(file)

	taskCounter := 0
	messages := make(chan int)

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
		// runes := []rune(springs + "?" + springs + "?" + springs + "?" + springs)
		runes := []rune(springs + "?" + springs + "?" + springs + "?" + springs + "?" + springs)
		lenth := len(runes)
		intialCombination := []rune(strings.Repeat(".", lenth))

		taskCounter++
		go getNextCombinations(checksums[0], intialCombination, checksums[0+1:], runes, messages, 0)

		fmt.Println("startedTask")
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	total := 0
	for i := 0; i < taskCounter; i++ {
		sub := <-messages
		fmt.Println("got sub", sub)

		total += sub
	}

	fmt.Println("total")
	fmt.Println(total)

}

func getNextCombinations(c string, runes []rune, followingChecksums []string, original []rune, out chan<- int, recLvl int) int {

	check := atoi(c)
	lengthOfAfter := sum(Map(followingChecksums, atoi)) + len(followingChecksums) - 1 // The lenghts of checksums, plus one separator for each of them
	start := 0
	lastIdx := lastIndex(runes, '#')
	if lastIdx != -1 {
		start = lastIdx + 2
	}

	end := len(runes) - lengthOfAfter

	innerTaskCounter := 0
	in := make(chan int)
	result := 0
	for i := start; i < end-check; i++ {
		runeCopy1 := make([]rune, len(runes))
		copy(runeCopy1, runes)
		for j := 0; j < check; j++ {
			runeCopy1[i+j] = '#'
		}

		if !violates(runeCopy1, original, i+check) {
			if len(followingChecksums) > 0 {
				innerTaskCounter++
				if recLvl < 3 {

					go getNextCombinations(followingChecksums[0], runeCopy1, followingChecksums[0+1:], original, in, recLvl+1)
				} else {
					result += getNextCombinations(followingChecksums[0], runeCopy1, followingChecksums[0+1:], original, in, recLvl+1)
				}
			} else {
				result += 1
			}
		}
	}

	if recLvl < 3 {

		for i := 0; i < innerTaskCounter; i++ {
			subsub := <-in
			result += subsub
		}
	}

	fmt.Println("Done", recLvl)
	if recLvl <= 3 {

		out <- result
		return 0
	} else {
		return result
	}
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
