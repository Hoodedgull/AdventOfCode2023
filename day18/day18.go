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
	file, err := os.Open("./input-day-18.txt")
	check(err)

	defer file.Close()

	scanner := bufio.NewScanner(file)

	instructions := []instruction{}

	for scanner.Scan() {
		// read line by line
		line := scanner.Text()
		split := strings.Split(line, " ")
		dir := split[0]
		len := atoi(split[1])
		color := strings.Replace(split[2], "(#", "", -1)
		color = strings.Replace(color, ")", "", -1)
		inst := instruction{color: color, dir: dir, len: len}
		instructions = append(instructions, inst)

	}

	// Create a map with _real_ safe margins
	totalRight := sum(Map(where(instructions, func(in instruction) bool {
		return in.dir == "R"
	}), func(in instruction) int {
		return in.len
	}))
	totaLeft := sum(Map(where(instructions, func(in instruction) bool {
		return in.dir == "L"
	}), func(in instruction) int {
		return in.len
	}))
	totalUp := sum(Map(where(instructions, func(in instruction) bool {
		return in.dir == "U"
	}), func(in instruction) int {
		return in.len
	}))
	totalDown := sum(Map(where(instructions, func(in instruction) bool {
		return in.dir == "D"
	}), func(in instruction) int {
		return in.len
	}))

	pipeMap := [][]section{}
	for i := 0; i < totalUp+totalDown+2; i++ {
		row := make([]section, totaLeft+totalRight+2)
		for j := 0; j < totaLeft+totalRight+2; j++ {
			row[j] = section{dirs: map[string]bool{}}
		}
		pipeMap = append(pipeMap, row)
	}

	fmt.Println("Made empty map")

	x, y := totaLeft, totalUp

	for _, in := range instructions {
		for i := 0; i < in.len; i++ {
			if in.dir == "R" {
				pipeMap[y][x].dirs["e"] = true
				pipeMap[y][x].isInLoop = "#"
				x = x + 1
				pipeMap[y][x].dirs["w"] = true
				pipeMap[y][x].isInLoop = "#"
			}

			if in.dir == "L" {
				pipeMap[y][x].dirs["w"] = true
				pipeMap[y][x].isInLoop = "#"
				x = x - 1
				pipeMap[y][x].dirs["e"] = true
				pipeMap[y][x].isInLoop = "#"
			}

			if in.dir == "U" {
				pipeMap[y][x].dirs["n"] = true
				pipeMap[y][x].isInLoop = "#"
				y = y - 1
				pipeMap[y][x].dirs["s"] = true
				pipeMap[y][x].isInLoop = "#"
			}

			if in.dir == "D" {
				pipeMap[y][x].dirs["s"] = true
				pipeMap[y][x].isInLoop = "#"
				y = y + 1
				pipeMap[y][x].dirs["n"] = true
				pipeMap[y][x].isInLoop = "#"
			}

		}
	}

	fmt.Println("Filled Out Map")
	visual := strings.Join(Map(pipeMap, func(row []section) string {
		return strings.Join(Map(row, func(sec section) string {
			if sec.isInLoop == "#" {
				return "#"
			} else {
				return "."
			}
		}), "")
	}), "\n")
	fmt.Println(visual)

	numberOfTrenchTiles := strings.Count(visual, "#")

	numberOfInteriorTiles := 0
	for _, line := range pipeMap {
		areWeInTheLoop := false
		hasNorthConnect := false
		hasSouthConnect := false
		for _, sect := range line {
			if sect.isInLoop == "#" {
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
