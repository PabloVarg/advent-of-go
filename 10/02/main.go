package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

var dirs = [][]int{
	{+0, +1},
	{+1, +0},
	{+0, -1},
	{-1, +0},
}

func main() {
	topMap := readInput(os.Stdin)

	result := 0
	for r := range topMap {
		for c := range topMap[r] {
			if topMap[r][c] != 0 {
				continue
			}

			result += dfs(r, c, topMap)
		}
	}

	fmt.Println(result)
}

func dfs(r, c int, topMap [][]int) int {
	m, n := len(topMap), len(topMap[0])

	if topMap[r][c] == 9 {
		return 1
	}

	result := 0
	for _, dir := range dirs {
		nr, nc := r+dir[0], c+dir[1]

		if !inBounds(nr, nc, m, n) {
			continue
		}

		if topMap[r][c]+1 != topMap[nr][nc] {
			continue
		}

		result += dfs(nr, nc, topMap)
	}

	return result
}

func inBounds(r, c, m, n int) bool {
	return r >= 0 && r < m && c >= 0 && c < n
}

func readInput(in io.Reader) [][]int {
	s := bufio.NewScanner(in)
	result := make([][]int, 0)

	for s.Scan() {
		line := s.Text()

		result = append(result, []int{})
		for _, c := range line {
			i, err := strconv.Atoi(string(c))
			if err != nil {
				panic("not a number")
			}

			result[len(result)-1] = append(result[len(result)-1], i)
		}
	}

	return result
}
