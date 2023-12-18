package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type instruction struct {
	dir   string
	len   int
	color string
}

type section struct {
	dirs     map[string]bool
	isInLoop string
}

func main() {
	file, err := os.Open("./sample-day-18.txt")
	check(err)

	defer file.Close()

	scanner := bufio.NewScanner(file)

	instructions := []instruction{}

	for scanner.Scan() {
		// read line by line
		line := scanner.Text()
		split := strings.Split(line, " ")
		dir := split[0]
		color := strings.Replace(split[2], "(#", "", -1)
		color = strings.Replace(color, ")", "", -1)
		len := parseHex(color[0:5])
		// len = atoi(split[1]) // For testing part1
		runes := []rune(color)
		if runes[5] == '0' {
			dir = "R"
		}
		if runes[5] == '1' {
			dir = "D"
		}
		if runes[5] == '2' {
			dir = "L"
		}
		if runes[5] == '3' {
			dir = "U"
		}
		// dir = split[0] // For testing part 1
		inst := instruction{color: color, dir: dir, len: len}
		instructions = append(instructions, inst)

	}

	pipeMap := map[int]map[int]section{}

	x, y := 0, 0
	numberOfTrenchTiles := 0

	for _, in := range instructions {
		for i := 0; i < in.len; {
			current := getSafe(pipeMap, x, y)
			if in.dir == "R" {
				current.dirs["e"] = true
				current.isInLoop = "#"
				pipeMap[y][x] = current
				x = x + in.len
				i += in.len
				numberOfTrenchTiles += in.len
				next := getSafe(pipeMap, x, y)
				next.dirs["w"] = true
				next.isInLoop = "#"
				pipeMap[y][x] = next
			}

			if in.dir == "L" {
				current.dirs["w"] = true
				current.isInLoop = "#"
				pipeMap[y][x] = current
				x = x - in.len
				i += in.len
				numberOfTrenchTiles += in.len
				next := getSafe(pipeMap, x, y)
				next.dirs["e"] = true
				next.isInLoop = "#"
				pipeMap[y][x] = next
			}

			if in.dir == "U" {
				current.dirs["n"] = true
				current.isInLoop = "#"
				pipeMap[y][x] = current
				y = y - 1
				i += 1
				numberOfTrenchTiles += 1
				next := getSafe(pipeMap, x, y)
				next.dirs["s"] = true
				next.isInLoop = "#"
				pipeMap[y][x] = next
			}

			if in.dir == "D" {
				current.dirs["s"] = true
				current.isInLoop = "#"
				pipeMap[y][x] = current
				y = y + 1
				i += 1
				numberOfTrenchTiles += 1
				next := getSafe(pipeMap, x, y)
				next.dirs["n"] = true
				next.isInLoop = "#"
				pipeMap[y][x] = next
			}

		}
	}

	fmt.Println("Filled Out Map")

	fmt.Println(numberOfTrenchTiles)

	yValues := make([]int, 0)
	for k, _ := range pipeMap {
		yValues = append(yValues, k)
	}
	sort.Ints(yValues)

	numberOfInteriorTiles := 0
	for _, y := range yValues {

		xValues := make([]int, 0)
		for k, _ := range pipeMap[y] {
			xValues = append(xValues, k)
		}
		sort.Ints(xValues)

		// if len(xValues) > 2 {
		// 	fmt.Println("exiting!")
		// }
		areWeInTheLoop := false
		hasNorthConnect := false
		hasSouthConnect := false
		xEntry := 0
		xExit := 0
		for _, x := range xValues {
			sect := pipeMap[y][x]

			if sect.isInLoop == "#" {
				if sect.dirs["n"] {
					if !hasNorthConnect && !hasSouthConnect {
						xExit = x
					}
					if hasNorthConnect && areWeInTheLoop {
						// we have been in a sideways trench, which is counted later, so we need to subtract it
						numberOfInteriorTiles = numberOfInteriorTiles - ((x - xExit) + 1)
					}
					hasNorthConnect = !hasNorthConnect
				}
				if sect.dirs["s"] {
					if !hasNorthConnect && !hasSouthConnect {
						xExit = x
					}
					if hasSouthConnect && areWeInTheLoop {
						// we have been in a sideways trench, which is counted later, so we need to subtract it
						numberOfInteriorTiles = numberOfInteriorTiles - ((x - xExit) + 1)
					}
					hasSouthConnect = !hasSouthConnect
				}

				if hasNorthConnect && hasSouthConnect {

					// if we are in the loop and breaking out, add all tiles since entry
					if areWeInTheLoop {
						numberOfInteriorTiles = numberOfInteriorTiles + (xExit - xEntry) - 1
					} else {
						xEntry = x
					}
					areWeInTheLoop = !areWeInTheLoop
					hasNorthConnect = false
					hasSouthConnect = false
				}
			} else {
				if areWeInTheLoop {
					numberOfInteriorTiles++
					//fmt.Println(x, ":", y)
				}
			}
		}
	}

	fmt.Println("loop")
	fmt.Println(numberOfInteriorTiles)

	fmt.Println("total:", numberOfInteriorTiles+numberOfTrenchTiles)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}

func getSafe(grid map[int]map[int]section, x int, y int) section {
	row, ok := grid[y]
	if !ok {
		row = map[int]section{}
		grid[y] = row
	}
	sect, ok2 := row[x]
	if !ok2 {
		sect = section{dirs: map[string]bool{}}
		row[x] = sect
	}
	return sect
}

func parseHex(input string) int {
	n, err := strconv.ParseInt(input, 16, 64)
	if err != nil {
		panic(err)
	}
	return int(n)
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
