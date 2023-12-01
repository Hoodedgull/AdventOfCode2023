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
	file, err := os.Open("./input-day-01.txt")
	check(err)

	defer file.Close()

	scanner := bufio.NewScanner(file)

	numbers := []int{}
	for scanner.Scan() {
		// read line by line
		line := scanner.Text()

		firstNum := FirstDigit(line)
		check(err)
		lastNum := LastDigit(line)
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
func FirstDigit(s string) string {
	validDigits := map[string]string{"1": "1", "2": "2", "3": "3", "4": "4", "5": "5", "6": "6", "7": "7", "8": "8", "9": "9", "one": "1", "two": "2", "three": "3", "four": "4", "five": "5", "six": "6", "seven": "7", "eight": "8", "nine": "9"}

	lowestIndex := 9999999
	valueWithLowestIndex := ""
	for key, value := range validDigits {
		index := strings.Index(s, key)
		if index > -1 && index < lowestIndex {
			lowestIndex = index
			valueWithLowestIndex = value
		}
	}

	return valueWithLowestIndex
}
func LastDigit(s string) string {
	validDigits := map[string]string{"1": "1", "2": "2", "3": "3", "4": "4", "5": "5", "6": "6", "7": "7", "8": "8", "9": "9", "one": "1", "two": "2", "three": "3", "four": "4", "five": "5", "six": "6", "seven": "7", "eight": "8", "nine": "9"}

	highest := -1
	bestValue := ""
	for key, value := range validDigits {
		index := strings.LastIndex(s, key)
		if index > highest {
			highest = index
			bestValue = value
		}
	}

	return bestValue
}

func sum(numbers []int) int {
	x := 0
	for _, n := range numbers {
		x = x + n
	}
	return x
}
