package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func parse(line string) (cardId int, winningNumbers []int, cardNumbers []int) {
	tokens := strings.Split(strings.TrimSpace(line), ":")
	cardId, _ = strconv.Atoi(strings.TrimSpace(tokens[0][4:]))
	numberTokens := strings.Split(strings.TrimSpace(tokens[1]), "|")
	winningTokens := strings.Split(strings.TrimSpace(numberTokens[0]), " ")
	for _, t := range winningTokens {
		t = strings.TrimSpace(t)
		if t == "" {
			continue
		}
		n, _ := strconv.Atoi(t)
		winningNumbers = append(winningNumbers, n)
	}
	cardTokens := strings.Split(strings.TrimSpace(numberTokens[1]), " ")
	for _, t := range cardTokens {
		t = strings.TrimSpace(t)
		if t == "" {
			continue
		}
		n, _ := strconv.Atoi(t)
		cardNumbers = append(cardNumbers, n)
	}
	return cardId, winningNumbers, cardNumbers
}

func main() {
	input, _ := os.Open("./4/input.txt")
	scanner := bufio.NewScanner(input)
	cardCopies := map[int]int{}
	sum := 0
	totalCards := 0
	for scanner.Scan() {
		cardId, winningNumbers, cardNumbers := parse(scanner.Text())
		cardCopies[cardId] += 1
		totalCards += cardCopies[cardId]
		repeated := 0
		for _, n := range winningNumbers {
			if slices.Contains(cardNumbers, n) {
				repeated += 1
			}
		}
		if repeated < 1 {
			continue
		}
		sum += 1 << (repeated - 1)
		for range cardCopies[cardId] {
			for i := cardId + 1; i <= cardId+repeated; i++ {
				cardCopies[i]++
			}
		}
	}
	fmt.Println("earned points", sum)
	fmt.Println("total cards", totalCards)
}
