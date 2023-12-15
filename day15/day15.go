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

type Lens struct {
	label string
	value int
}

func main() {
	file, err := os.Open("./input-day-15.txt")
	check(err)

	defer file.Close()

	scanner := bufio.NewScanner(file)

	steps := []string{}
	for scanner.Scan() {
		// read line by line
		line := scanner.Text()
		steps = strings.Split(line, ",")
	}

	fmt.Println("Found", len(steps))

	hashMap := map[int][]Lens{}
	for _, v := range steps {
		split := strings.Split(v, "=")
		action := "add"
		if len(split) != 2 {
			split = strings.Split(v, "-")
			action = "rm"
		}
		label := split[0]
		focal := atoi(split[1])
		lens := Lens{
			label: label,
			value: focal,
		}

		hash := hash(label)
		list, ok := hashMap[hash]
		if ok {

			if action == "add" {
				existingIndex := slices.IndexFunc(list, func(lens Lens) bool {
					return lens.label == label
				})
				if existingIndex != -1 {
					list[existingIndex] = lens
				} else {

					list = append(list, lens)
				}
			} else {
				list = where(list, func(l Lens) bool {
					return l.label != label
				})
			}
		} else {
			if action == "add" {
				list = []Lens{lens}
			} else {
				list = []Lens{}
			}
		}
		hashMap[hash] = list

	}

	total := 0
	for i := 0; i < 256; i++ {
		box := hashMap[i]
		for j := 0; j < len(box); j++ {
			value := box[j]

			total += (1 + i) * (1 + j) * value.value
		}
	}
	fmt.Println("total")
	fmt.Println(total)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}

func hash(in string) int {
	value := 0
	for _, v := range in {
		value += int(v)
		value *= 17
		value = value % 256
	}
	return value
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
