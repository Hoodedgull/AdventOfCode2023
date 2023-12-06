package main

import (
	"bufio"
	"cmp"
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

type destAndLength struct {
	sourceStart int
	destStart   int
	length      int
}

type SeedMap struct {
	source string
	dest   string
	theMap map[int]destAndLength
}

func main() {
	file, err := os.Open("./input-day-05.txt")
	check(err)

	defer file.Close()

	scanner := bufio.NewScanner(file)
	seeds := []destAndLength{}

	mapmap := map[string]SeedMap{}
	currentlyBuilding := ""

	for scanner.Scan() {
		// read line by line
		line := scanner.Text()

		if strings.Contains(line, "seeds") {
			// Store seeds some place
			seedsInput := Map(strings.Split(strings.Split(line, "seeds: ")[1], " "), atoi)
			for i := 0; i < len(seedsInput); i += 2 {
				seedStart := seedsInput[i]
				seedLength := seedsInput[i+1]
				seeds = append(seeds, destAndLength{destStart: seedStart, length: seedLength, sourceStart: seedStart})
			}
		} else if strings.Contains(line, "map") {
			// Start building a new map
			source := strings.Split(line, "-to-")[0]
			dest := strings.Split(strings.Split(line, "-to-")[1], " ")[0]
			amap := SeedMap{source: source, dest: dest, theMap: make(map[int]destAndLength)}
			currentlyBuilding = source
			mapmap[source] = amap

		} else if instAllWhiteSpace(line) {
			// Add numbers to currently building map
			destStart := atoi(strings.Split(line, " ")[0])
			srcStart := atoi(strings.Split(line, " ")[1])
			howMany := atoi(strings.Split(line, " ")[2])

			mapmap[currentlyBuilding].theMap[srcStart] = destAndLength{destStart: destStart, length: howMany, sourceStart: srcStart}

		}

	}

	fmt.Println("Built the maps!")

	minLocation := 99999999
	// Get location number of each seed using maps

	for _, seed := range seeds {
		someSlices := []destAndLength{seed}
		nextMap := "seed"
		for nextMap != "location" {
			fmt.Println(nextMap)
			fmt.Println(someSlices)
			newRange := []destAndLength{}
			map2 := mapmap[nextMap]
			for _, rangee := range someSlices {

				lol := getNextLayer(rangee.destStart, rangee.length, map2)
				newRange = append(newRange, lol...)
			}
			nextMap = map2.dest
			someSlices = newRange
		}
		localMin := slices.MinFunc(someSlices, func(loc destAndLength, loc2 destAndLength) int {
			return loc.destStart - loc2.destStart
		})

		if minLocation > localMin.destStart {
			minLocation = localMin.destStart
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(minLocation)

}

func getNextLayer(start int, length int, aamap SeedMap) []destAndLength {
	bmap := aamap.theMap
	newSlices := []destAndLength{}
	for sourceStart, v := range bmap {

		sourceEnd := sourceStart + v.length
		seedEnd := start + length
		// |----------| <-- seed range
		// ----|---------| <-- source map range
		// -----\\\\\\\\\\
		// ------|--------| <-- dest map range
		if sourceStart >= start && sourceStart < seedEnd && sourceEnd > seedEnd {
			rangeLength := v.length - (sourceEnd - seedEnd)
			newSlices = append(newSlices, destAndLength{destStart: v.destStart, length: rangeLength, sourceStart: sourceStart})
		} else
		// |----------| <-- seed range
		// ----|-----|-- <-- source map range
		// -----\\\\\\\--
		// ------|-----|- <-- dest map range
		if sourceStart >= start && sourceStart < seedEnd && sourceEnd < seedEnd && sourceEnd > start {
			rangeLength := v.length
			newSlices = append(newSlices, destAndLength{destStart: v.destStart, length: rangeLength, sourceStart: sourceStart})
		} else

		// ---|----------| <-- seed range
		// -|---------|--- <-- source map range
		// --\\\\\\\\\\---
		// ---|--------|--- <-- dest map range
		if sourceStart < start && sourceEnd <= seedEnd && sourceEnd > start {
			rangeStart := v.destStart + (start - sourceStart)
			rangeLength := v.length - (start - sourceStart)
			newSlices = append(newSlices, destAndLength{destStart: rangeStart, length: rangeLength, sourceStart: start})
		} else

		// -----|------|-- <-- seed range
		// ----|---------| <-- source map range
		// -----\\\\\\\\\\
		// ------|--------| <-- dest map range
		if sourceStart < start && sourceEnd >= seedEnd {
			rangeStart := v.destStart + (start - sourceStart)
			rangeLength := length
			newSlices = append(newSlices, destAndLength{destStart: rangeStart, length: rangeLength, sourceStart: start})
		}
	}

	fmt.Println("Mapped", start)

	slices.SortFunc(newSlices,
		func(a, b destAndLength) int {
			return cmp.Compare(a.sourceStart, b.sourceStart)
		})

	// Insert "fake" slices for parts of the seed range that is not covered by the map.

	// --|------------------------|------  <-- seed range
	// ------|--|-----|---|---------|---| <-- source map ranges
	for i := start; i < start+length; {
		mapsThatCover := where(newSlices, func(b destAndLength) bool {
			return b.sourceStart >= i
		})
		if len(mapsThatCover) > 0 {
			firstCover := mapsThatCover[0]
			uncoveredRange := firstCover.sourceStart - i
			if uncoveredRange > 0 {

				newSlices = append(newSlices, destAndLength{sourceStart: i, destStart: i, length: uncoveredRange})
			}
			i = firstCover.sourceStart + firstCover.length
		} else {
			uncoveredRange := start + length - i
			newSlices = append(newSlices, destAndLength{sourceStart: i, destStart: i, length: uncoveredRange})
			break
		}

	}

	return newSlices
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
