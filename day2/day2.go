package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	file, err := os.Open("./day2/input-day-02.txt")
	check(err)

	defer file.Close()

	scanner := bufio.NewScanner(file)
	powers := []int{}
	for scanner.Scan() {
		// read line by line
		line := scanner.Text()

		split := strings.Split(line, ":")
		// header := split[0]
		// gameid, _ := strconv.Atoi(strings.Split(header, " ")[1])

		data := split[1]
		draws := strings.Split(data, ";")

		r, _ := regexp.Compile("(\\d+) (red|blue|green)")

		maxRed := 0
		maxBlue := 0
		maxGreen := 0
		for _, draw := range draws {
			colors := strings.Split(draw, ",")
			for _, color := range colors {
				matches := r.FindStringSubmatch(color)
				num := matches[1]
				numnum, _ := strconv.Atoi(num)
				col := matches[2]
				if col == "red" && numnum > maxRed {
					maxRed = numnum
				}
				if col == "blue" && numnum > maxBlue {
					maxBlue = numnum
				}
				if col == "green" && numnum > maxGreen {
					maxGreen = numnum
				}
			}

		}

		powers = append(powers, maxRed*maxBlue*maxGreen)

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(powers)
	fmt.Println(sum(powers))

}

func sum(numbers []int) int {
	x := 0
	for _, n := range numbers {
		x = x + n
	}
	return x
}
