package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"slices"
	"strconv"
)

func main() {
	fs := readInput(os.Stdin)

	fmt.Println(checksum(fs))
}

func checksum(fs []int) int {
	originalFS := slices.Clone(fs)
	res := 0

	index := 0
	l, r := 0, len(fs)-1
	for l <= r {

		if l&1 == 0 {
			for range originalFS[l] {
				if fs[l] != 0 {
					res += index * l / 2
				}
				index++
			}
			l++
			continue
		}

		if l&1 == 1 {
			i, found := findLastThatFits(fs, l, fs[l])
			if !found {
				index += fs[l]
				fs[l] = 0
				l++
				continue
			}

			for range fs[i] {
				res += index * i / 2
				index++
			}
			fs[l] -= fs[i]
			fs[i] = 0
			continue
		}

		index++

	}

	return res
}

func findLastThatFits(fs []int, l, target int) (int, bool) {
	r := len(fs) - 1
	for r > 0 && l <= r {
		if fs[r] > 0 && fs[r] <= target {
			return r, true
		}

		r -= 2
	}

	return 0, false
}

func readInput(r io.Reader) []int {
	s := bufio.NewScanner(r)

	res := make([]int, 0)
	for s.Scan() {
		for _, c := range s.Text() {
			i, err := strconv.Atoi(string(c))
			if err != nil {
				panic("not a number")
			}

			res = append(res, i)

		}
	}

	return res
}
