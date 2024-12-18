package main

import (
	"fmt"
	"os"

	"github.com/PabloVarg/advent-of-go/input"
)

var dirs = [][]int{
	{+0, +1},
	{+1, +0},
	{+0, -1},
	{-1, +0},
}

func main() {
	garden := readInput()
	visited := make([][]bool, len(garden))
	for i := range garden {
		visited[i] = make([]bool, len(garden[i]))
	}

	res := 0
	for row := range garden {
		for col := range garden[row] {
			if visited[row][col] {
				continue
			}

			area, perimeter := dfs(row, col, garden, visited)
			res += area * perimeter
		}
	}

	fmt.Println(res)
}

func dfs(r, c int, garden [][]rune, visited [][]bool) (int, int) {
	m, n := len(garden), len(garden[0])

	if visited[r][c] {
		return 0, 0
	}
	visited[r][c] = true

	area, perimeter := 1, 0
	for _, dir := range dirs {
		nr, nc := r+dir[0], c+dir[1]

		if !inBounds(nr, nc, n, m) {
			perimeter++
			continue
		}

		if garden[nr][nc] != garden[r][c] {
			perimeter++
			continue
		}

		a, p := dfs(nr, nc, garden, visited)
		area += a
		perimeter += p
	}

	return area, perimeter
}

func inBounds(r, c, m, n int) bool {
	return r >= 0 && c >= 0 && r < m && c < n
}

func readInput() [][]rune {
	return input.ReadRuneMat(os.Stdin)
}
