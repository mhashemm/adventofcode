package main

import (
	"bufio"
	"crypto/sha256"
	"fmt"
	"os"
)

func north(grid [][]byte) {
	for c := 0; c < len(grid[0]); c++ {
		last := 0
		for r := 0; r < len(grid); r++ {
			if grid[r][c] == '.' && grid[last][c] != '.' {
				last = r
			}
			if grid[r][c] == 'O' && grid[last][c] == '.' {
				grid[last][c] = 'O'
				grid[r][c] = '.'
				last += 1
			}
			if grid[r][c] == '#' {
				last = r
			}
		}
	}
}

func south(grid [][]byte) {
	for c := len(grid[0]) - 1; c >= 0; c-- {
		last := len(grid) - 1
		for r := len(grid) - 1; r >= 0; r-- {
			if grid[r][c] == '.' && grid[last][c] != '.' {
				last = r
			}
			if grid[r][c] == 'O' && grid[last][c] == '.' {
				grid[last][c] = 'O'
				grid[r][c] = '.'
				last -= 1
			}
			if grid[r][c] == '#' {
				last = r
			}
		}
	}
}

func west(grid [][]byte) {
	for r := 0; r < len(grid); r++ {
		last := 0
		for c := 0; c < len(grid[r]); c++ {
			if grid[r][c] == '.' && grid[r][last] != '.' {
				last = c
			}
			if grid[r][c] == 'O' && grid[r][last] == '.' {
				grid[r][last] = 'O'
				grid[r][c] = '.'
				last += 1
			}
			if grid[r][c] == '#' {
				last = c
			}
		}
	}
}

func east(grid [][]byte) {
	for r := 0; r < len(grid); r++ {
		last := len(grid[r]) - 1
		for c := len(grid[r]) - 1; c >= 0; c-- {
			if grid[r][c] == '.' && grid[r][last] != '.' {
				last = c
			}
			if grid[r][c] == 'O' && grid[r][last] == '.' {
				grid[r][last] = 'O'
				grid[r][c] = '.'
				last -= 1
			}
			if grid[r][c] == '#' {
				last = c
			}
		}
	}
}

func totalLoad(grid [][]byte) int {
	sum := 0
	for r := 0; r < len(grid); r++ {
		count := 0
		for c := 0; c < len(grid[r]); c++ {
			if grid[r][c] == 'O' {
				count++
			}
		}
		sum += count * (len(grid) - r)
	}
	return sum
}

func checksum(grid [][]byte) string {
	sum := sha256.New()
	for _, r := range grid {
		sum.Write(r)
	}
	return fmt.Sprintf("%x", sum.Sum(nil))
}

func main() {
	input, _ := os.Open("./14/input.txt")
	scanner := bufio.NewScanner(input)
	grid := [][]byte{}
	for scanner.Scan() {
		r := []byte{}
		line := scanner.Text()
		for i := range line {
			r = append(r, line[i])
		}
		grid = append(grid, r)
	}
	cache := map[string]int{}

	for i := 0; i < 1000000000; i++ {
		north(grid)
		west(grid)
		south(grid)
		east(grid)
		j, ok := cache[checksum(grid)]
		if ok {
			fmt.Println(i, j, 1000000000-i, i-j, (1000000000-i)%(i-j), totalLoad(grid))
			if (1000000000-i)%(i-j) == 1 {
				break
			}
		}
		cache[checksum(grid)] = i
	}

	fmt.Println(totalLoad(grid))
}
