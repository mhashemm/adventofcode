package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
)

func inBound(graph []string, r int, c int) bool {
	return r >= 0 && r < len(graph) && c >= 0 && c < len(graph[r])
}

func dfs(graph []string, visited map[[4]int]struct{}, r int, c int, dr int, dc int) {
	k := [4]int{r, c, dr, dc}
	_, ok := visited[k]
	if !inBound(graph, r, c) || ok {
		return
	}
	visited[k] = struct{}{}

	r += dr
	c += dc
	if !inBound(graph, r, c) {
		return
	}
	switch graph[r][c] {
	case '.':
		dfs(graph, visited, r, c, dr, dc)
	case '-':
		if dc != 0 {
			dfs(graph, visited, r, c, 0, dc)
		} else {
			dfs(graph, visited, r, c, 0, 1)
			dfs(graph, visited, r, c, 0, -1)
		}
	case '|':
		if dr != 0 {
			dfs(graph, visited, r, c, dr, 0)
		} else {
			dfs(graph, visited, r, c, 1, 0)
			dfs(graph, visited, r, c, -1, 0)
		}
	case '\\':
		if dr == 1 {
			dfs(graph, visited, r, c, 0, 1)
		} else if dr == -1 {
			dfs(graph, visited, r, c, 0, -1)
		} else if dc == 1 {
			dfs(graph, visited, r, c, 1, 0)
		} else if dc == -1 {
			dfs(graph, visited, r, c, -1, 0)
		} else {
			log.Fatal(r, c, dr, dc)
		}
	case '/':
		if dr == 1 {
			dfs(graph, visited, r, c, 0, -1)
		} else if dr == -1 {
			dfs(graph, visited, r, c, 0, 1)
		} else if dc == 1 {
			dfs(graph, visited, r, c, -1, 0)
		} else if dc == -1 {
			dfs(graph, visited, r, c, 1, 0)
		} else {
			log.Fatal(r, c, dr, dc)
		}
	}
}

func energize(graph []string, r int, c int, dr int, dc int) int {
	visited := make([][]int, len(graph))
	for i := range graph {
		visited[i] = make([]int, len(graph[i]))
	}
	visitedm := map[[4]int]struct{}{}

	dfs(graph, visitedm, r, c, dr, dc)
	energized := 0
	for k := range visitedm {
		r, c := k[0], k[1]
		visited[r][c] += 1
	}

	for r := 0; r < len(visited); r++ {
		for c := 0; c < len(visited[r]); c++ {
			if visited[r][c] > 0 {
				energized++
			}
		}
	}
	return energized
}

func main() {
	input, _ := os.Open("./16/input.txt")
	scanner := bufio.NewScanner(input)
	graph := []string{}
	for scanner.Scan() {
		graph = append(graph, scanner.Text())
	}
	dr, dc := 0, 0
	switch graph[0][0] {
	case '.', '-':
		dc = 1
	case '\\', '|':
		dr = 1
	case '/':
		dr = -1
	}
	fmt.Println(energize(graph, 0, 0, dr, dc))

	energized := 0
	for r, c := 0, 0; c < len(graph[r]); c++ {
		if slices.Contains([]byte{'.', '|', '\\'}, graph[r][c]) {
			if e := energize(graph, r, c, 1, 0); e > energized {
				energized = e
			}
		}
	}
	for r, c := len(graph)-1, 0; c < len(graph[r]); c++ {
		if slices.Contains([]byte{'.', '|', '/'}, graph[r][c]) {
			if e := energize(graph, r, c, -1, 0); e > energized {
				energized = e
			}
		}
	}
	for r, c := 0, 0; r < len(graph); r++ {
		if slices.Contains([]byte{'.', '-', '/'}, graph[r][c]) {
			if e := energize(graph, r, c, 0, 1); e > energized {
				energized = e
			}
		}
	}
	for r, c := 0, len(graph[0])-1; r < len(graph); r++ {
		if slices.Contains([]byte{'.', '-', '\\'}, graph[r][c]) {
			if e := energize(graph, r, c, 0, -1); e > energized {
				energized = e
			}
		}
	}
	fmt.Println(energized)
}
