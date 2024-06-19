package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var alphanumric map[string]int = map[string]int{
	"zero":  0,
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func getNum(str string, last bool) int {
	i := len(str)
	if last {
		i = 0
	}
	n := -1
	for k, v := range alphanumric {
		if last {
			if j := strings.Index(str, k); j >= 0 && j > i {
				n = v
			}
		} else {
			if j := strings.Index(str, k); j >= 0 && j < i {
				n = v
			}
		}
	}
	return n
}

func trimLeft(str string) int {
	for r, l := 0, 0; l < len(str) && r < len(str); {
		if str[l] >= '0' && str[l] <= '9' {
			return int(str[l] - '0')
		}
		l++
		if r-l > 5 {
			r++
		}
		if n := getNum(str[r:l+1], false); n >= 0 {
			return n
		}
	}
	return 0
}

func trimRight(str string) int {
	for r, l := len(str)-1, len(str)-1; l >= 0 && r >= 0; {
		if str[r] >= '0' && str[r] <= '9' {
			return int(str[r] - '0')
		}
		r--
		if r-l > 5 {
			l--
		}
		if n := getNum(str[r:l+1], true); n >= 0 {
			return n
		}
	}
	return 0
}

func trim(str string) int {
	number, _ := strconv.Atoi(fmt.Sprintf("%d%d", trimLeft(str), trimRight(str)))
	return number
}

func main() {
	input, _ := os.Open("./1/input.txt")
	scanner := bufio.NewScanner(input)
	counter := 0
	for scanner.Scan() {
		line := scanner.Text()
		counter += trim(line)
	}
	fmt.Println(counter)
}
