package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	fs := readInput(os.Stdin)

	fmt.Println(checksum(fs))
}

func checksum(fs []int) int {
	res := 0

	index := 0
	l, r := 0, len(fs)-1
	for l <= r {
		if fs[l] == 0 {
			l++
			continue
		}

		if l&1 == 0 {
			res += index * l / 2
			fs[l]--

		}

		if l&1 == 1 {
			if fs[r] == 0 {
				r -= 2
				continue
			}

			res += index * r / 2
			fs[l]--
			fs[r]--
		}

		index++
	}

	return res
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
