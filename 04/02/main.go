package main

import (
	"fmt"
	"os"

	"github.com/PabloVarg/advent-of-go/input"
)

func main() {
	mat := input.ReadRuneMat(os.Stdin)

	result := 0
	for r := range mat {
		for c := range mat[r] {
			if mat[r][c] != 'A' {
				continue
			}

			if !check(r, c, mat) {
				continue
			}

			result++
		}
	}

	fmt.Println(result)
}

func check(r, c int, mat [][]rune) bool {
	if r <= 0 || c <= 0 || r+1 >= len(mat) || c+1 >= len(mat[r]) {
		return false
	}

	dirs := [][]int{
		{+1, +1},
		{+1, -1},
		{-1, +1},
		{-1, -1},
	}

	for d, dir := range dirs {
		nr, nc := r+dir[0], c+dir[1]
		or, oc := r+dirs[len(dirs)-d-1][0], c+dirs[len(dirs)-d-1][1]

		if mat[nr][nc] != 'M' && mat[nr][nc] != 'S' {
			return false
		}

		if mat[nr][nc] == mat[or][oc] {
			return false
		}
	}

	return true
}
