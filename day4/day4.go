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
	file, err := os.Open("./input-day04.txt")
	check(err)

	defer file.Close()

	scanner := bufio.NewScanner(file)
	numbers := []int{}

	for scanner.Scan() {
		// read line by line
		line := scanner.Text()
		data := strings.Split(line, ": ")[1]
		winningNums := Map(where(strings.Split(strings.Split(data, "|")[0], " "), instAllWhiteSpace), atoi)
		haveNums := Map(where(strings.Split(strings.Split(data, "|")[1], " "), instAllWhiteSpace), atoi)

		cardValue := 0
		for _, num := range haveNums {
			if includes(winningNums, func(i int) bool { return i == num }) {
				if cardValue == 0 {
					cardValue = 1
				} else {
					cardValue = cardValue * 2
				}
			}
		}

		numbers = append(numbers, cardValue)

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(numbers)
	fmt.Println(sum(numbers))

}

func atoi(s string) int {
	lol, _ := strconv.Atoi(s)
	return lol
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

func sum(numbers []int) int {
	x := 0
	for _, n := range numbers {
		x = x + n
	}
	return x
}
