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
	file, err := os.Open("./input-day-11.txt")
	check(err)

	defer file.Close()

	scanner := bufio.NewScanner(file)
	galaxyCoordinates := [][2]int{} // x, y
	lengthOfSomeRow := 0

	y := 0
	for scanner.Scan() {
		// read line by line
		line := scanner.Text()
		rowHasGalaxies := false
		for x, v := range strings.Split(line, "") {

			if v == "#" {
				galaxyCoordinates = append(galaxyCoordinates, [2]int{x, y})
				rowHasGalaxies = true
			}

		}

		if !rowHasGalaxies {
			y++
		}

		lengthOfSomeRow = len(line)
		y++

	}

	fmt.Println("made map ")
	fmt.Println(galaxyCoordinates)

	colsWithOutGalaxies := []int{}
	for i := 0; i < lengthOfSomeRow; i++ {

		if !some(galaxyCoordinates, func(coord [2]int) bool {
			return coord[0] == i
		}) {
			colsWithOutGalaxies = append(colsWithOutGalaxies, i)
		}
	}

	fmt.Println("found cols without")
	fmt.Println(colsWithOutGalaxies)

	// Pair galaxies

	galaxyPairs := [][2][2]int{}
	for i := 0; i < len(galaxyCoordinates); i++ {
		for j := i + 1; j < len(galaxyCoordinates); j++ {
			galaxyPairs = append(galaxyPairs, [2][2]int{galaxyCoordinates[i], galaxyCoordinates[j]})
		}
	}

	fmt.Println("Made pairs")
	fmt.Println(len(galaxyPairs))

	total := 0

	// Do manhattan dist on each pair
	for _, pair := range galaxyPairs {
		item1 := pair[0]
		item2 := pair[1]

		// remember to account for expansion between the pairs
		expandedCols := where(colsWithOutGalaxies, func(n int) bool {
			return n > item1[0] && n < item2[0] || n < item1[0] && n > item2[0]
		})
		expansion := len(expandedCols)

		xDist := Abs(item1[0]-item2[0]) + expansion
		yDist := Abs(item1[1] - item2[1]) // vertical expansion is accounted for when building coord list

		dist := xDist + yDist
		total = total + dist
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("total")
	fmt.Println(total)

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
