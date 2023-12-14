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
	cardCopies := map[int]int{}

	for scanner.Scan() {
		// read line by line
		line := scanner.Text()
		split := strings.Split(line, ": ")
		cardId, _ := strconv.Atoi(strings.TrimSpace(strings.Split(split[0], "Card")[1]))
		data := split[1]
		winningNums := Map(where(strings.Split(strings.Split(data, "|")[0], " "), instAllWhiteSpace), atoi)
		haveNums := Map(where(strings.Split(strings.Split(data, "|")[1], " "), instAllWhiteSpace), atoi)

		numOfCards := cardCopies[cardId]
		numOfCards = numOfCards + 1 // Add original to number of copies
		numbers = append(numbers, numOfCards)

		// Make card copies for winning numbers
		winningNumsWeHave := where(haveNums, func(num int) bool { return includes(winningNums, func(i int) bool { return i == num }) })

		for i := 1; i <= len(winningNumsWeHave); i++ {
			wonCardId := cardId + i
			cardCopies[wonCardId] = cardCopies[wonCardId] + numOfCards
		}

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
