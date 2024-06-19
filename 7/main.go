package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

const (
	HighCard     int = iota // 1+1+1+1+1
	OnePair                 // 1+1+1+2
	TwoPair                 // 1+2+2
	ThreeOfAKind            // 1+1+3
	FullHouse               // 2+3
	FourOfAKind             // 1+4
	FiveOfAKind             // 5
)

// var strength = []byte{'2', '3', '4', '5', '6', '7', '8', '9', 'T', 'J', 'Q', 'K', 'A'}
var strength = []byte{'J', '2', '3', '4', '5', '6', '7', '8', '9', 'T', 'Q', 'K', 'A'}

func handType(s string) int {
	cards := make(map[rune]int, len(s))
	for _, c := range s {
		cards[c]++
	}
	distribution := make([]int, 0, len(cards))
	for _, v := range cards {
		distribution = append(distribution, v)
	}

	slices.Sort(distribution)

	if slices.Compare(distribution, []int{1, 1, 1, 1, 1}) == 0 {
		if _, ok := cards['J']; ok {
			return OnePair
		}
		return HighCard
	}

	if slices.Compare(distribution, []int{1, 1, 1, 2}) == 0 {
		if _, ok := cards['J']; ok {
			return ThreeOfAKind
		}
		return OnePair
	}

	if slices.Compare(distribution, []int{1, 2, 2}) == 0 {
		if n, ok := cards['J']; ok {
			if n == 2 {
				return FourOfAKind
			}
			if n == 1 {
				return FullHouse
			}
		}
		return TwoPair
	}

	if slices.Compare(distribution, []int{1, 1, 3}) == 0 {
		if _, ok := cards['J']; ok {
			return FourOfAKind
		}
		return ThreeOfAKind
	}

	if slices.Compare(distribution, []int{2, 3}) == 0 {
		if _, ok := cards['J']; ok {
			return FiveOfAKind
		}
		return FullHouse
	}

	if slices.Compare(distribution, []int{1, 4}) == 0 {
		if _, ok := cards['J']; ok {
			return FiveOfAKind
		}
		return FourOfAKind
	}

	if slices.Compare(distribution, []int{5}) == 0 {
		return FiveOfAKind
	}
	return -1
}

func main() {
	input, _ := os.Open("./7/input.txt")
	scanner := bufio.NewScanner(input)
	handBids := map[string]uint64{}
	hands := []string{}
	for scanner.Scan() {
		hand := strings.Split(scanner.Text(), " ")
		handBids[hand[0]], _ = strconv.ParseUint(hand[1], 10, 64)
		hands = append(hands, hand[0])
	}

	slices.SortFunc(hands, func(a string, b string) int {
		arank, brank := handType(a), handType(b)
		if arank > brank {
			return 1
		} else if arank < brank {
			return -1
		}
		for i := 0; i < len(a); i++ {
			ai, bi := slices.Index(strength, a[i]), slices.Index(strength, b[i])
			if ai > bi {
				return 1
			} else if ai < bi {
				return -1
			}
		}
		return 0
	})

	sum := uint64(0)
	for i, hand := range hands {
		sum += (uint64(i) + uint64(1)) * handBids[hand]
	}
	fmt.Println(sum)
}
