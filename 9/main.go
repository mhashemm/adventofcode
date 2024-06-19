package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func nextHistory(h []int) int {
	if len(h) == 0 {
		return 0
	}
	uniq := make(map[int]struct{}, len(h))
	for _, v := range h {
		uniq[v] = struct{}{}
	}
	if len(uniq) == 1 {
		return h[0]
	}

	nh := make([]int, 0, len(h)-1)

	for i := 0; i < len(h)-1; i++ {
		nh = append(nh, h[i+1]-h[i])
	}
	return h[len(h)-1] + nextHistory(nh)
}

func prevHistory(h []int) int {
	if len(h) == 0 {
		return 0
	}
	uniq := make(map[int]struct{}, len(h))
	for _, v := range h {
		uniq[v] = struct{}{}
	}
	if len(uniq) == 1 {
		return h[0]
	}

	ph := make([]int, 0, len(h)-1)

	for i := 0; i < len(h)-1; i++ {
		ph = append(ph, h[i+1]-h[i])
	}
	return h[0] - prevHistory(ph)
}

func main() {
	input, _ := os.Open("./9/input.txt")
	scanner := bufio.NewScanner(input)
	histories := [][]int{}
	for scanner.Scan() {
		line := scanner.Text()
		history := []int{}
		for _, token := range strings.Split(line, " ") {
			n, _ := strconv.Atoi(token)
			history = append(history, n)
		}
		histories = append(histories, history)
	}
	prevSum, nextSum := 0, 0
	for _, h := range histories {
		prevSum += prevHistory(h)
		nextSum += nextHistory(h)
	}
	fmt.Println(prevSum)
	fmt.Println(nextSum)
}
