package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Condition struct {
	category byte
	operator byte
	value    int
	then     string
}

func (c Condition) String() string {
	return fmt.Sprintf("%c%c%d:%s", c.category, c.operator, c.value, c.then)
}

func (c Condition) bound(gt, lt int) (int, int) {
	switch c.operator {
	case '<':
		lt = c.value - 1
	case '>':
		gt = c.value + 1
	}
	return gt, lt
}

func (c Condition) oppositeBound(gt, lt int) (int, int) {
	switch c.operator {
	case '<':
		gt = c.value
	case '>':
		lt = c.value
	}
	return gt, lt
}

func (c Condition) evaluate(category byte, value int) bool {
	if c.category != category {
		return false
	}
	switch c.operator {
	case '<':
		return value < c.value
	case '>':
		return value > c.value
	}
	return false
}

type Workflow struct {
	conditions []Condition
	base       string
}

type Part struct {
	category []byte
	value    []int
}

func evaluate(wfs map[string]Workflow, p Part, w Workflow) byte {
	next := w.base
	matched := false
	for i := 0; i < len(w.conditions) && !matched; i++ {
		c := w.conditions[i]
		for j := 0; j < len(p.category) && !matched; j++ {
			if c.evaluate(p.category[j], p.value[j]) {
				if c.then == "A" || c.then == "R" {
					return c.then[0]
				}
				next = c.then
				matched = true
				break
			}
		}
	}

	if next == "A" || next == "R" {
		return next[0]
	}

	return evaluate(wfs, p, wfs[next])
}

type entry struct {
	w                                      string
	xgt, xlt, mgt, mlt, agt, alt, sgt, slt int
}

func dfs(wfs map[string]Workflow, e entry) int {
	if e.w == "R" {
		return 0
	}
	if e.w == "A" {
		return (e.xlt - e.xgt + 1) * (e.mlt - e.mgt + 1) * (e.alt - e.agt + 1) * (e.slt - e.sgt + 1)
	}

	combinations := 0
	nn := e

	for _, c := range wfs[e.w].conditions {
		ne := nn
		ne.w = c.then
		switch c.category {
		case 'x':
			ne.xgt, ne.xlt = c.bound(ne.xgt, ne.xlt)
			nn.xgt, nn.xlt = c.oppositeBound(nn.xgt, nn.xlt)
		case 'm':
			ne.mgt, ne.mlt = c.bound(ne.mgt, ne.mlt)
			nn.mgt, nn.mlt = c.oppositeBound(nn.mgt, nn.mlt)
		case 'a':
			ne.agt, ne.alt = c.bound(ne.agt, ne.alt)
			nn.agt, nn.alt = c.oppositeBound(nn.agt, nn.alt)
		case 's':
			ne.sgt, ne.slt = c.bound(ne.sgt, ne.slt)
			nn.sgt, nn.slt = c.oppositeBound(nn.sgt, nn.slt)
		}
		combinations += dfs(wfs, ne)
	}

	nn.w = wfs[e.w].base
	combinations += dfs(wfs, nn)
	return combinations
}

func main() {
	input, _ := os.Open("./19/input.txt")
	scanner := bufio.NewScanner(input)
	workflows := map[string]Workflow{}
	parts := []Part{}
	workflowsParse := true
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			workflowsParse = false
			continue
		}
		if workflowsParse {
			tokens := strings.Split(line, "{")
			wf := Workflow{}
			conditions := strings.Split(tokens[1][:len(tokens[1])-1], ",")
			for _, c := range conditions {
				if !strings.ContainsAny(c, "<>:") {
					wf.base = c
					continue
				}
				condition := Condition{
					category: c[0],
					operator: c[1],
				}
				i := 2
				for ; i < len(c) && c[i] != ':'; i++ {
				}
				condition.value, _ = strconv.Atoi(c[2:i])
				condition.then = c[i+1:]
				wf.conditions = append(wf.conditions, condition)
			}
			workflows[tokens[0]] = wf
		} else {
			ps := strings.Split(line[1:len(line)-1], ",")
			part := Part{category: make([]byte, 4), value: make([]int, 4)}
			for i, p := range ps {
				part.category[i] = p[0]
				part.value[i], _ = strconv.Atoi(p[2:])
			}
			parts = append(parts, part)
		}
	}

	sum := 0
	for _, part := range parts {
		res := evaluate(workflows, part, workflows["in"])
		if res == 'A' {
			for _, v := range part.value {
				sum += v
			}
		}
	}

	fmt.Println(sum)

	fmt.Println(dfs(workflows, entry{w: "in", xgt: 1, xlt: 4000, mgt: 1, mlt: 4000, agt: 1, alt: 4000, sgt: 1, slt: 4000}))
}
