package main

import (
	"bufio"
	"fmt"
	"os"
)

func transpose(grid [][]byte) [][]byte {
	transposed := make([][]byte, len(grid[0]))
	for c := 0; c < len(grid[0]); c++ {
		transposed[c] = make([]byte, len(grid))
		for r := 0; r < len(grid); r++ {
			transposed[c][r] = grid[r][c]
		}
	}
	return transposed
}

func difference(a []byte, b []byte) int {
	c := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			c++
		}
	}
	return c
}

func horizontalMirror(grid [][]byte) int {
	for i := 0; i < len(grid)-1; i++ {
		diff := 0
		t, d := i, i+1
		for t >= 0 && d < len(grid) && diff <= 1 {
			diff += difference(grid[t], grid[d])
			t--
			d++
		}
		if diff == 1 {
			return i + 1
		}
	}
	return 0
}

func main() {
	input, _ := os.Open("./13/input.txt")
	scanner := bufio.NewScanner(input)
	grids := [][][]byte{}
	grid := [][]byte{}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			grids = append(grids, grid)
			grid = [][]byte{}
			continue
		}
		grid = append(grid, []byte(line))
	}
	grids = append(grids, grid)

	sum := 0
	for _, grid := range grids {
		sum += 100 * horizontalMirror(grid)
		sum += horizontalMirror(transpose(grid))
	}
	fmt.Println(sum)
}
