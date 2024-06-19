package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func inBound(graph []string, r int, c int) bool {
	return r >= 0 && r < len(graph) && c >= 0 && c < len(graph[r])
}

func getOpposite(t byte) byte {
	switch t {
	case 'F':
		return 'J'
	case 'J':
		return 'F'
	case 'L':
		return '7'
	case '7':
		return 'L'
	default:
		return '0'
	}
}

func getDirection(graph []string, r int, c int) [][]int {
	if !inBound(graph, r, c) {
		return nil
	}

	dirs := make([][]int, 0, 2)

	switch graph[r][c] {
	case 'S':
		if inBound(graph, r+1, c) && slices.Contains([]byte{'|', 'L', 'J'}, graph[r+1][c]) {
			dirs = append(dirs, []int{1, 0})
		}
		if inBound(graph, r-1, c) && slices.Contains([]byte{'|', 'F', '7'}, graph[r-1][c]) {
			dirs = append(dirs, []int{-1, 0})
		}
		if inBound(graph, r, c+1) && slices.Contains([]byte{'-', 'J', '7'}, graph[r][c+1]) {
			dirs = append(dirs, []int{0, 1})
		}
		if inBound(graph, r, c-1) && slices.Contains([]byte{'-', 'F', 'L'}, graph[r][c-1]) {
			dirs = append(dirs, []int{0, -1})
		}
	case '|':
		if inBound(graph, r+1, c) && slices.Contains([]byte{'|', 'J', 'L'}, graph[r+1][c]) {
			dirs = append(dirs, []int{1, 0})
		}
		if inBound(graph, r-1, c) && slices.Contains([]byte{'|', 'F', '7'}, graph[r-1][c]) {
			dirs = append(dirs, []int{-1, 0})
		}
	case '-':
		if inBound(graph, r, c+1) && slices.Contains([]byte{'-', 'J', '7'}, graph[r][c+1]) {
			dirs = append(dirs, []int{0, 1})
		}
		if inBound(graph, r, c-1) && slices.Contains([]byte{'-', 'F', 'L'}, graph[r][c-1]) {
			dirs = append(dirs, []int{0, -1})
		}
	case 'L':
		if inBound(graph, r-1, c) && slices.Contains([]byte{'|', 'F', '7'}, graph[r-1][c]) {
			dirs = append(dirs, []int{-1, 0})
		}
		if inBound(graph, r, c+1) && slices.Contains([]byte{'-', 'J', '7'}, graph[r][c+1]) {
			dirs = append(dirs, []int{0, 1})
		}
	case 'J':
		if inBound(graph, r-1, c) && slices.Contains([]byte{'|', 'F', '7'}, graph[r-1][c]) {
			dirs = append(dirs, []int{-1, 0})
		}
		if inBound(graph, r, c-1) && slices.Contains([]byte{'-', 'F', 'L'}, graph[r][c-1]) {
			dirs = append(dirs, []int{0, -1})
		}
	case '7':
		if inBound(graph, r+1, c) && slices.Contains([]byte{'|', 'L', 'J'}, graph[r+1][c]) {
			dirs = append(dirs, []int{1, 0})
		}
		if inBound(graph, r, c-1) && slices.Contains([]byte{'-', 'F', 'L'}, graph[r][c-1]) {
			dirs = append(dirs, []int{0, -1})
		}
	case 'F':
		if inBound(graph, r+1, c) && slices.Contains([]byte{'|', 'J', 'L'}, graph[r+1][c]) {
			dirs = append(dirs, []int{1, 0})
		}
		if inBound(graph, r, c+1) && slices.Contains([]byte{'-', 'J', '7'}, graph[r][c+1]) {
			dirs = append(dirs, []int{0, 1})
		}
	}
	return dirs
}

func bfs(graph []string, sr int, sc int) (int, [][]bool) {
	visited, distance := make([][]bool, len(graph)), make([][]int, len(graph))
	for r := range graph {
		visited[r], distance[r] = make([]bool, len(graph[r])), make([]int, len(graph[r]))
	}
	q, maxDistance := [][]int{{sr, sc}}, 0
	for len(q) > 0 {
		r, c := q[0][0], q[0][1]
		q = q[1:]

		if visited[r][c] {
			if distance[r][c] > maxDistance {
				maxDistance = distance[r][c]
			}
			continue
		}

		visited[r][c] = true

		for _, dir := range getDirection(graph, r, c) {
			nr, nc := r+dir[0], c+dir[1]
			if inBound(graph, nr, nc) && graph[nr][nc] != '.' && !visited[nr][nc] {
				q = append(q, []int{nr, nc})
				distance[nr][nc] = distance[r][c] + 1
			}
		}
	}

	return maxDistance, visited
}

func intersections(graph []string, visited [][]bool, r int, c int) int {
	n := 0
	for c = c + 1; c < len(graph[r]); c++ {
		if !visited[r][c] {
			continue
		}
		if graph[r][c] == '|' {
			n++
			continue
		}
		opposite := getOpposite(graph[r][c])
		if slices.Contains([]byte{'F', 'L'}, graph[r][c]) {
			for c = c + 1; c < len(graph[r]); c++ {
				if graph[r][c] != '-' {
					break
				}
			}
		}
		if !inBound(graph, r, c) || graph[r][c] == opposite {
			n++
		}
	}

	return n
}

func getS(graph []string) (int, int) {
	for r := 0; r < len(graph); r++ {
		for c := 0; c < len(graph[r]); c++ {
			if graph[r][c] == 'S' {
				return r, c
			}
		}
	}
	return 0, 0
}

func replaceS(graph []string, r int, c int) {
	if inBound(graph, r-1, c) && slices.Contains([]byte{'|', 'F', '7'}, graph[r-1][c]) &&
		inBound(graph, r+1, c) && slices.Contains([]byte{'|', 'J', 'L'}, graph[r+1][c]) {
		graph[r] = strings.Replace(graph[r], "S", "|", 1)
	} else if inBound(graph, r, c-1) && slices.Contains([]byte{'-', 'F', 'L'}, graph[r][c-1]) &&
		inBound(graph, r, c+1) && slices.Contains([]byte{'-', 'J', '7'}, graph[r][c+1]) {
		graph[r] = strings.Replace(graph[r], "S", "-", 1)
	} else if inBound(graph, r-1, c) && slices.Contains([]byte{'|', 'F', '7'}, graph[r-1][c]) &&
		inBound(graph, r, c+1) && slices.Contains([]byte{'-', 'J', '7'}, graph[r][c+1]) {
		graph[r] = strings.Replace(graph[r], "S", "L", 1)
	} else if inBound(graph, r-1, c) && slices.Contains([]byte{'|', 'F', '7'}, graph[r-1][c]) &&
		inBound(graph, r, c-1) && slices.Contains([]byte{'-', 'F', 'L'}, graph[r][c-1]) {
		graph[r] = strings.Replace(graph[r], "S", "J", 1)
	} else if inBound(graph, r+1, c) && slices.Contains([]byte{'|', 'J', 'L'}, graph[r+1][c]) &&
		inBound(graph, r, c+1) && slices.Contains([]byte{'-', 'J', '7'}, graph[r][c+1]) {
		graph[r] = strings.Replace(graph[r], "S", "F", 1)
	} else if inBound(graph, r-1, c) && slices.Contains([]byte{'|', 'J', 'L'}, graph[r-1][c]) &&
		inBound(graph, r, c-1) && slices.Contains([]byte{'-', 'F', 'L'}, graph[r][c-1]) {
		graph[r] = strings.Replace(graph[r], "S", "7", 1)
	}
}

func main() {
	input, _ := os.Open("./10/input.txt")
	scanner := bufio.NewScanner(input)
	graph := []string{}
	for scanner.Scan() {
		graph = append(graph, scanner.Text())
	}

	r, c := getS(graph)
	replaceS(graph, r, c)
	maxDistance, visited := bfs(graph, r, c)
	fmt.Println(maxDistance)
	tiles := 0
	for r := 0; r < len(graph); r++ {
		for c := 0; c < len(graph[r]); c++ {
			if !visited[r][c] {
				n := intersections(graph, visited, r, c)
				if n%2 == 1 {
					tiles++
				}
			}
		}
	}
	fmt.Println(tiles)
}
