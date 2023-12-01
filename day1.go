package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	file, err := os.Open("./input-day-01.txt")
	check(err)

	defer file.Close()

	scanner := bufio.NewScanner(file)

	numbers := []int{}
	for scanner.Scan() {
		// read line by line
		line := scanner.Text()

		firstNum, err := FirstDigit(line)
		check(err)
		lastNum, err := LastDigit(line)
		check(err)
		var lineNumStr string = string(firstNum) + string(lastNum)
		lineNum, err := strconv.Atoi(lineNumStr)
		check(err)

		numbers = append(numbers, lineNum)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(sum(numbers))

}
func FirstDigit(s string) (rune, error) {
	for _, r := range s {
		if unicode.IsDigit(r) {
			return r, nil
		}
	}
	return 0, errors.New("no number!")
}
func LastDigit(s string) (rune, error) {
	runes := []rune(s)
	for i := len(runes) - 1; i >= 0; i-- {
		r := runes[i]
		if unicode.IsDigit(r) {
			return r, nil
		}
	}
	return 0, errors.New("no number!")
}

func sum(numbers []int) int {
	x := 0
	for _, n := range numbers {
		x = x + n
	}
	return x
}
