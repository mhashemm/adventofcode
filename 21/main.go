package main

import (
	"bufio"
	"crypto/sha256"
	"fmt"
	"os"
)

var directions = [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

func inBound(graph [][]byte, r int, c int) bool {
	return r >= 0 && r < len(graph) && c >= 0 && c < len(graph[r])
}

func pace(graph [][]byte) {
	visited := make([][]bool, len(graph))
	for i := range graph {
		visited[i] = make([]bool, len(graph[i]))
	}

	for r := 0; r < len(graph); r++ {
		for c := 0; c < len(graph[r]); c++ {
			if graph[r][c] == '#' || graph[r][c] == '.' || visited[r][c] {
				continue
			}

			graph[r][c] = '.'
			for _, dir := range directions {
				nr, nc := r+dir[0], c+dir[1]
				if nr < 0 {
					nr = len(graph) - 1
				}
				if nr >= len(graph) {
					nr = 0
				}
				if nc < 0 {
					nc = len(graph[r]) - 1
				}
				if nc >= len(graph[r]) {
					nc = 0
				}

				if !inBound(graph, nr, nc) || graph[nr][nc] == '#' || visited[nr][nc] {
					continue
				}
				graph[nr][nc] = 'O'
				visited[nr][nc] = true
			}
		}
	}
}

func checksum(grid [][]byte) string {
	sum := sha256.New()
	for _, r := range grid {
		sum.Write(r)
	}
	return fmt.Sprintf("%x", sum.Sum(nil))
}

func main() {
	input, _ := os.Open("./21/input.txt")
	scanner := bufio.NewScanner(input)

	graph := [][]byte{}
	for scanner.Scan() {
		graph = append(graph, []byte(scanner.Text()))
	}

	// cache := map[string]int{}
	max := 26501365
	// max := 64
	for i := 1; i <= max; i++ {
		pace(graph)
		// c := checksum(graph)
		// j, ok := cache[c]
		// if ok {
		// 	i = i + (max / j)
		// } else {
		// 	cache[c] = i
		// }
	}

	count := 0
	for r := 0; r < len(graph); r++ {
		for c := 0; c < len(graph[r]); c++ {
			if graph[r][c] == 'O' {
				count++
			}
			// fmt.Printf(string(graph[r][c]))
		}
		// fmt.Printf("\n")
	}
	fmt.Println(count)
}
