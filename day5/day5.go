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

type seedMap struct {
	source string
	dest   string
	theMap map[int]int
}

func main() {
	file, err := os.Open("./input-day-05.txt")
	check(err)

	defer file.Close()

	scanner := bufio.NewScanner(file)
	seeds := []int{}

	mapmap := map[string]seedMap{}
	currentlyBuilding := ""

	for scanner.Scan() {
		fmt.Println("Are we making progress?")
		// read line by line
		line := scanner.Text()

		if strings.Contains(line, "seeds") {
			// Store seeds some place
			seeds = Map(strings.Split(strings.Split(line, "seeds: ")[1], " "), atoi)
		} else if strings.Contains(line, "map") {
			// Start building a new map
			source := strings.Split(line, "-to-")[0]
			dest := strings.Split(strings.Split(line, "-to-")[1], " ")[0]
			amap := seedMap{source: source, dest: dest, theMap: make(map[int]int)}
			currentlyBuilding = source
			mapmap[source] = amap

		} else if instAllWhiteSpace(line) {
			// Add numbers to currently building map
			destStart := atoi(strings.Split(line, " ")[0])
			srcStart := atoi(strings.Split(line, " ")[1])
			howMany := atoi(strings.Split(line, " ")[2])
			for i := 0; i < howMany; i++ {
				mapmap[currentlyBuilding].theMap[srcStart+i] = destStart + i
			}
		}

	}

	fmt.Println("Built the maps!")

	locations := []int{}

	// Get location number of each seed using maps
	for _, seed := range seeds {
		value := seed
		nextMap := "seed"
		for nextMap != "location" {
			map2 := mapmap[nextMap]
			newValue, ok := map2.theMap[value]
			if ok {
				value = newValue
			}

			nextMap = map2.dest

		}
		locations = append(locations, value)
	}

	// Find lowest location number
	lowest := slices.Min(locations)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(locations)
	fmt.Println(lowest)

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
