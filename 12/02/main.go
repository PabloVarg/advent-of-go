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
	if visited[r][c] {
		return 0, 0
	}
	visited[r][c] = true

	area, perimeter := 1, 0
	perimeter += corners(r, c, garden) + convexCorners(r, c, garden)
	for _, dir := range dirs {
		nr, nc := r+dir[0], c+dir[1]

		if !inBounds(nr, nc, garden, garden[r][c]) {
			continue
		}

		a, p := dfs(nr, nc, garden, visited)
		area += a
		perimeter += p
	}

	return area, perimeter
}

func convexCorners(r, c int, garden [][]rune) int {
	res := 0

	for i, dir := range dirs {
		dir2 := dirs[(i+1)%len(dirs)]
		dir3 := []int{dir[0] + dir2[0], dir[1] + dir2[1]}

		if inBounds(r+dir[0], c+dir[1], garden, garden[r][c]) &&
			inBounds(r+dir2[0], c+dir2[1], garden, garden[r][c]) &&
			!inBounds(r+dir3[0], c+dir3[1], garden, garden[r][c]) {
			res++
		}
	}

	return res
}

func corners(r, c int, garden [][]rune) int {
	res := 0
	for i, dir := range dirs {
		dir2 := dirs[(i+1)%len(dirs)]

		if !inBounds(r+dir[0], c+dir[1], garden, garden[r][c]) &&
			!inBounds(r+dir2[0], c+dir2[1], garden, garden[r][c]) {
			res++
		}
	}

	return res
}

func inBounds(r, c int, grid [][]rune, curr rune) bool {
	m, n := len(grid), len(grid[0])

	return r >= 0 && c >= 0 && r < m && c < n && grid[r][c] == curr
}

func readInput() [][]rune {
	return input.ReadRuneMat(os.Stdin)
}
