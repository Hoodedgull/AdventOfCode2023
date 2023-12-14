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

type section struct {
	dirs     map[string]bool
	isInLoop string
}

func main() {
	file, err := os.Open("./input-day-10.txt")
	check(err)

	defer file.Close()

	scanner := bufio.NewScanner(file)
	startCoordinates := [2]int{0, 0}
	pipeMap := [][]section{}

	y := 0
	for scanner.Scan() {
		// read line by line
		line := scanner.Text()

		lineSections := []section{}

		for x, v := range strings.Split(line, "") {
			dirs := map[string]bool{
				"n": false,
				"s": false,
				"e": false,
				"w": false,
			}
			if v == "S" {
				startCoordinates = [2]int{x, y}
			}
			if v == "|" {
				dirs["n"] = true
				dirs["s"] = true
			}
			if v == "-" {
				dirs["e"] = true
				dirs["w"] = true
			}
			if v == "J" {
				dirs["n"] = true
				dirs["w"] = true
			}
			if v == "7" {
				dirs["w"] = true
				dirs["s"] = true
			}
			if v == "F" {
				dirs["e"] = true
				dirs["s"] = true
			}
			if v == "L" {
				dirs["n"] = true
				dirs["e"] = true
			}
			sect := section{
				dirs:     dirs,
				isInLoop: "?",
			}
			lineSections = append(lineSections, sect)
		}

		pipeMap = append(pipeMap, lineSections)
		y++

	}

	fmt.Println("made map")
	fmt.Println(startCoordinates)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	startSect, dir1, dir2 := findStartConnections(pipeMap, startCoordinates)
	pipeMap[startCoordinates[1]][startCoordinates[0]] = startSect

	fmt.Println("start")
	fmt.Println(pipeMap[startCoordinates[1]][startCoordinates[0]])

	coord1, sect1, from1 := goDirection(dir1, pipeMap, startCoordinates)
	coord2, sect2, from2 := goDirection(dir2, pipeMap, startCoordinates)
	pipeMap[coord1[1]][coord1[0]].isInLoop = "L"
	pipeMap[coord2[1]][coord2[0]].isInLoop = "L"
	dist := 1
	for coord1 != coord2 {
		coord1, sect1, from1 = goDirection(getNextDir(sect1, from1), pipeMap, coord1)
		coord2, sect2, from2 = goDirection(getNextDir(sect2, from2), pipeMap, coord2)
		pipeMap[coord1[1]][coord1[0]].isInLoop = "L"
		pipeMap[coord2[1]][coord2[0]].isInLoop = "L"
		dist++
	}

	fmt.Println("dist")
	fmt.Println(dist)

	numberInLoop := 0
	for y, line := range pipeMap {
		areWeInTheLoop := false
		hasNorthConnect := false
		hasSouthConnect := false
		for x, sect := range line {
			if sect.isInLoop == "L" {
				if sect.dirs["n"] {

					hasNorthConnect = !hasNorthConnect
				}
				if sect.dirs["s"] {

					hasSouthConnect = !hasSouthConnect
				}

				if hasNorthConnect && hasSouthConnect {

					areWeInTheLoop = !areWeInTheLoop
					hasNorthConnect = false
					hasSouthConnect = false
				}
			} else {
				if areWeInTheLoop {
					numberInLoop++
					fmt.Println(x, ":", y)
				}
			}
		}
	}

	fmt.Println("loop")
	fmt.Println(numberInLoop)

}

func getNextDir(section section, cameFrom string) string {
	if cameFrom != "n" && section.dirs["n"] {
		return "n"
	}
	if cameFrom != "e" && section.dirs["e"] {
		return "e"
	}
	if cameFrom != "w" && section.dirs["w"] {
		return "w"
	}
	if cameFrom != "s" && section.dirs["s"] {
		return "s"
	}
	return ""
}

func goDirection(dir string, pipeMap [][]section, coordinates [2]int) (newCoords [2]int, sect section, cameFrom string) {

	if dir == "w" && coordinates[0] > 0 {
		return [2]int{coordinates[0] - 1, coordinates[1]}, pipeMap[coordinates[1]][coordinates[0]-1], "e"
	}
	if dir == "e" && coordinates[0] < len(pipeMap[coordinates[1]])-1 {
		return [2]int{coordinates[0] + 1, coordinates[1]}, pipeMap[coordinates[1]][coordinates[0]+1], "w"
	}

	if dir == "n" && coordinates[1] > 0 {
		return [2]int{coordinates[0], coordinates[1] - 1}, pipeMap[coordinates[1]-1][coordinates[0]], "s"
	}

	if dir == "s" && coordinates[1] < len(pipeMap)-1 {
		return [2]int{coordinates[0], coordinates[1] + 1}, pipeMap[coordinates[1]+1][coordinates[0]], "n"
	}

	return [2]int{0, 0}, section{}, ""
}

func findStartConnections(pipeMap [][]section, startCoordinates [2]int) (section, string, string) {
	dirs := map[string]bool{}
	dir1 := ""
	dir2 := ""
	_, leftNeighbor, _ := goDirection("w", pipeMap, startCoordinates)
	if leftNeighbor.dirs["e"] {
		dirs["w"] = true
		dir1 = "w"
	}
	_, rightNeighbor, _ := goDirection("e", pipeMap, startCoordinates)
	if rightNeighbor.dirs["w"] {
		dirs["e"] = true
		if dir1 == "" {
			dir1 = "e"
		} else {
			dir2 = "e"
		}
	}

	_, topNeightbor, _ := goDirection("n", pipeMap, startCoordinates)
	if topNeightbor.dirs["s"] {
		dirs["n"] = true
		if dir1 == "" {
			dir1 = "n"
		} else {
			dir2 = "n"
		}
	}
	_, downNeightbor, _ := goDirection("s", pipeMap, startCoordinates)

	if downNeightbor.dirs["n"] {
		dirs["s"] = true
		if dir1 == "" {
			dir1 = "s"
		} else {
			dir2 = "s"
		}
	}

	sect := section{
		dirs:     dirs,
		isInLoop: "L",
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
