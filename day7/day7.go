package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
	"unicode/utf8"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type labelledHand struct {
	cards string
	// 5 of a kind -> 7
	// 4 of a kind -> 6
	// Full house  -> 5
	// 3 of a kind -> 4
	// 2 pairs     -> 3
	// 2 of a kind -> 2
	// 1 of a kind -> 1
	handValue int
	bid       int
}

func main() {
	file, err := os.Open("./input-day-07.txt")
	check(err)

	defer file.Close()

	scanner := bufio.NewScanner(file)
	hands := []labelledHand{}

	for scanner.Scan() {
		// read line by line
		line := scanner.Text()
		split := strings.Split(line, " ")
		hand := split[0]
		bid := split[1]
		handType := labelHand(hand)
		hands = append(hands, labelledHand{cards: hand, handValue: handType, bid: atoi(bid)})
	}

	fmt.Println("Got hands")

	slices.SortFunc(hands, func(a labelledHand, b labelledHand) int {
		typeDiff := a.handValue - b.handValue
		if typeDiff != 0 {
			return typeDiff
		}

		for i := 0; i < len(a.cards); i++ {
			aRune, _ := utf8.DecodeRuneInString(a.cards[i:])
			bRune, _ := utf8.DecodeRuneInString(b.cards[i:])
			aCard := cardValue(string(aRune))
			bCard := cardValue(string(bRune))
			valueDiff := aCard - bCard
			if valueDiff != 0 {
				return valueDiff
			}

		}

		return 0

	})
	fmt.Println("sorted hands")

	totalWinnings := 0
	for i, hand := range hands {
		rank := i + 1
		winnings := rank * hand.bid
		totalWinnings += winnings
	}

	fmt.Println(totalWinnings)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}

func cardValue(card string) int {

	switch card {
	case "A":
		return 14
	case "K":
		return 13
	case "Q":
		return 12
	case "J":
		return 1
	case "T":
		return 10
	default:
		return atoi(card)

	}
}

func labelHand(hand string) int {
	// Find type
	chars := strings.Split(hand, "")
	sort.Strings(chars)

	numberOfSameCards := 1
	secondMostSameCards := 0
	highestNumberOfSameCards := 0
	highestChar := ""
	prevChar := ""
	numJokers := 0
	for i := 0; i < len(chars); i++ {
		if chars[i] == "J" {
			numJokers++
		}
		if chars[i] == prevChar && chars[i] != "J" {
			numberOfSameCards++
		} else {
			numberOfSameCards = 1
		}
		if numberOfSameCards > highestNumberOfSameCards {
			highestNumberOfSameCards = numberOfSameCards
			highestChar = chars[i]
		}
		if numberOfSameCards > secondMostSameCards && chars[i] != highestChar {
			secondMostSameCards = numberOfSameCards
		}
		prevChar = chars[i]
	}
	if numJokers == 5 {
		highestNumberOfSameCards = 5
	} else {

		highestNumberOfSameCards += numJokers
	}

	if highestNumberOfSameCards == 5 {
		return 7
	}
	if highestNumberOfSameCards == 4 {
		return 6
	}
	if highestNumberOfSameCards == 3 {

		if secondMostSameCards == 2 {
			return 5
		} else {
			return 4
		}
	}
	if highestNumberOfSameCards == 2 {
		if secondMostSameCards == 2 {
			return 3
		}
		return 2
	}
	return 1

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
