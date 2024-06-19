package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	Left int = iota
	Right
)

func travers(instructions string, graph map[string][2]string) int {
	steps, ic := 0, 0
	current := graph["AAA"]

	for {
		if ic >= len(instructions) {
			ic = 0
		}
		steps += 1
		instruction := 0
		switch instructions[ic] {
		case 'L':
			instruction = Left
		case 'R':
			instruction = Right
		}
		if current[instruction] == "ZZZ" {
			return steps
		}
		current = graph[current[Left]]
		ic += 1
	}
}

func inAllZNodes(instruction int, current [][2]string) bool {
	for _, node := range current {
		if !strings.HasSuffix(node[instruction], "Z") {
			return false
		}
	}
	return true
}

func fastTravers(instructions string, graph map[string][2]string, start string) int {
	steps, ic := 0, 0
	current := graph[start]

	for {
		if steps < 0 {
			panic(steps)
		}
		if ic >= len(instructions) {
			ic = 0
		}
		steps += 1
		instruction := 0
		switch instructions[ic] {
		case 'L':
			instruction = Left
		case 'R':
			instruction = Right
		}
		if strings.HasSuffix(current[instruction], "Z") {
			return steps
		}

		current = graph[current[instruction]]
		ic += 1
	}
}

func higherFactor(x int, y int) int {
	fmt.Println(x, y)
	if y == 0 {
		return x
	}
	return higherFactor(y, x%y)
}

func main() {
	input, _ := os.Open("./8/input.txt")
	scanner := bufio.NewScanner(input)
	scanner.Scan()
	instructions := scanner.Text()
	scanner.Scan()
	graph := map[string][2]string{}
	for scanner.Scan() {
		line := scanner.Text()
		graph[line[:3]] = [2]string{line[7:10], line[12:15]}
	}
	// fmt.Println("start at AAA end at ZZZ", travers(instructions, graph))

	totalSteps := []int{}
	for node := range graph {
		if strings.HasSuffix(node, "A") {
			steps := fastTravers(instructions, graph, node)
			totalSteps = append(totalSteps, steps)
		}
	}

	steps := totalSteps[0]
	for i := 1; i < len(totalSteps); i++ {
		// fmt.Println(steps)
		steps = (steps * totalSteps[i]) / higherFactor(steps, totalSteps[i])
	}

	fmt.Println("start at A end at Z", steps)
}
