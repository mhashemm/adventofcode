package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func checksum(str string) int {
	sum := 0
	for i := 0; i < len(str) && str[i] != '=' && str[i] != '-'; i++ {
		sum += int(str[i])
		sum *= 17
		sum %= 256
	}
	return sum
}

func main() {
	input, _ := os.Open("./15/input.txt")
	scanner := bufio.NewScanner(input)
	steps := []string{}
	for scanner.Scan() {
		steps = strings.Split(scanner.Text(), ",")
	}
	sum := 0
	boxes := [256][]string{}
	for _, step := range steps {
		hash := checksum(step)
		if step[len(step)-1] == '-' {
			i := slices.IndexFunc(boxes[hash], func(e string) bool {
				return strings.Compare(strings.Split(e, "=")[0], step[:len(step)-1]) == 0
			})
			if i >= 0 {
				if i == len(boxes[hash])-1 {
					boxes[hash] = append([]string{}, boxes[hash][:i]...)
				} else if i < len(boxes[hash])-1 {
					right, left := boxes[hash][:i], boxes[hash][i+1:]
					boxes[hash] = append([]string{}, right...)
					boxes[hash] = append(boxes[hash], left...)
				}
			}
		} else {
			i := slices.IndexFunc(boxes[hash], func(e string) bool {
				return strings.Compare(strings.Split(e, "=")[0], strings.Split(step, "=")[0]) == 0
			})
			if i < 0 {
				boxes[hash] = append(boxes[hash], step)
			} else {
				boxes[hash][i] = step
			}
		}
	}

	for i, box := range boxes {
		for j, lens := range box {
			focalLen, _ := strconv.Atoi(strings.Split(lens, "=")[1])
			sum += (i + 1) * (j + 1) * (focalLen)
		}
	}
	fmt.Println(sum)
}
