package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func parse(lines []string) (time []int, distance []int) {
	timeTokens := strings.Split(lines[0], " ")[1:]
	timeTokens = []string{strings.Join(timeTokens, "")}
	for _, t := range timeTokens {
		t := strings.TrimSpace(t)
		if t == "" {
			continue
		}
		n, _ := strconv.Atoi(t)
		time = append(time, n)
	}
	disTokens := strings.Split(lines[1], " ")[1:]
	disTokens = []string{strings.Join(disTokens, "")}
	for _, t := range disTokens {
		t := strings.TrimSpace(t)
		if t == "" {
			continue
		}
		n, _ := strconv.Atoi(t)
		distance = append(distance, n)
	}
	return time, distance
}

func main() {
	input, _ := os.Open("./6/input.txt")
	scanner := bufio.NewScanner(input)
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	time, distance := parse(lines)

	mul := 1
	for i, t := range time {
		d := distance[i]
		counter := 0
		// for s := 0; s < t; s++ {
		// 	if s*(t-s) > d {
		// 		counter++
		// 	}
		// }

		l := (float64(t) - math.Sqrt(float64((t*t)-(4*d)))) / 2
		if float64(int(l)) == l {
			l += 1
		} else {
			l = math.Ceil(l)
		}
		h := (float64(t) + math.Sqrt(float64((t*t)-(4*d)))) / 2
		if float64(int(h)) == h {
			h -= 1
		} else {
			h = math.Floor(h)
		}
		counter = int(h - l + 1)
		mul *= counter
		fmt.Println(t, d, counter)
	}
	fmt.Println(mul)
}
