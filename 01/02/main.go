package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	list1, counts2 := readInput()
	result := 0
	for _, num := range list1 {
		count, ok := counts2[num]
		if !ok {
			continue
		}

		result += num * count
	}

	fmt.Println(result)
}

func readInput() ([]int, map[int]int) {
	s := bufio.NewScanner(os.Stdin)
	list1 := make([]int, 0)
	counts2 := make(map[int]int)

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
		counts2[r]++
	}

	return list1, counts2
}
