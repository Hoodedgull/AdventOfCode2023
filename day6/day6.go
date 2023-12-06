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
	file, err := os.Open("./input-day-06.txt")
	check(err)

	defer file.Close()

	scanner := bufio.NewScanner(file)
	times := []int{}
	dists := []int{}
	isFirst := true

	for scanner.Scan() {
		// read line by line
		line := scanner.Text()
		split := strings.Split(line, ": ")
		data := split[1]
		if isFirst {
			times = Map(where(strings.Split(data, " "), instAllWhiteSpace), atoi)
			isFirst = false
		} else {
			dists = Map(where(strings.Split(data, " "), instAllWhiteSpace), atoi)

		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	total := 1
	for i := 0; i < len(times); i++ {
		time := times[i]
		recordDist := dists[i]

		lowerLimit := findLowerLimit(0, time, time, recordDist)
		upperLimit := findUpperLimit(0, time, time, recordDist)
		betterWays := (upperLimit - lowerLimit) + 1
		fmt.Println(lowerLimit, upperLimit, betterWays)
		total = total * betterWays
	}

	fmt.Println(total)

}

func findLowerLimit(minTime int, maxTime int, totalTime int, recordDist int) int {
	newTime := (minTime + maxTime) / 2
	speed := newTime
	sailingTime := totalTime - newTime
	dist := speed * sailingTime

	if newTime == minTime || newTime == maxTime {
		return -1
	}

	if dist > recordDist {

		newLimit := findLowerLimit(minTime, newTime, totalTime, recordDist)
		if newLimit == -1 {
			return newTime
		} else {
			return newLimit
		}

	} else if dist <= recordDist {
		return findLowerLimit(newTime, maxTime, totalTime, recordDist)
	}
	return -1
}

func findUpperLimit(minTime int, maxTime int, totalTime int, recordDist int) int {
	newTime := (minTime + maxTime) / 2
	speed := newTime
	sailingTime := totalTime - newTime
	dist := speed * sailingTime

	if newTime == minTime || newTime == maxTime {
		return -1
	}

	if dist > recordDist {

		newLimit := findUpperLimit(newTime, maxTime, totalTime, recordDist)
		if newLimit == -1 {
			return newTime
		} else {
			return newLimit
		}

	} else if dist <= recordDist {
		return findUpperLimit(minTime, newTime, totalTime, recordDist)
	}
	return -1
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
