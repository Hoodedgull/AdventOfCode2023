package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Beam struct {
	x    int
	y    int
	dir  string
	cost int
	path string
}

func main() {
	file, err := os.Open("./input-day-17.txt")
	check(err)

	defer file.Close()

	scanner := bufio.NewScanner(file)

	grid := [][]int{}

	for scanner.Scan() {
		// read line by line
		line := scanner.Text()
		grid = append(grid, Map(strings.Split(line, ""), atoi))

	}

	costGrid := [][]map[string]map[int]int{}
	for _, row := range grid {
		costGrid = append(costGrid, make([]map[string]map[int]int, len(row)))
	}

	start := Beam{x: 0, y: 0, dir: "e", cost: 0, path: ""}
	goalx := len(grid[0]) - 1
	goaly := len(grid) - 1
	beams := []Beam{start}
	bestBeam := Beam{}
	currentBestCost := 999999999
	for len(beams) > 0 {
		beam := beams[0]      // take first item out to process it
		newBeams := beams[1:] // remove first item

		options := getOptions(beam, grid)
		for _, opt := range options {
			if opt.x == goalx && opt.y == goaly {
				fmt.Println("Found one way with cost: ", opt.cost)
				if opt.cost < currentBestCost {
					currentBestCost = opt.cost
					bestBeam = opt
					newBeams = []Beam{}
					break
				}
			}
			currentBestCost := costGrid[opt.y][opt.x]
			if currentBestCost == nil {
				currentBestCost = map[string]map[int]int{}
				costGrid[opt.y][opt.x] = currentBestCost
			}
			bestCostForDir := currentBestCost[opt.dir]
			if bestCostForDir == nil {
				bestCostForDir = make(map[int]int)
				currentBestCost[opt.dir] = bestCostForDir
			}

			stepSize := countSuffix(beam, opt.dir)
			foundAGoodCase := false
			foundABetterCase := false
			for step := 0; step <= 10; step++ {
				bestCostForStepSize := bestCostForDir[step]
				if bestCostForStepSize == 0 || opt.cost < bestCostForStepSize {
					if step >= stepSize {
						bestCostForDir[stepSize] = opt.cost
					}
					if step == stepSize {
						foundAGoodCase = true
					}
					if step < stepSize && bestCostForStepSize != 0 && bestCostForStepSize < opt.cost {
						foundABetterCase = true
					}
				}
			}

			if foundAGoodCase && !foundABetterCase {
				//costGrid[opt.y][opt.x] = opt.cost
				newBeams = append(newBeams, opt)
			}
		}

		slices.SortFunc(newBeams, func(a Beam, b Beam) int {

			// adist := Abs(a.x-goalx) + Abs(a.y-goaly)
			// bdist := Abs(b.x-goalx) + Abs(b.y-goaly)
			// return (adist + a.cost) - (bdist + b.cost)
			return a.cost - b.cost
		})
		beams = newBeams

	}

	fmt.Println("Got to an exit?!")
	fmt.Println(currentBestCost)
	fmt.Println(bestBeam) // Found 947-> too high

	// fmt.Println("Best cost", costGrid[goaly][goalx])

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}

func countSuffix(beam Beam, dir string) int {
	s := []rune(beam.path)
	last3 := string(s[max(len(s)-10, 0):])

	stepSize := 0
	for step := 10; step > 0; step-- {
		if strings.HasSuffix(last3, strings.Repeat(dir, step)) {
			stepSize = step
			break
		}
	}
	return stepSize
}

func getOptions(beam Beam, grid [][]int) []Beam {
	beams := []Beam{}
	esInEnd := countSuffix(beam, "e")
	wsInEnd := countSuffix(beam, "w")
	nsInEnd := countSuffix(beam, "n")
	ssInEnd := countSuffix(beam, "s")
	isForcedNorth := nsInEnd > 0 && nsInEnd < 4
	isForcedEast := esInEnd > 0 && esInEnd < 4
	isForcedSouth := ssInEnd > 0 && ssInEnd < 4
	isForcedWest := wsInEnd > 0 && wsInEnd < 4
	if beam.dir != "w" && esInEnd < 10 && !isForcedNorth && !isForcedSouth && !isForcedWest {
		x := beam.x + 1
		y := beam.y

		if isWithinGrid(x, y, grid) && canStopInGoal(x, y, grid, esInEnd+1, 0) {
			newCost := grid[y][x]
			beams = append(beams, Beam{x: x, y: y, dir: "e", cost: beam.cost + newCost, path: beam.path + "e"})
		}
	}
	if beam.dir != "e" && wsInEnd < 10 && !isForcedNorth && !isForcedSouth && !isForcedEast {
		x := beam.x - 1
		y := beam.y

		if isWithinGrid(x, y, grid) && canStopInGoal(x, y, grid, esInEnd, ssInEnd) {
			newCost := grid[y][x]
			beams = append(beams, Beam{x: x, y: y, dir: "w", cost: beam.cost + newCost, path: beam.path + "w"})
		}
	}
	if beam.dir != "s" && nsInEnd < 10 && !isForcedEast && !isForcedSouth && !isForcedWest {
		x := beam.x
		y := beam.y - 1

		if isWithinGrid(x, y, grid) && canStopInGoal(x, y, grid, 0, ssInEnd+1) {
			newCost := grid[y][x]
			beams = append(beams, Beam{x: x, y: y, dir: "n", cost: beam.cost + newCost, path: beam.path + "n"})
		}
	}
	if beam.dir != "n" && ssInEnd < 10 && !isForcedNorth && !isForcedEast && !isForcedWest {
		x := beam.x
		y := beam.y + 1

		if isWithinGrid(x, y, grid) && canStopInGoal(x, y, grid, esInEnd, ssInEnd) {
			newCost := grid[y][x]
			beams = append(beams, Beam{x: x, y: y, dir: "s", cost: beam.cost + newCost, path: beam.path + "s"})
		}
	}
	return beams
}

func canStopInGoal(x int, y int, grid [][]int, esInEnd int, ssInEnd int) bool {
	if y < len(grid)-1 || x < len(grid[0])-1 {
		return true
	}

	return esInEnd >= 4 || ssInEnd >= 4
}

func isWithinGrid(x int, y int, grid [][]int) bool {
	if x < 0 || x >= len(grid[0]) {
		return false
	}
	if y < 0 || y >= len(grid) {
		return false
	}
	return true
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
