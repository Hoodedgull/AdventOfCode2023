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

type Beam struct {
	x   int
	y   int
	dir string
}

func main() {
	file, err := os.Open("./input-day-16.txt")
	check(err)

	defer file.Close()

	scanner := bufio.NewScanner(file)

	grid := [][]string{}
	energized := [][]string{}
	for scanner.Scan() {
		// read line by line
		line := scanner.Text()
		grid = append(grid, strings.Split(line, ""))
		energized = append(energized, make([]string, len(line)))
	}

	beams := []Beam{Beam{x: -1, y: 0, dir: "e"}}

	for len(beams) > 0 {
		newBeams := []Beam{}
		for _, beam := range beams {

			x, y := getNextPos(beam)
			if x < 0 || x >= len(grid[0]) {
				continue
			}
			if y < 0 || y >= len(grid) {
				continue
			}

			item := grid[y][x]
			energizedItem := energized[y][x]
			if strings.Contains(energizedItem, beam.dir) {
				continue
			}
			energized[y][x] = energizedItem + beam.dir

			if item == "." {
				newBeams = append(newBeams, Beam{x, y, beam.dir})
			} else if item == "/" {
				if beam.dir == "e" {
					newBeams = append(newBeams, Beam{x, y, "n"})
				}
				if beam.dir == "w" {
					newBeams = append(newBeams, Beam{x, y, "s"})
				}
				if beam.dir == "n" {
					newBeams = append(newBeams, Beam{x, y, "e"})
				}
				if beam.dir == "s" {
					newBeams = append(newBeams, Beam{x, y, "w"})
				}
			} else if item == "\\" {
				if beam.dir == "e" {
					newBeams = append(newBeams, Beam{x, y, "s"})
				}
				if beam.dir == "w" {
					newBeams = append(newBeams, Beam{x, y, "n"})
				}
				if beam.dir == "n" {
					newBeams = append(newBeams, Beam{x, y, "w"})
				}
				if beam.dir == "s" {
					newBeams = append(newBeams, Beam{x, y, "e"})
				}
			} else if item == "-" {
				if beam.dir == "e" {
					newBeams = append(newBeams, Beam{x, y, "e"})
				}
				if beam.dir == "w" {
					newBeams = append(newBeams, Beam{x, y, "w"})
				}
				if beam.dir == "n" {
					newBeams = append(newBeams, Beam{x, y, "e"})
					newBeams = append(newBeams, Beam{x, y, "w"})
				}
				if beam.dir == "s" {
					newBeams = append(newBeams, Beam{x, y, "e"})
					newBeams = append(newBeams, Beam{x, y, "w"})
				}
			} else if item == "|" {
				if beam.dir == "e" {
					newBeams = append(newBeams, Beam{x, y, "n"})
					newBeams = append(newBeams, Beam{x, y, "s"})
				}
				if beam.dir == "w" {
					newBeams = append(newBeams, Beam{x, y, "n"})
					newBeams = append(newBeams, Beam{x, y, "s"})
				}
				if beam.dir == "n" {
					newBeams = append(newBeams, Beam{x, y, "n"})
				}
				if beam.dir == "s" {
					newBeams = append(newBeams, Beam{x, y, "s"})
				}
			}
		}
		beams = newBeams
	}

	fmt.Println("Got to an exit?!")

	total := sum(Map(energized, func(row []string) int {
		return len(where(row, instAllWhiteSpace))
	}))
	fmt.Println(total)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}

func getNextPos(beam Beam) (int, int) {
	if beam.dir == "e" {
		return beam.x + 1, beam.y
	}
	if beam.dir == "w" {
		return beam.x - 1, beam.y
	}
	if beam.dir == "n" {
		return beam.x, beam.y - 1
	}
	if beam.dir == "s" {
		return beam.x, beam.y + 1
	}
	fmt.Println("PANIC")
	return 0, 0
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
