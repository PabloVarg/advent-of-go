package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/PabloVarg/advent-of-go/input"
)

func main() {
	adj, cases := readInput(os.Stdin)

	result := 0
	for _, order := range cases {
		if check(order, adj) {
			continue
		}

		nodes := make(map[int]struct{})
		for _, n := range order {
			nodes[n] = struct{}{}
		}

		topo := topologicalSort(adj, nodes)

		result += topo[len(order)/2]
	}

	fmt.Println(result)
}

func check(level []int, adj map[int][]int) bool {
	seen := make(map[int]struct{})
	for _, num := range level {
		for _, v := range adj[num] {
			if _, ok := seen[v]; ok {
				return false
			}
		}

		seen[num] = struct{}{}
	}

	return true
}

func topologicalSort(adj map[int][]int, nodes map[int]struct{}) []int {
	stack := make([]int, 0)
	seen := make(map[int]struct{})
	for key := range adj {
		if _, ok := seen[key]; ok {
			continue
		}
		if _, ok := nodes[key]; !ok {
			continue
		}

		dfs(key, adj, &stack, seen, nodes)
	}

	return stack
}

func dfs(v int, adj map[int][]int, stack *[]int, seen map[int]struct{}, nodes map[int]struct{}) {
	if _, ok := nodes[v]; !ok {
		return
	}
	if _, ok := seen[v]; ok {
		return
	}
	seen[v] = struct{}{}

	for _, u := range adj[v] {
		dfs(u, adj, stack, seen, nodes)
	}

	*stack = append(*stack, v)
}

func readInput(r io.Reader) (map[int][]int, [][]int) {
	s := bufio.NewScanner(r)

	adj := make(map[int][]int)
	for s.Scan() {
		if s.Text() == "" {
			break
		}

		before, after, found := strings.Cut(s.Text(), "|")
		if !found {
			panic("no separator found for order list")
		}

		u, err := strconv.Atoi(before)
		if err != nil {
			panic("expected number")
		}

		v, err := strconv.Atoi(after)
		if err != nil {
			panic("expected number")
		}

		adj[u] = append(adj[u], v)
	}

	updates := make([][]int, 0)
	for s.Scan() {
		line := input.ScanIntSlice(s, ",")
		updates = append(updates, line)
	}

	return adj, updates
}
