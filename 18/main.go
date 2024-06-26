package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Point [2]int64

func (a Point) multiply(b Point) int64 {
	return (a[0] * b[1]) - (a[1] * b[0])
}

func area(points []Point) int64 {
	sum := int64(0)
	for i := 0; i < len(points); i++ {
		sum += points[i].multiply(points[(i+1)%len(points)])
	}
	if sum < 0 {
		sum *= -1
	}
	return sum / 2
}

type Step struct {
	dir   byte
	count int64
}

func stepsToPoints(steps []Step) []Point {
	points := []Point{}
	x, y := int64(0), int64(0)
	for i := 0; i < len(steps); i++ {
		step := steps[i]
		switch step.dir {
		case 'U':
			y += step.count
		case 'D':
			y -= step.count
		case 'L':
			x -= step.count
		case 'R':
			x += step.count
		}
		points = append(points, Point{x, y})
	}
	return points
}

func main() {
	input, _ := os.Open("./18/input.txt")
	scanner := bufio.NewScanner(input)
	steps := []Step{}
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		tokens := strings.Split(scanner.Text(), " ")
		// n, _ := strconv.Atoi(tokens[1])
		// steps = append(steps, Step{dir: tokens[0][0], count: n})
		// continue
		n, _ := strconv.ParseInt(tokens[2][2:7], 16, 64)
		var dir byte
		switch tokens[2][7] {
		case '0':
			dir = 'R'
		case '1':
			dir = 'D'
		case '2':
			dir = 'L'
		case '3':
			dir = 'U'
		}
		steps = append(steps, Step{dir: dir, count: n})
	}
	b := int64(0)
	for _, s := range steps {
		b += s.count
	}
	points := stepsToPoints(steps)
	// a = i + b/2 - 1
	// i = a - b/2 + 1
	i := area(points) - (b / 2) + 1
	fmt.Println(i + b)
}
