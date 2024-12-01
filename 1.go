package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	list1, list2 := readInput()
	if len(list1) != len(list2) {
		fmt.Fprintln(os.Stderr, "mismatch on lists length")
		return
	}

	distance := 0
	for i := range len(list1) {
		l, r := list1[i], list2[i]
		if r > l {
			l, r = r, l
		}
		distance += l - r
	}

	fmt.Println(distance)
}

func readInput() ([]int, []int) {
	s := bufio.NewScanner(os.Stdin)
	list1, list2 := make([]int, 0), make([]int, 0)

	for s.Scan() {
		line := s.Text()

		fields := strings.Fields(line)
		if len(fields) != 2 {
			fmt.Fprintln(os.Stderr, "malformed input")
		}

		l, err := strconv.Atoi(fields[0])
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return nil, nil
		}

		r, err := strconv.Atoi(fields[1])
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return nil, nil
		}

		list1 = append(list1, l)
		list2 = append(list2, r)
	}

	slices.Sort(list1)
	slices.Sort(list2)

	return list1, list2
}
