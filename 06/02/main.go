package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"slices"
	"strings"
)

var dirs = [][]int{
	{+0, +1},
	{+1, +0},
	{+0, -1},
	{-1, +0},
}

func main() {
	mat, guard, dir := readInput(os.Stdin)

	result := 0
	for row := range mat {
		for col := range mat[row] {
			if guard[0] == row && guard[1] == col {
				continue
			}

			if mat[row][col] == true {
				continue
			}

			mat[row][col] = true
			if walk(mat, slices.Clone(guard), dir) {
				result++
			}
			mat[row][col] = false
		}
	}

	fmt.Println(result)
}

func walk(mat [][]bool, guard []int, initialDir int) bool {
	H := len(mat)
	W := len(mat[0])

	visited := make([][][]bool, len(dirs))
	for i := range visited {
		visited[i] = make([][]bool, H)

		for j := range visited[i] {
			visited[i][j] = make([]bool, W)
		}
	}

	dir := initialDir
	for inBounds(mat, guard[0], guard[1]) {
		for inBounds(mat, guard[0], guard[1]) {
			if visited[dir][guard[0]][guard[1]] {
				return true
			}
			visited[dir][guard[0]][guard[1]] = true

			nr, nc := guard[0]+dirs[dir][0], guard[1]+dirs[dir][1]
			if !inBounds(mat, nr, nc) {
				guard[0], guard[1] = nr, nc
				break
			}
			if mat[nr][nc] {
				break
			}
			guard[0], guard[1] = nr, nc
		}

		dir = (dir + 1) % len(dirs)
	}

	return false
}

func inBounds(mat [][]bool, r, c int) bool {
	return r >= 0 && c >= 0 && r < len(mat) && c < len(mat[r])
}

func readInput(r io.Reader) ([][]bool, []int, int) {
	guards := map[string]int{
		">": 0,
		"v": 1,
		"<": 2,
		"^": 3,
	}

	dir := 0
	guard := make([]int, 2)
	mat := make([][]bool, 0)

	s := bufio.NewScanner(r)
	s.Split(bufio.ScanLines)

	for s.Scan() {
		row := make([]bool, 0)
		rowScanner := bufio.NewScanner(strings.NewReader(s.Text()))
		rowScanner.Split(bufio.ScanRunes)

		for rowScanner.Scan() {
			if _, ok := guards[rowScanner.Text()]; ok {
				guard[0] = len(mat)
				guard[1] = len(row)
				dir = guards[rowScanner.Text()]
			}

			if rowScanner.Text() == "#" {
				row = append(row, true)
				continue
			}

			row = append(row, false)
		}
		mat = append(mat, row)
	}

	return mat, guard, dir
}
