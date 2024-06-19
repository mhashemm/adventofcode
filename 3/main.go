package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var directions [][]int = [][]int{
	{1, 0},
	{-1, 0},
	{0, 1},
	{0, -1},

	{1, 1},
	{-1, -1},
	{1, -1},
	{-1, 1},
}

func inBound(graph []string, r int, c int) bool {
	return r >= 0 && r < len(graph) && c >= 0 && c < len(graph[r])
}

func isSymbol(graph []string, r int, c int) bool {
	return inBound(graph, r, c) && graph[r][c] != '.' && (graph[r][c] < '0' || graph[r][c] > '9')
}

func travers(graph []string, r int, c int) bool {
	for _, direction := range directions {
		row, col := direction[0], direction[1]
		if isSymbol(graph, r+row, c+col) {
			return true
		}
	}
	return false
}

func getNumber(graph []string, visited [][]bool, r int, c int) int {
	if !inBound(graph, r, c) || visited[r][c] || graph[r][c] < '0' || graph[r][c] > '9' {
		return -1
	}
	start, end := c, c

	for start = c; start >= 0 && graph[r][start] >= '0' && graph[r][start] <= '9'; start-- {
		visited[r][start] = true
	}
	for end = c; end < len(graph[r]) && graph[r][end] >= '0' && graph[r][end] <= '9'; end++ {
		visited[r][end] = true
	}

	number, _ := strconv.Atoi(graph[r][start+1 : end])
	return number
}

func main() {
	graph := make([]string, 0)
	input, _ := os.Open("./3/input.txt")
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		graph = append(graph, scanner.Text())
	}

	sumPartNumbers := 0
	gearRatios := 0
	for r := 0; r < len(graph); r++ {
		for c := 0; c < len(graph[r]); c++ {
			end, isPartNumber := 0, false

			for end = c; end < len(graph[r]) && graph[r][end] >= '0' && graph[r][end] <= '9'; end++ {
				isPartNumber = isPartNumber || travers(graph, r, end)
			}

			if isPartNumber {
				n, _ := strconv.Atoi(graph[r][c:end])
				sumPartNumbers += n
			}

			c = end
		}

		for c := 0; c < len(graph[r]); c++ {
			if graph[r][c] != '*' {
				continue
			}

			visited := make([][]bool, len(graph))
			for i := 0; i < len(visited); i++ {
				visited[i] = make([]bool, len(graph[i]))
			}

			partNumberCount := 0
			ratio := 1
			for _, direction := range directions {
				row, col := direction[0], direction[1]
				if n := getNumber(graph, visited, r+row, c+col); n >= 0 {
					partNumberCount += 1
					ratio *= n
				}
			}

			if partNumberCount != 2 {
				continue
			}
			gearRatios += ratio

		}
	}
	fmt.Println("sum of part numbers", sumPartNumbers)
	fmt.Println("sum of ratios", gearRatios)
}
