package main

import (
	"fmt"
	"os"

	"github.com/PabloVarg/advent-of-go/input"
)

func main() {
	stones := readInput()

	for range 25 {
		newStones := make([]int, 0, len(stones))

		for _, stone := range stones {
			switch {
			case stone == 0:
				newStones = append(newStones, 1)
			case numDigits(stone)&1 == 0:
				first, second := partition(stone, numDigits(stone))
				newStones = append(newStones, first)
				newStones = append(newStones, second)
			default:
				newStones = append(newStones, stone*2024)
			}
		}

		stones = newStones
	}

	fmt.Println(len(stones))
}

func numDigits(x int) int {
	res := 0

	for x > 0 {
		x /= 10
		res++
	}

	return res
}

func partition(x, numDigits int) (int, int) {
	second := 0
	exp := 1

	for range numDigits / 2 {
		second += x % 10 * exp
		x /= 10
		exp *= 10
	}

	return x, second
}

func readInput() []int {
	return input.ReadIntSlice(os.Stdin)
}
