package main

import (
	"fmt"
	"os"

	"github.com/PabloVarg/advent-of-go/input"
)

func main() {
	in := readInput()

	stones := make(map[uint64]uint64)
	for _, stone := range in {
		stones[uint64(stone)]++
	}

	for range 75 {
		stones = simulate(stones)
	}

	res := uint64(0)
	for _, count := range stones {
		res += count
	}

	fmt.Println(res)
}

func simulate(stones map[uint64]uint64) map[uint64]uint64 {
	res := make(map[uint64]uint64)

	for stone, count := range stones {
		switch {
		case stone == 0:
			res[1] += count
		case numDigits(stone)&1 == 0:
			first, second := partition(stone, numDigits(stone))
			res[first] += count
			res[second] += count
		default:
			res[stone*2024] += count
		}
	}

	return res
}

func numDigits(x uint64) uint64 {
	res := uint64(0)

	for x > 0 {
		x /= 10
		res++
	}

	return res
}

func partition(x, numDigits uint64) (uint64, uint64) {
	second := uint64(0)
	exp := uint64(1)

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
