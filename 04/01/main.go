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
			if mat[r][c] != 'X' {
				continue
			}

			result += check(r, c, mat)
		}
	}

	fmt.Println(result)
}

func check(r, c int, mat [][]rune) int {
	counter := 0
	dirs := [][]int{
		{+0, +1},
		{+0, -1},
		{+1, +0},
		{-1, +0},
		{+1, +1},
		{+1, -1},
		{-1, -1},
		{-1, +1},
	}

	for _, dir := range dirs {
		for i := range 4 {
			nr, nc := dir[0]*i+r, dir[1]*i+c

			if nr < 0 || nc < 0 || nr >= len(mat) || nc >= len(mat[nr]) {
				break
			}

			if mat[nr][nc] != []rune("XMAS")[i] {
				break
			}

			if i == 3 {
				counter++
			}
		}
	}

	return counter
}
