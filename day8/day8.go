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
	file, err := os.Open("./input-day-08.txt")
	check(err)

	defer file.Close()

	scanner := bufio.NewScanner(file)
	nodeMap := map[string][]string{}
	instructions := []string{}

	index := 0

	for scanner.Scan() {
		// read line by line
		line := scanner.Text()

		if index == 0 {
			instructions = strings.Split(line, "")
		}
		if index > 1 {
			split := strings.Split(line, " = ")
			node := split[0]
			data := split[1]
			children := strings.Split(strings.Replace(strings.Replace(data, "(", "", -1), ")", "", -1), ", ")
			nodeMap[node] = children

		}
		index++

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("made map")

	node := "AAA"
	steps := 0
	index = 0
	for node != "ZZZ" {
		steps++
		nextNodes := nodeMap[node]
		instruction := instructions[index]
		if instruction == "L" {
			node = nextNodes[0]
		} else {
			node = nextNodes[1]
		}
		index++
		if index >= len(instructions) {
			index = 0
		}
	}

	fmt.Println("Steps")
	fmt.Println(steps)

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
