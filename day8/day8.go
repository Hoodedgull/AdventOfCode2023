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
	startingNodes := []string{}

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
			if strings.HasSuffix(node, "A") {
				startingNodes = append(startingNodes, node)
			}
		}
		index++

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("made map")

	nodes := startingNodes
	lastTimeMap := map[int]int{}
	steps := 0
	index = 0
	done := false
	for !done {
		steps++
		nextNodes := []string{}
		for _, n := range nodes {

			children := nodeMap[n]
			instruction := instructions[index]
			if instruction == "L" {
				nextNodes = append(nextNodes, children[0])
			} else {
				nextNodes = append(nextNodes, children[1])
			}
		}
		nodes = nextNodes
		stepsSinceLastTime := []int{}
		anyHasZ := false
		for i, n := range nodes {

			if strings.HasSuffix(n, "Z") {
				lastTime := lastTimeMap[i]
				stepsSinceLastTime = append(stepsSinceLastTime, steps-lastTime)
				lastTimeMap[i] = steps
				anyHasZ = true
			} else {

				stepsSinceLastTime = append(stepsSinceLastTime, -1)
			}
		}

		if anyHasZ {

			fmt.Println(stepsSinceLastTime)
		}
		if every(nodes, func(n string) bool { return strings.HasSuffix(n, "Z") }) {
			done = true
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
