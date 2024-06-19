package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

var directions = [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

type Pair [2]int

func (p Pair) String() string {
	return fmt.Sprintf("%d,%d", p[0], p[1])
}

func (p Pair) Points(graph [][]int) (sr int, sc int, er int, ec int) {
	found := 0
	for r := 0; r < len(graph); r++ {
		for c := 0; c < len(graph[r]); c++ {
			if found == 2 {
				return
			}
			if graph[r][c] == p[0] {
				sr, sc = r, c
				found++
			}
			if graph[r][c] == p[1] {
				er, ec = r, c
				found++
			}
		}
	}
	return
}

func getExpandable(graph [][]int) (rowsToExpand []int, colsToExpand []int) {
	for r := 0; r < len(graph); r++ {
		hasGalaxy := false
		for c := 0; c < len(graph[r]); c++ {
			if graph[r][c] != 0 {
				hasGalaxy = true
				break
			}
		}
		if !hasGalaxy {
			rowsToExpand = append(rowsToExpand, r)
		}
	}
	for c := 0; c < len(graph[0]); c++ {
		hasGalaxy := false
		for r := 0; r < len(graph); r++ {
			if graph[r][c] != 0 {
				hasGalaxy = true
				break
			}
		}
		if !hasGalaxy {
			colsToExpand = append(colsToExpand, c)
		}
	}
	return
}

func getPairs(from int, to int) []Pair {
	pairs := []Pair{}
	for j := from + 1; j <= to; j++ {
		pairs = append(pairs, Pair{from, j})
	}
	return pairs
}

func inBound(graph [][]int, r int, c int) bool {
	return r >= 0 && r < len(graph) && c >= 0 && c < len(graph[r])
}

func bfs(graph [][]int, rowsToExpand []int, colsToExpand []int, sr int, sc int) [][]uint64 {
	visited, distance := make([][]bool, len(graph)), make([][]uint64, len(graph))
	for r := range graph {
		visited[r], distance[r] = make([]bool, len(graph[r])), make([]uint64, len(graph[r]))
	}

	q := [][]int{{sr, sc}}
	for len(q) > 0 {
		r, c := q[0][0], q[0][1]
		q = q[1:]

		if visited[r][c] {
			continue
		}

		visited[r][c] = true

		for _, dir := range directions {
			nr, nc := r+dir[0], c+dir[1]
			if inBound(graph, nr, nc) && !visited[nr][nc] {
				q = append(q, []int{nr, nc})
				distance[nr][nc] = distance[r][c] + 1
				if slices.Contains(rowsToExpand, nr) {
					distance[nr][nc] += 1e6 - 1
				}
				if slices.Contains(colsToExpand, nc) {
					distance[nr][nc] += 1e6 - 1
				}
			}
		}
	}

	return distance
}

func main() {
	input, _ := os.Open("./11/input.txt")
	scanner := bufio.NewScanner(input)
	graph := [][]int{}
	n := 0
	for scanner.Scan() {
		line := scanner.Text()
		cols := make([]int, 0, len(line))
		for i := range line {
			id := 0
			if line[i] == '#' {
				n++
				id = n
			}
			cols = append(cols, id)
		}
		graph = append(graph, cols)
	}
	rowsToExpand, colsToExpand := getExpandable(graph)
	sum := uint64(0)

	for r := 0; r < len(graph); r++ {
		for c := 0; c < len(graph[r]); c++ {
			if graph[r][c] == 0 {
				continue
			}
			pairs := getPairs(graph[r][c], n)
			if len(pairs) == 0 {
				continue
			}

			distance := bfs(graph, rowsToExpand, colsToExpand, r, c)
			for _, pair := range pairs {
				_, _, er, ec := pair.Points(graph)
				sum += distance[er][ec]
			}
		}
	}

	fmt.Println(sum)
}
