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
	file, err := os.Open("./input-day-10.txt")
	check(err)

	defer file.Close()

	scanner := bufio.NewScanner(file)
	startCoordinates := [2]int{0, 0}
	pipeMap := [][]map[string]bool{}

	y := 0
	for scanner.Scan() {
		// read line by line
		line := scanner.Text()

		lineSections := []map[string]bool{}

		for x, v := range strings.Split(line, "") {
			sect := map[string]bool{
				"n": false,
				"s": false,
				"e": false,
				"w": false,
			}
			if v == "S" {
				startCoordinates = [2]int{x, y}
			}
			if v == "|" {
				sect["n"] = true
				sect["s"] = true
			}
			if v == "-" {
				sect["e"] = true
				sect["w"] = true
			}
			if v == "J" {
				sect["n"] = true
				sect["w"] = true
			}
			if v == "7" {
				sect["w"] = true
				sect["s"] = true
			}
			if v == "F" {
				sect["e"] = true
				sect["s"] = true
			}
			if v == "L" {
				sect["n"] = true
				sect["e"] = true
			}
			lineSections = append(lineSections, sect)
		}

		pipeMap = append(pipeMap, lineSections)
		y++

	}

	fmt.Println("made map")

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	startSect, dir1, dir2 := findStartConnections(pipeMap, startCoordinates)
	pipeMap[startCoordinates[0]][startCoordinates[1]] = startSect

	fmt.Println("start")
	fmt.Println(pipeMap[startCoordinates[0]][startCoordinates[1]])

	coord1, sect1, from1 := goDirection(dir1, pipeMap, startCoordinates)
	coord2, sect2, from2 := goDirection(dir2, pipeMap, startCoordinates)
	dist := 1
	for coord1 != coord2 {
		coord1, sect1, from1 = goDirection(getNextDir(sect1, from1), pipeMap, coord1)
		coord2, sect2, from2 = goDirection(getNextDir(sect2, from2), pipeMap, coord2)
		dist++
	}

	fmt.Println("dist")
	fmt.Println(dist)

}

func getNextDir(section map[string]bool, cameFrom string) string {
	if cameFrom != "n" && section["n"] {
		return "n"
	}
	if cameFrom != "e" && section["e"] {
		return "e"
	}
	if cameFrom != "w" && section["w"] {
		return "w"
	}
	if cameFrom != "s" && section["s"] {
		return "s"
	}
	return ""
}

func goDirection(dir string, pipeMap [][]map[string]bool, coordinates [2]int) (newCoords [2]int, section map[string]bool, cameFrom string) {

	if dir == "w" && coordinates[0] > 0 {
		return [2]int{coordinates[0] - 1, coordinates[1]}, pipeMap[coordinates[1]][coordinates[0]-1], "e"
	}
	if dir == "e" && coordinates[0] < len(pipeMap)-1 {
		return [2]int{coordinates[0] + 1, coordinates[1]}, pipeMap[coordinates[1]][coordinates[0]+1], "w"
	}

	if dir == "n" && coordinates[1] > 0 {
		return [2]int{coordinates[0], coordinates[1] - 1}, pipeMap[coordinates[1]-1][coordinates[0]], "s"
	}

	if dir == "s" && coordinates[1] < len(pipeMap[coordinates[0]])-1 {
		return [2]int{coordinates[0], coordinates[1] + 1}, pipeMap[coordinates[1]+1][coordinates[0]], "n"
	}

	return [2]int{0, 0}, map[string]bool{}, ""
}

func findStartConnections(pipeMap [][]map[string]bool, startCoordinates [2]int) (map[string]bool, string, string) {
	sect := map[string]bool{}
	dir1 := ""
	dir2 := ""
	_, leftNeighbor, _ := goDirection("w", pipeMap, startCoordinates)
	if leftNeighbor["e"] {
		sect["w"] = true
		dir1 = "w"
	}
	_, rightNeighbor, _ := goDirection("e", pipeMap, startCoordinates)
	if rightNeighbor["w"] {
		sect["e"] = true
		if dir1 == "" {
			dir1 = "e"
		} else {
			dir2 = "e"
		}
	}

	_, topNeightbor, _ := goDirection("n", pipeMap, startCoordinates)
	if topNeightbor["s"] {
		sect["n"] = true
		if dir1 == "" {
			dir1 = "n"
		} else {
			dir2 = "n"
		}
	}
	_, downNeightbor, _ := goDirection("s", pipeMap, startCoordinates)

	if downNeightbor["n"] {
		sect["s"] = true
		if dir1 == "" {
			dir1 = "s"
		} else {
			dir2 = "s"
		}
	}

	return sect, dir1, dir2
}

func atoi(s string) int {
	lol, _ := strconv.Atoi(s)
	return lol
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
