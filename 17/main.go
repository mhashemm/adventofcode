package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

// r, c, dr, dc, steps, weight
type node [6]int

func (n *node) r() int {
	return n[0]
}
func (n *node) c() int {
	return n[1]
}
func (n *node) dr() int {
	return n[2]
}
func (n *node) dc() int {
	return n[3]
}
func (n *node) steps() int {
	return n[4]
}
func (n *node) weight() int {
	return n[5]
}
func (n *node) prev() (int, int) {
	return n.r() - n.dr(), n.c() - n.dc()
}
func (n *node) key() [5]int {
	return [5]int{n.r(), n.c(), n.dr(), n.dc(), n.steps()}
}

type MinPQ struct {
	q []node
	n int
}

func NewMinPQ(n int) *MinPQ {
	return &MinPQ{q: make([]node, n+1), n: 0}
}

func (pq *MinPQ) insert(n node) {
	pq.n++
	if pq.n >= len(pq.q) {
		pq.resize()
	}
	pq.q[pq.n] = n
	pq.swim(pq.n)
}

func (pq *MinPQ) greater(i, j int) bool {
	return pq.q[i].weight() > pq.q[j].weight()
}

func (pq *MinPQ) swim(k int) {
	for k > 1 && pq.greater(int(k/2), k) {
		pq.exch(k, (k / 2))
		k = int(k / 2)
	}
}
func (pq *MinPQ) resize() {
	nq := make([]node, len(pq.q)*2)
	copy(nq, pq.q)
	pq.q = nq
}

func (pq *MinPQ) sink(k int) {
	for 2*k <= pq.n {
		j := 2 * k
		if j < pq.n && pq.greater(j, j+1) {
			j++
		}
		if !pq.greater(k, j) {
			break
		}
		pq.exch(k, j)
		k = j
	}
}

func (pq *MinPQ) exch(i, j int) {
	pq.q[i], pq.q[j] = pq.q[j], pq.q[i]
}

func (pq *MinPQ) empty() bool {
	return pq.n == 0
}

func (pq *MinPQ) delMin() node {
	if pq.empty() {
		return node{}
	}
	n := pq.q[1]
	pq.exch(1, pq.n)
	pq.n--
	pq.sink(1)
	pq.q[pq.n+1] = node{}

	return n
}

func inBound(graph [][]int, r int, c int) bool {
	return r >= 0 && r < len(graph) && c >= 0 && c < len(graph[r])
}

var directions [][]int = [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

func dijkstra(graph [][]int, sr int, sc int) map[[5]int]int {
	q := NewMinPQ(len(graph) * len(graph[0]))
	q.insert(node{sr, sc, 0, 0, 0, 0})
	distance := map[[5]int]int{}

	for !q.empty() {
		n := q.delMin()
		if _, ok := distance[n.key()]; ok {
			continue
		}
		distance[n.key()] = n.weight()

		for _, d := range directions {
			dr, dc := d[0], d[1]
			nr, nc := n.r()+dr, n.c()+dc
			if !inBound(graph, nr, nc) ||
				(n.dr() == 1 && dr == -1) || (n.dr() == -1 && dr == 1) || (n.dc() == 1 && dc == -1) || (n.dc() == -1 && dc == 1) {
				continue
			}
			nn := node{nr, nc, dr, dc, 1, n.weight() + graph[nr][nc]}
			if n.steps() < 4 && dr != n.dr() && dc != n.dc() {
				continue
			}
			if n.dr() == dr && n.dc() == dc {
				if n.steps() == 10 {
					continue
				}
				nn[4] = n.steps() + 1
			}
			q.insert(nn)
		}
	}

	return distance
}

func main() {
	input, _ := os.Open("./17/input.txt")
	scanner := bufio.NewScanner(input)
	graph := [][]int{}
	for scanner.Scan() {
		line := scanner.Text()
		r := make([]int, len(line))
		for i := range line {
			r[i] = int(line[i] - '0')
		}
		graph = append(graph, r)
	}

	distance := dijkstra(graph, 0, 0)
	sum := math.MaxInt
	for d, r := range distance {
		if d[0] == len(graph)-1 && d[1] == len(graph[0])-1 {
			if sum > r {
				sum = r
			}
		}
	}
	fmt.Println(sum)
}
