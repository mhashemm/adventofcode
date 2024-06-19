package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Key [3]int
type Spring struct {
	springs    []byte
	conditions []int
	cache      map[Key]uint64
}

func (s Spring) String() string {
	return fmt.Sprintf("%s %v", string(s.springs), s.conditions)
}

func backtrack(s Spring, i int, c int, curr int) uint64 {
	if i < len(s.springs) {
		char := s.springs[i]
		defer func() { s.springs[i] = char }()
	}

	if i >= len(s.springs) &&
		((c >= len(s.conditions) && curr == 0) ||
			(c == len(s.conditions)-1 && curr == s.conditions[c])) {
		return 1
	} else if i >= len(s.springs) {
		return 0
	}

	count, ok := s.cache[Key{i, c, curr}]
	if ok {
		return count
	}

	count = uint64(0)

	switch s.springs[i] {
	case '.':
		if curr > 0 && c < len(s.conditions) && curr == s.conditions[c] {
			count += backtrack(s, i+1, c+1, 0)
			s.cache[Key{i + 1, c + 1, 0}] = count
		} else if curr == 0 {
			count += backtrack(s, i+1, c, curr)
			s.cache[Key{i + 1, c, curr}] = count
		}
	case '#':
		count += backtrack(s, i+1, c, curr+1)
		s.cache[Key{i + 1, c, curr + 1}] = count
	case '?':
		for _, guess := range []byte{'.', '#'} {
			s.springs[i] = guess
			count += backtrack(s, i, c, curr)
		}
	}

	return count
}

func main() {
	input, _ := os.Open("./12/input.txt")
	scanner := bufio.NewScanner(input)
	count := 0
	sum := uint64(0)
	ch := make(chan uint64)
	for scanner.Scan() {
		count++
		line := strings.Split(scanner.Text(), " ")
		line[0] = fmt.Sprintf("%s?%s?%s?%s?%s", line[0], line[0], line[0], line[0], line[0])
		line[1] = fmt.Sprintf("%s,%s,%s,%s,%s", line[1], line[1], line[1], line[1], line[1])
		s := Spring{
			cache: make(map[Key]uint64),
		}
		for i := range line[0] {
			s.springs = append(s.springs, line[0][i])
		}
		for _, str := range strings.Split(line[1], ",") {
			n, _ := strconv.Atoi(str)
			s.conditions = append(s.conditions, n)
		}
		go func(s Spring) {
			v := backtrack(s, 0, 0, 0)
			ch <- v
		}(s)
	}
	for range count {
		sum += <-ch
	}
	close(ch)
	fmt.Println(sum)
}
