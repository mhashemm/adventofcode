package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var maxBalls map[string]int = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func parse(line string) (int, []map[string]int) {
	sets := make([]map[string]int, 0)

	tokens := strings.Split(strings.TrimSpace(line), ":")
	gameId, _ := strconv.Atoi(strings.Split(strings.TrimSpace(tokens[0]), " ")[1])
	setTokens := strings.Split(strings.TrimSpace(tokens[1]), ";")
	for _, setToken := range setTokens {
		set := make(map[string]int, 3)
		ballTokens := strings.Split(strings.TrimSpace(setToken), ",")
		for _, ballToken := range ballTokens {
			ball := strings.Split(strings.TrimSpace(ballToken), " ")
			set[strings.TrimSpace(ball[1])], _ = strconv.Atoi(strings.TrimSpace(ball[0]))
		}

		sets = append(sets, set)
	}
	return gameId, sets
}

func main() {
	input, _ := os.Open("./2/input.txt")
	scanner := bufio.NewScanner(input)
	sumGameId := 0
	sumPower := 0
	for scanner.Scan() {
		line := scanner.Text()
		gameId, sets := parse(line)
		valid := true
		max := make(map[string]int, 3)
		for _, set := range sets {
			for k, v := range set {
				if v > maxBalls[k] {
					valid = false
				}
				if max[k] < v {
					max[k] = v
				}
			}
		}
		power := 1
		for _, v := range max {
			power *= v
		}
		sumPower += power
		if valid {
			sumGameId += gameId
		}
	}
	fmt.Println("sum of valid game id", sumGameId)
	fmt.Println("sum of set power", sumPower)
}
