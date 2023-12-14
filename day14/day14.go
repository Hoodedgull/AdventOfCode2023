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
	file, err := os.Open("./sample-day-14.txt")
	check(err)

	defer file.Close()

	scanner := bufio.NewScanner(file)

	grid := [][]string{}
	for scanner.Scan() {
		// read line by line
		line := scanner.Text()
		rocks := strings.Split(line, "")
		grid = append(grid, rocks)
	}

	fmt.Println("Found rows x", len(grid))

	cycles := 1_000_000_000
	consecutiveEqualLoads := 0
	for i := 0; i < cycles; i++ {
		prevLoad := calculateLoad(grid)
		if i%1000000 == 0 {
			fmt.Println("Got through ", i/1000000, "million cycles")
		}
		grid = moveNorth(grid)
		grid = moveWest(grid)
		grid = moveSouth(grid)
		grid = moveEast(grid)

		load := calculateLoad(grid)
		if prevLoad == load {
			consecutiveEqualLoads++
		} else {
			consecutiveEqualLoads = 0
		}
		if consecutiveEqualLoads > 100000 {
			fmt.Println("Found 100 loads in a row that were the same. Assuming we're done")
			fmt.Println("Load is:", load)
			break
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}

func calculateLoad(grid [][]string) int {
	total := 0
	for i, row := range grid {
		distanceFromSouth := len(grid) - i
		rowString := strings.Join(row, "")
		rocksInRow := strings.Count(rowString, "O")
		total += rocksInRow * distanceFromSouth
	}
	return total
}

func moveNorth(grid [][]string) [][]string {

	grid = rotateMatrix(grid)
	grid = moveEast(grid)
	grid = rotateMatrix(grid)
	grid = rotateMatrix(grid)
	grid = rotateMatrix(grid)
	return grid
}

func moveWest(grid [][]string) [][]string {

	grid = rotateMatrix(grid)
	grid = rotateMatrix(grid)
	grid = moveEast(grid)
	grid = rotateMatrix(grid)
	grid = rotateMatrix(grid)
	return grid
}

func moveSouth(grid [][]string) [][]string {

	grid = rotateMatrix(grid)
	grid = rotateMatrix(grid)
	grid = rotateMatrix(grid)
	grid = moveEast(grid)
	grid = rotateMatrix(grid)
	return grid
}

func moveEast(grid [][]string) [][]string {

	resultGrid := [][]string{}
	for _, row := range grid {
		rowString := strings.Join(row, "")
		chunks := strings.Split(rowString, "#")
		resultChunks := []string{}
		for _, chunk := range chunks {
			numRocks := strings.Count(chunk, "O")
			numNotRocks := len(chunk) - numRocks

			resultChunks = append(resultChunks, strings.Repeat(".", numNotRocks)+strings.Repeat("O", numRocks))
		}
		resultString := strings.Join(resultChunks, "#")
		resultGrid = append(resultGrid, strings.Split(resultString, ""))
	}
	return resultGrid
}

// https://github.com/procrypt/CrackingTheCodingInterview-Golang/blob/master/Ch-1-Arrays-and-Strings/1.7-RotateMatrix/rotateMatrix.go
func rotateMatrix[T any](matrix [][]T) [][]T {

	// reverse the matrix
	for i, j := 0, len(matrix)-1; i < j; i, j = i+1, j-1 {
		matrix[i], matrix[j] = matrix[j], matrix[i]
	}

	// transpose it
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	return matrix
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
