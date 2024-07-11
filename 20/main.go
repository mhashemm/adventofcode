package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

type Pulse byte

const (
	LowPulse Pulse = iota
	HighPulse
)

type Module struct {
	name     string
	inputs   []string
	outputs  []string
	typ      byte
	on       bool
	received map[string]Pulse
	pulse    Pulse
}

func (m *Module) send(modules map[string]*Module, input string, pulse Pulse, cache map[string]struct{}) [2]int {
	counter := [2]int{}
	counter[pulse]++

	switch m.typ {
	case 0: // output module
		m.pulse = pulse
		return counter

	case '%':
		if pulse == HighPulse {
			return counter
		}
		m.on = !m.on
		pulse = LowPulse
		if m.on {
			pulse = HighPulse
		}

	case '&':
		m.received[input] = pulse
		pulse = HighPulse
		allHigh := true
		for _, i := range m.inputs {
			if m.received[i] == LowPulse {
				allHigh = false
				break
			}
		}
		if allHigh {
			pulse = LowPulse
		} else {
			cache[m.name] = struct{}{}
		}
	}

	for _, i := range m.outputs {
		c := modules[i].send(modules, m.name, pulse, cache)
		counter[LowPulse] += c[LowPulse]
		counter[HighPulse] += c[HighPulse]
	}
	return counter
}

func higherFactor(x int, y int) int {
	if y == 0 {
		return x
	}
	return higherFactor(y, x%y)
}

func main() {
	input, _ := os.Open("./20/input.txt")
	scanner := bufio.NewScanner(input)

	modules := map[string]*Module{}

	for scanner.Scan() {
		line := scanner.Text()
		tokens := strings.Split(line, "->")
		name := strings.TrimSpace(tokens[0])
		typ := byte(0)
		if name == "broadcaster" {
			typ = '*'
		} else {
			typ = name[0]
			name = name[1:]
		}
		m, ok := modules[name]
		if !ok {
			m = &Module{}
			modules[name] = m
		}
		m.typ = typ
		m.name = name
		if m.typ == '&' {
			m.received = map[string]Pulse{}
		}

		for _, o := range strings.Split(tokens[1], ",") {
			o = strings.TrimSpace(o)
			om, ok := modules[o]
			if !ok {
				om = &Module{}
				modules[o] = om
			}
			m.outputs = append(m.outputs, o)
			om.inputs = append(om.inputs, name)
		}
	}

	counter := [2]int{}
	cache := map[string]struct{}{}
	// for range 1000 {
	// 	c := modules["broadcaster"].send(modules, "button", LowPulse, cache)
	// 	counter[LowPulse] += c[LowPulse]
	// 	counter[HighPulse] += c[HighPulse]
	// }
	fmt.Println(counter, counter[LowPulse]*counter[HighPulse])

	o := map[string]int{}
	for i := range math.MaxInt {
		modules["broadcaster"].send(modules, "button", LowPulse, cache)
		if len(o) == len(modules[modules["rx"].inputs[0]].inputs) {
			break
		}
		for _, c := range modules[modules["rx"].inputs[0]].inputs {
			_, ok := cache[c]
			if ok {
				o[c] = i + 1
			}
		}
		cache = map[string]struct{}{}
	}
	fmt.Println(o)
	total := 1
	for _, v := range o {
		total = (total * v) / higherFactor(total, v)
	}
	fmt.Println(total)
}
