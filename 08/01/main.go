package main

import (
	"fmt"
	"os"

	"github.com/PabloVarg/advent-of-go/input"
)

func main() {
	mat := readInput()
	result := simulate(mat)

	fmt.Println(result)
}

func simulate(mat [][]rune) int {
	m, n := len(mat), len(mat[0])

	positions := make(map[rune][][]int)
	for row := range mat {
		for col, v := range mat[row] {
			if v == '.' {
				continue
			}

			positions[v] = append(positions[v], []int{row, col})
		}
	}

	seen := make([][]bool, len(mat))
	for row := range seen {
		seen[row] = make([]bool, len(mat[row]))
	}

	result := 0
	for _, pos := range positions {
		for i := range pos {
			for j := range pos {
				if i == j {
					continue
				}

				vec := []int{pos[i][0] - pos[j][0], pos[i][1] - pos[j][1]}

				r, c := pos[i][0]+vec[0], pos[i][1]+vec[1]
				if inBounds(m, n, r, c) && !seen[r][c] {
					seen[r][c] = true
					result++
				}
			}
		}
	}

	return result
}

func inBounds(m, n, r, c int) bool {
	return r >= 0 && r < m && c >= 0 && c < n
}

func readInput() [][]rune {
	return input.ReadRuneMat(os.Stdin)
}
